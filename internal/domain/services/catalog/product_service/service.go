package product_service

import (
	"fmt"
	"math"
	"runtime"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
	"wikreate/fimex/internal/domain/entities/catalog/product_entities"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/inputs"
	"wikreate/fimex/internal/helpers"
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
		total        = s.deps.ProductRepository.CountTotalForGenerateNames(payload)
		limit        = 700
		iterations   = int(math.Ceil(float64(total) / float64(limit)))
		wg           = sync.WaitGroup{}
		jobs         = make(chan int, iterations*2)
		workersCount = runtime.NumCPU()
	)

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				ids := s.deps.ProductRepository.GetIdsForGenerateNames(payload, limit, n*limit)
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
			}
		}()
	}

	go func() {
		for i := 0; i < iterations; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	wg.Wait()

	fmt.Println("GenerateNames", time.Since(start))
}

func (s ProductService) Sort() {
	type job struct {
		products  map[any]map[any]map[any][]product_entities.ProductSortDto
		iteration int
	}

	start := time.Now()
	var (
		total        = s.deps.ProductRepository.CountTotal()
		limit        = 1000
		wg           = sync.WaitGroup{}
		jobs         = make(chan job)
		workersCount = 5
	)

	for i := 0; i < workersCount; i++ {
		go func() {
			for job := range jobs {
				iteration := job.iteration
				var insert []product_entities.ProductInsertSortDto
				for _, brandProducts := range job.products {
					for _, catProducts := range brandProducts {
						for _, subcatProducts := range catProducts {

							sort.Slice(subcatProducts, func(a, b int) bool {
								var aProd = subcatProducts[a]
								var bProd = subcatProducts[b]

								var aPup = strings.Split(aProd.Position, ",")
								var bPup = strings.Split(bProd.Position, ",")

								for key := 0; key < len(aPup) && key < len(bPup); key++ {
									if aPup[key] != bPup[key] {
										return aPup[key] < bPup[key]
									}
								}
								return true
							})

							for _, prod := range subcatProducts {
								insert = append(insert, product_entities.ProductInsertSortDto{
									ID:       prod.ID,
									Position: iteration,
								})
								iteration++
							}
						}
					}
				}

				fmt.Println("Hello")

				//s.deps.ProductRepository.UpdatePosition(insert, "id")
			}
		}()
	}

	iterations := int(math.Ceil(float64(total) / float64(limit)))
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {

			grouped := make(map[any]map[any]map[any][]product_entities.ProductSortDto)

			result := s.deps.ProductRepository.GetForSort(limit, i*limit)

			for _, product := range result {
				grouped[product.IdBrand][product.IdCategory][product.IdSubcategory] =
					append(grouped[product.IdBrand][product.IdCategory][product.IdSubcategory], product)
			}

			jobs <- job{
				products:  grouped,
				iteration: i,
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(jobs)
	}()

	fmt.Println("hello", time.Since(start))
}
