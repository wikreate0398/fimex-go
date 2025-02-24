package product_service

import (
	"fmt"
	"go.uber.org/fx"
	"math"
	"runtime"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"wikreate/fimex/internal/domain/entities/catalog/product"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/catalog_dto"
	"wikreate/fimex/internal/helpers"
	"wikreate/fimex/pkg/workerpool"
)

type Params struct {
	fx.In

	ProductRepository     ProductRepository
	ProductCharRepository ProductCharRepository
	Logger                interfaces.Logger
}

type ProductService struct {
	*Params
}

func NewProductService(params Params) *ProductService {
	return &ProductService{&params}
}

func (s ProductService) GenerateNames(payload *catalog_dto.GenerateNamesInputDto) {
	start := time.Now()

	var (
		total      = s.ProductRepository.CountTotalForGenerateNames(payload)
		limit      = 700
		iterations = int(math.Ceil(float64(total) / float64(limit)))
	)

	pool := workerpool.NewWorkerPool(runtime.NumCPU())

	pool.Start()

	for i := 0; i < iterations; i++ {
		pool.AddJob(func(i int) func() {
			return func() {
				ids := s.ProductRepository.GetIdsForGenerateNames(payload, limit, i*limit)
				grouped := make(map[any][]catalog_dto.ProductCharQueryDto)

				data := s.ProductCharRepository.GetByProductIds(ids)

				for _, char := range data {
					grouped[char.IdProduct] = append(grouped[char.IdProduct], char)
				}

				var products [][]catalog_dto.ProductCharQueryDto

				for _, items := range grouped {
					slices.SortFunc(items, func(a, b catalog_dto.ProductCharQueryDto) int {
						return strings.Compare(a.Position, b.Position)
					})
					products = append(products, items)
				}

				var insert []catalog_dto.ProductNameStoreDto
				for _, productChars := range products {
					var productNameChars []string
					for _, char := range productChars {
						name := product.NewProductChar(char).PrepareNameForProduct()
						if name != "" {
							productNameChars = append(productNameChars, name)
						}
					}

					insert = append(insert, catalog_dto.ProductNameStoreDto{
						Id:        productChars[0].IdProduct,
						Name:      strings.Join(productNameChars, " "),
						UpdatedAt: time.Now().Format(helpers.FullTimeFormat),
					})
				}

				if len(insert) > 0 {
					s.ProductRepository.UpdateNames(insert, "id")
				}
			}
		}(i))
	}

	pool.Wait()

	pool.Stop()

	s.Logger.Info(fmt.Sprintf("GenerateNames %v", time.Since(start)))
}

func (s ProductService) Sort() {
	start := time.Now()

	type job struct {
		products  []catalog_dto.ProductSortQueryDto
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

				var insert []catalog_dto.ProductSortStoreDto

				sort.Slice(subcatProducts, func(a, b int) bool {
					var aPup = strings.Split(subcatProducts[a].Position.String, ",")
					var bPup = strings.Split(subcatProducts[b].Position.String, ",")

					for k := 0; k < len(aPup) && k < len(bPup); k++ {
						aVal, _ := strconv.Atoi(aPup[k])
						bVal, _ := strconv.Atoi(bPup[k])

						if aVal != bVal {
							return aVal < bVal
						}
					}
					return false
				})

				for _, prod := range subcatProducts {
					insert = append(insert, catalog_dto.ProductSortStoreDto{
						ID:        prod.ID,
						Position:  iteration,
						UpdatedAt: time.Now().Format(helpers.FullTimeFormat),
					})
					iteration++
				}

				s.ProductRepository.UpdatePosition(insert, "id")
			}
		}()
	}

	go func() {
		grouped := make(map[any][]catalog_dto.ProductSortQueryDto)
		var orderedKeys []string

		var data = s.ProductRepository.GetForSort()

		for _, prod := range data {
			var key = fmt.Sprintf("%v.%v.%v", prod.IdBrand, prod.IdCategory, prod.IdSubcategory)

			if _, exists := grouped[key]; !exists {
				orderedKeys = append(orderedKeys, key)
			}

			grouped[key] = append(grouped[key], prod)
		}

		var num = 1
		for _, key := range orderedKeys {
			var products = grouped[key]

			jobs <- job{products: products, iteration: num}
			num += len(products)
		}

		close(jobs)
	}()

	wg.Wait()

	s.Logger.Info(fmt.Sprintf("Sort Products %v", time.Since(start)))
}
