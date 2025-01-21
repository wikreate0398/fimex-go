package payment_history_service

import (
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/cashbox"
)

type UserRepository interface {
	SelectWhitchHasPaymentHistory(id_user int, cashbox cashbox.Cashbox) []user_dto.UserQueryDto
}
