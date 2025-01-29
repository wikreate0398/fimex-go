package payment_history_entity

import (
	"wikreate/fimex/internal/domain/structure/dto/payment_dto"
	"wikreate/fimex/internal/domain/structure/vo/payment_vo"
)

type PaymentHistory struct {
	id       int
	idUser   int
	increase payment_vo.Increase
	sum      float64
	ballance float64
	date     string
}

func NewPaymentHistory(dto payment_dto.PaymentHistoryQueryDto) PaymentHistory {
	return PaymentHistory{
		id:       dto.ID,
		idUser:   dto.IdUser,
		increase: payment_vo.Increase(dto.Increase),
		sum:      dto.Sum,
		ballance: dto.Ballance,
		date:     dto.Date,
	}
}

func (p *PaymentHistory) ID() int {
	return p.id
}

func (p *PaymentHistory) IdUser() int {
	return p.idUser
}

func (p *PaymentHistory) Increase() payment_vo.Increase {
	return p.increase
}

func (p *PaymentHistory) Sum() float64 {
	return p.sum
}

func (p *PaymentHistory) Ballance() float64 {
	return p.ballance
}

func (p *PaymentHistory) Date() string {
	return p.date
}
