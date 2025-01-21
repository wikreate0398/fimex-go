package payment_history_service

import (
	"runtime"
	"wikreate/fimex/internal/domain/structure/dto/payment_dto"
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/cashbox"
	"wikreate/fimex/pkg/workerpool"
)

type Deps struct {
	UserRepo UserRepository
}

type PaymentHistoryService struct {
	deps *Deps
}

func NewPaymentHistoryService(deps *Deps) *PaymentHistoryService {
	return &PaymentHistoryService{deps: deps}
}

func (s PaymentHistoryService) RecalcBallances(payload *payment_dto.RecalcBallanceInputDto) {
	var cashboxes []cashbox.Cashbox
	for _, val := range cashbox.GetCashboxes() {
		var eqWithVal = payload.Cashbox.String() != "" && payload.Cashbox == val
		var exceptIfEmpty = payload.Cashbox.String() == "" && val != cashbox.PurchaseLimit
		if eqWithVal || exceptIfEmpty {
			cashboxes = append(cashboxes, val)
		}
	}

	pool := workerpool.NewWorkerPool(runtime.NumCPU())

	pool.Start()

	for _, val := range cashboxes {
		var users = s.deps.UserRepo.SelectWhitchHasPaymentHistory(payload.IdUser, val)

		for _, user := range users {
			pool.AddJob(func(user user_dto.UserQueryDto) func() {
				return func() {

				}
			}(user))
		}
	}

	pool.Stop()
	pool.Wait()
}
