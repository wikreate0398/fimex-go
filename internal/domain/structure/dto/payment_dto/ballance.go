package payment_dto

import "wikreate/fimex/internal/domain/structure/vo/cashbox"

type RecalcBallanceInputDto struct {
	IdUser  int             `json:"id_user"`
	Cashbox cashbox.Cashbox `json:"cashbox"`
}
