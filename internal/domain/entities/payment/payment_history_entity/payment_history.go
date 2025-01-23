package payment_history_entity

import "wikreate/fimex/internal/domain/structure/vo/payment_vo"

type PaymentHistory struct {
	id       int
	idUser   int
	increase payment_vo.Increase
	sum      float64
	ballance float64
	date     string
}

func NewPaymentHistory() *PaymentHistory {
	return &PaymentHistory{}
}

func (p *PaymentHistory) SetID(id int) {
	p.id = id
}

func (p *PaymentHistory) SetIdUser(idUser int) {
	p.idUser = idUser
}

func (p *PaymentHistory) SetIncrease(increase string) {
	p.increase = payment_vo.Increase(increase)
}

func (p *PaymentHistory) SetSum(sum float64) {
	p.sum = sum
}

func (p *PaymentHistory) SetBallance(ballance float64) {
	p.ballance = ballance
}

func (p *PaymentHistory) SetDate(date string) {
	p.date = date
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
