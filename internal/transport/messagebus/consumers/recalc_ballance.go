package consumers

import (
	"encoding/json"
	"wikreate/fimex/internal/domain/services/payment_history_service"
	"wikreate/fimex/internal/domain/structure/dto/payment_dto"
)

type RecalcHistoryBallanceConsumer struct {
	service *payment_history_service.PaymentHistoryService
}

func NewRecalcHistoryBallanceConsumer(
	service *payment_history_service.PaymentHistoryService,
) *RecalcHistoryBallanceConsumer {

	return &RecalcHistoryBallanceConsumer{service}
}

func (r *RecalcHistoryBallanceConsumer) Handle(result []byte) error {
	var input = new(payment_dto.RecalcBallanceInputDto)
	if err := json.Unmarshal(result, &input); err != nil {
		return err
	}

	r.service.RecalcBallances(input)

	return nil
}
