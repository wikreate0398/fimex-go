package consumers

type SortProductsConsumer struct {
	service ProductService
}

func NewSortProductsConsumer(service ProductService) *SortProductsConsumer {
	return &SortProductsConsumer{service}
}

func (r *SortProductsConsumer) Handle(result []byte) error {
	r.service.Sort()
	return nil
}
