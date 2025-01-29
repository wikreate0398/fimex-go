package user_entity

import (
	"wikreate/fimex/internal/domain/entities/payment/payment_history_entity"
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/payment_vo"
)

type User struct {
	id                int
	deposit           float64
	ballance          float64
	penalty_ballance  float64
	purchase_ballance float64

	paymentsHistory []payment_history_entity.PaymentHistory
}

func NewUser(dto user_dto.UserQueryDto) *User {
	return &User{
		id:                dto.ID,
		deposit:           dto.Deposit,
		ballance:          dto.Ballance,
		penalty_ballance:  dto.PenaltyBallance,
		purchase_ballance: dto.PurchaseBallance,
	}
}

func (u *User) SetPaymentHistory(paymentsHistory []payment_history_entity.PaymentHistory) {
	u.paymentsHistory = paymentsHistory
}

func (u *User) ID() int {
	return u.id
}

func (u *User) PaymentsHistory() []payment_history_entity.PaymentHistory {
	return u.paymentsHistory
}

func (u *User) BallanceValueByCashbox(val payment_vo.Cashbox) float64 {
	switch val {
	case payment_vo.Deposit:
		return u.deposit
	case payment_vo.Balance:
		return u.ballance
	case payment_vo.Penalty:
		return u.penalty_ballance
	case payment_vo.PurchaseLimit:
		return u.purchase_ballance
	default:
		return 0
	}
}

func (u *User) CountInitialBallance(cashbox payment_vo.Cashbox) float64 {
	var currentBallance = u.BallanceValueByCashbox(cashbox)
	for _, item := range u.PaymentsHistory() {
		var sum = item.Sum()
		if item.Increase().IsUp() {
			currentBallance -= sum
		} else {
			currentBallance += sum
		}
	}

	return currentBallance
}
