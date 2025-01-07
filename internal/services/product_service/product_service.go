package product_service

import (
	"fmt"
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

	start := time.Now()

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
						var charName string
						if char.UseEmoji {
							str := gomoji.FindAll(char.Name)

							if len(str) > 0 {
								charName = str[0].Character
							}
							charName = char.Name
						} else {
							charName = gomoji.RemoveEmojis(char.Name)
						}
						productNameChars = append(productNameChars, charName)
					}

					productName := strings.Join(productNameChars, " ")
					insert = append(insert, product{
						Id:        productChars[0].IdProduct,
						Name:      productName,
						UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
					})
				}

				s.deps.ProductRepository.UpdateNames(insert, "id")
			}
		}()
	}

	for i := 0; i < int(math.Ceil(float64(total)/float64(limit))); i++ {
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
		wg.Wait()   // Ждем завершения всех горутин
		close(jobs) // Закрываем канал
	}()

	fmt.Println(time.Since(start))
}
