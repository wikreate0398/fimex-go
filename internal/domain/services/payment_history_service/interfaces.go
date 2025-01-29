package payment_history_service

import (
	"wikreate/fimex/internal/domain/entities/payment/payment_history_entity"
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/payment_vo"
)

type UserRepository interface {
	SelectWhitchHasPaymentHistory(id_user int, cashbox payment_vo.Cashbox) []user_dto.UserQueryDto
}

type PaymentHistoryRepository interface {
	SelectUserHistory(id_user int, cashbox payment_vo.Cashbox) ([]payment_history_entity.PaymentHistory, error)
	BatchUpdate(arg interface{}, identifier string)
}
