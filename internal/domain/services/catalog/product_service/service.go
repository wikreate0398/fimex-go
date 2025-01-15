package product_service

import (
	"cmp"
	"fmt"
	"math"
	"runtime"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"wikreate/fimex/internal/domain/entities/catalog/product_entities"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/inputs"
	"wikreate/fimex/internal/helpers"
	"wikreate/fimex/pkg/workerpool"
)

type Deps struct {
	ProductRepository     interfaces.ProductRepository
	ProductCharRepository interfaces.ProductCharRepository
}

type ProductService struct {
	deps *Deps
}

func NewProductService(deps *Deps) *ProductService {
	return &ProductService{deps: deps}
}

func (s ProductService) GenerateNames(payload *inputs.GenerateNamesPayloadInput) {
	start := time.Now()

	var (
		total      = s.deps.ProductRepository.CountTotalForGenerateNames(payload)
		limit      = 700
		iterations = int(math.Ceil(float64(total) / float64(limit)))
	)

	pool := workerpool.NewWorkerPool(runtime.NumCPU())

	pool.Start()

	for i := 0; i < iterations; i++ {
		pool.AddJob(func() {
			fmt.Println("Generating names", i)
			ids := s.deps.ProductRepository.GetIdsForGenerateNames(payload, limit, i*limit)
			grouped := make(map[any][]product_entities.ProductCharDto)

			data := s.deps.ProductCharRepository.GetByProductIds(ids)

			for _, char := range data {
				grouped[char.IdProduct] = append(grouped[char.IdProduct], char)
			}

			var products [][]product_entities.ProductCharDto

			for _, items := range grouped {
				slices.SortFunc(items, func(a, b product_entities.ProductCharDto) int {
					return strings.Compare(a.Position, b.Position)
				})
				products = append(products, items)
			}

			var insert []product_entities.ProductNameDto
			for _, productChars := range products {
				var productNameChars []string
				for _, char := range productChars {
					name := char.ToEntity().PrepareNameForProduct()
					if name != "" {
						productNameChars = append(productNameChars, name)
					}
				}

				insert = append(insert, product_entities.ProductNameDto{
					Id:        productChars[0].IdProduct,
					Name:      strings.Join(productNameChars, " "),
					UpdatedAt: time.Now().Format(helpers.FullTimeFormat),
				})
			}

			if len(insert) > 0 {
				s.deps.ProductRepository.UpdateNames(insert, "id")
			}
		})
	}

	pool.Wait()

	pool.Stop()

	fmt.Println("GenerateNames", time.Since(start))
}

func (s ProductService) Sort() {
	start := time.Now()

	type job struct {
		products  []product_entities.ProductSortDto
		iteration int
	}

	var (
		wg           = sync.WaitGroup{}
		jobs         = make(chan job)
		workersCount = runtime.NumCPU()
	)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {

				var (
					subcatProducts = job.products
					iteration      = job.iteration
				)

				var insert []product_entities.ProductInsertSortDto

				sort.Slice(subcatProducts, func(a, b int) bool {
					var aProd = subcatProducts[a]
					var bProd = subcatProducts[b]

					var aPup = strings.Split(aProd.Position.String, ",")
					var bPup = strings.Split(bProd.Position.String, ",")

					for key := 0; key < len(aPup) && key < len(bPup); key++ {
						aVal, _ := strconv.Atoi(aPup[key])
						bVal, _ := strconv.Atoi(bPup[key])

						if aVal != bVal {
							return aVal < bVal
						}
					}

					return len(aPup) < len(bPup)
				})

				for _, prod := range subcatProducts {
					insert = append(insert, product_entities.ProductInsertSortDto{
						ID:        prod.ID,
						Position:  iteration,
						UpdatedAt: time.Now().Format(helpers.FullTimeFormat),
					})
					iteration++
				}

				s.deps.ProductRepository.UpdatePosition(insert, "id")
			}
		}()
	}

	go func() {
		grouped := make(map[any][]product_entities.ProductSortDto)
		var orderedKeys []string

		var data = s.deps.ProductRepository.GetForSort()

		slices.SortFunc(data, func(a, b product_entities.ProductSortDto) int {
			return cmp.Or(
				cmp.Compare(a.BrandPosition, b.BrandPosition),
				cmp.Compare(a.CatPosition, b.CatPosition),
				cmp.Compare(a.SubCatPosition, b.SubCatPosition),
			)
		})

		for _, product := range data {
			var key = fmt.Sprintf("%v.%v.%v", product.IdBrand, product.IdCategory, product.IdSubcategory)

			if _, exists := grouped[key]; !exists {
				orderedKeys = append(orderedKeys, key)
			}

			grouped[key] = append(grouped[key], product)
		}

		var num = 1
		for i, key := range orderedKeys {
			var products = grouped[key]
			jobs <- job{products: products, iteration: num}
			num = i * (len(products) + 1)
		}

		close(jobs)
	}()

	wg.Wait()

	fmt.Println("sort", time.Since(start))
}
