package consumers

import (
	"wikreate/fimex/internal/domain/core"
)

type Consumers struct {
	GenerateNamesQueue Consumer
}

func NewHandlers(application *core.Application) *Consumers {
	return &Consumers{
		GenerateNamesQueue: NewGenerateNamesConsumer(application.Service.ProductService),
	}
}
