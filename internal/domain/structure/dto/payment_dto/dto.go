package payment_dto

import "wikreate/fimex/internal/domain/structure/vo/payment_vo"

type RecalcBallanceInputDto struct {
	IdUser  int                `json:"id_user"`
	Cashbox payment_vo.Cashbox `json:"cashbox"`
}

type PaymentHistoryBallanceStoreDto struct {
	ID       int     `db:"id"`
	Ballance float64 `db:"ballance"`
}
