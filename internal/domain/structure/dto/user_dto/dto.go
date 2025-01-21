package user_dto

type UserQueryDto struct {
	ID               int     `db:"id"`
	Deposit          float64 `db:"deposit"`
	Ballance         float64 `db:"ballance"`
	PenaltyBallance  float64 `db:"penalty_ballance"`
	PurchaseBallance float64 `db:"purchase_ballance"`
}
