package user

type User struct {
	id                int
	deposit           float64
	ballance          float64
	penalty_ballance  float64
	purchase_ballance float64
}

func NewUser() *User {

	return &User{}
}
