package user_repository

import (
	"wikreate/fimex/internal/domain/structure/dto/user_dto"
	"wikreate/fimex/internal/domain/structure/vo/cashbox"
	"wikreate/fimex/internal/dto/repo_dto"
)

type UserRepositoryImpl struct {
	deps *repo_dto.Deps
}

func NewUserRepository(deps *repo_dto.Deps) *UserRepositoryImpl {
	return &UserRepositoryImpl{deps}
}

func (repo UserRepositoryImpl) SelectWhitchHasPaymentHistory(
	id_user int, cashbox cashbox.Cashbox,
) []user_dto.UserQueryDto {
	var input []user_dto.UserQueryDto
	repo.deps.DbManager.Select(&input, `
		select id,deposit,ballance,penalty_ballance
		from users
		where id=?
		and exists(select * from payment_history where id_user = id and deleted_at is null and cashbox = ?)
	`, id_user, cashbox.String())
	return input
}
