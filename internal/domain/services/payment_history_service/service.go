package payment_history_service

import (
	"fmt"
	"runtime"
	"time"
	"wikreate/fimex/internal/domain/entities/user_entity"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/payment_dto"
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/payment_vo"
	"wikreate/fimex/pkg/workerpool"
)

type Deps struct {
	UserRepo           UserRepository
	PaymentHistoryRepo PaymentHistoryRepository
	Logger             interfaces.Logger
}

type PaymentHistoryService struct {
	deps Deps
}

func NewPaymentHistoryService(deps Deps) *PaymentHistoryService {
	return &PaymentHistoryService{deps: deps}
}

func (s PaymentHistoryService) RecalcBallances(payload *payment_dto.RecalcBallanceInputDto) {
	var start = time.Now()

	var cashboxes []payment_vo.Cashbox
	for _, val := range payment_vo.GetCashboxes() {
		var eqWithVal = payload.Cashbox.String() != "" && payload.Cashbox == val
		var exceptIfEmpty = payload.Cashbox.String() == "" && val != payment_vo.PurchaseLimit

		if eqWithVal || exceptIfEmpty {
			cashboxes = append(cashboxes, val)
		}
	}

	pool := workerpool.NewWorkerPool(runtime.NumCPU())

	pool.Start()

	for _, val := range cashboxes {
		var users = s.deps.UserRepo.SelectWhitchHasPaymentHistory(payload.IdUser, val)

		for _, user := range users {
			pool.AddJob(func(user user_dto.UserQueryDto, cashboxType payment_vo.Cashbox) func() {
				return func() {
					var userEntity = user_entity.NewUser(user)

					history, err := s.deps.PaymentHistoryRepo.SelectUserHistory(userEntity.ID(), cashboxType)

					if err != nil {
						s.deps.Logger.PanicOnErr(
							err,
							fmt.Sprintf("Failed to get user history, ID %v", userEntity.ID()),
						)
					}

					if len(history) <= 0 {
						return
					}

					userEntity.SetPaymentHistory(history)

					var initialBallance = userEntity.CountInitialBallance(cashboxType)

					inserts := make([]payment_dto.PaymentHistoryBallanceStoreDto, 0, len(history))

					for _, item := range userEntity.PaymentsHistory() {
						if item.Increase().IsUp() {
							initialBallance += item.Sum()
						} else {
							initialBallance -= item.Sum()
						}

						inserts = append(inserts, payment_dto.PaymentHistoryBallanceStoreDto{
							ID:       item.ID(),
							Ballance: initialBallance,
						})
					}

					if len(inserts) > 0 {
						s.deps.PaymentHistoryRepo.BatchUpdate(inserts, "id")
					}
				}
			}(user, val))
		}
	}

	pool.Stop()
	pool.Wait()

	fmt.Println("payment history", time.Since(start))
}
