package user_repository

import (
	"fmt"
	"wikreate/fimex/internal/domain/interfaces"
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/payment_vo"
)

type UserRepositoryImpl struct {
	dbManager interfaces.DbManager
}

func NewUserRepository(db interfaces.DbManager) *UserRepositoryImpl {
	return &UserRepositoryImpl{dbManager: db}
}

func (repo UserRepositoryImpl) SelectWhitchHasPaymentHistory(id_user int, cashbox payment_vo.Cashbox) []user_dto.UserQueryDto {
	var input []user_dto.UserQueryDto

	var where []string
	args := []interface{}{}

	var userCond string
	if id_user > 0 {
		userCond = "id=? and"
		where = append(where, "id=?")
		args = append(args, id_user)
	}

	var cashboxCond string
	if len(cashbox) > 0 {
		cashboxCond = "and cashbox = ?"
		args = append(args, cashbox.String())
	}

	var query = fmt.Sprintf(`
		select id,deposit,ballance,penalty_ballance
		from users
		where %s exists(select * from payment_history where id_user = id and deleted_at is null %s)
	`, userCond, cashboxCond)

	repo.dbManager.Select(&input, query, args...)

	return input
}
