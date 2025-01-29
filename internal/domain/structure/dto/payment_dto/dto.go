package payment_dto

import "wikreate/fimex/internal/domain/structure/vo/payment_vo"

type RecalcBallanceInputDto struct {
	IdUser  int                `json:"id_user"`
	Cashbox payment_vo.Cashbox `json:"cashbox"`
}

type PaymentHistoryQueryDto struct {
	ID       int     `db:"id"`
	IdUser   int     `db:"id_user"`
	Increase string  `db:"increase"`
	Ballance float64 `db:"ballance"`
	Sum      float64 `db:"sum"`
	Date     string  `db:"date"`
}

type PaymentHistoryBallanceStoreDto struct {
	ID       int     `db:"id"`
	Ballance float64 `db:"ballance"`
}
