package repositories

import (
	"wikreate/fimex/internal/domain/entities/payment/payment_history_entity"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/payment_dto"
	"wikreate/fimex/internal/domain/structure/vo/payment_vo"
)

type PaymentHistoryRepositoryImpl struct {
	dbManager interfaces.DbManager
}

func NewPaymentHistoryRepository(db interfaces.DbManager) *PaymentHistoryRepositoryImpl {
	return &PaymentHistoryRepositoryImpl{dbManager: db}
}

func (repo PaymentHistoryRepositoryImpl) SelectUserHistory(
	id_user int, cashbox payment_vo.Cashbox,
) ([]payment_history_entity.PaymentHistory, error) {
	rows := repo.dbManager.Query(`
	   select id,id_user,increase,sum,ballance,date
		from payment_history
		WHERE id_user=? and cashbox=? and deleted_at is null
		order by date asc, id asc
		FOR UPDATE
	`, id_user, cashbox.String())

	defer rows.Close()

	var history []payment_history_entity.PaymentHistory
	for rows.Next() {
		var id int
		var idUser int
		var increase string
		var sum float64
		var ballance float64
		var date string

		err := rows.Scan(&id, &idUser, &increase, &sum, &ballance, &date)

		if err != nil {
			return nil, err
		}

		dto := payment_dto.PaymentHistoryQueryDto{
			ID:       id,
			IdUser:   idUser,
			Increase: increase,
			Sum:      sum,
			Ballance: ballance,
			Date:     date,
		}

		history = append(history, payment_history_entity.NewPaymentHistory(dto))
	}

	return history, nil
}

func (p PaymentHistoryRepositoryImpl) BatchUpdate(arg interface{}, identifier string) {
	p.dbManager.BatchUpdate("payment_history", identifier, arg)
}
