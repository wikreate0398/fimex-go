package product_service

import (
	"github.com/lovelydeng/gomoji"
	"math"
	"slices"
	"strings"
	"sync"
	"time"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure"
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

func (s ProductService) GenerateNames(payload *structure.GenerateNamesPayloadInput) {

	var (
		total        = s.deps.ProductRepository.CountTotal(payload)
		limit        = 1000
		wg           = sync.WaitGroup{}
		jobs         = make(chan [][]structure.ProductChar)
		workersCount = 5
	)

	for i := 0; i < workersCount; i++ {
		go func() {
			for products := range jobs {

				type product struct {
					Id        int    `db:"id"`
					Name      string `db:"name"`
					UpdatedAt string `db:"updated_at"`
				}

				var insert []product
				for _, productChars := range products {
					var productNameChars []string
					for _, char := range productChars {
						productNameChars = append(
							productNameChars,
							s.prepareChar(char.Name, char.UseEmoji),
						)
					}

					insert = append(insert, product{
						Id:        productChars[0].IdProduct,
						Name:      strings.Join(productNameChars, " "),
						UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
					})
				}

				s.deps.ProductRepository.UpdateNames(insert, "id")
			}
		}()
	}

	iterations := int(math.Ceil(float64(total) / float64(limit)))
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			ids := s.deps.ProductRepository.GetIds(payload, limit, i*limit)

			grouped := make(map[any][]structure.ProductChar)

			for _, char := range s.deps.ProductCharRepository.GetByProductIds(ids) {
				grouped[char.IdProduct] = append(grouped[char.IdProduct], char)
			}

			var products [][]structure.ProductChar

			for _, items := range grouped {
				slices.SortFunc(items, func(a, b structure.ProductChar) int {
					return strings.Compare(a.Position, b.Position)
				})
				products = append(products, items)
			}

			jobs <- products
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(jobs)
	}()
}

func (s ProductService) prepareChar(char string, useEmoji bool) string {
	if useEmoji {
		str := gomoji.FindAll(char)
		if len(str) > 0 {
			return str[0].Character
		}
		return char
	}

	return gomoji.RemoveEmojis(char)
}
