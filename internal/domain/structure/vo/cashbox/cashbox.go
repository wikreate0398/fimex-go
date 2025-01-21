package cashbox

import "github.com/pkg/errors"

type Cashbox string

const (
	Deposit       Cashbox = "deposit"
	Balance       Cashbox = "ballance"
	Penalty       Cashbox = "penalty"
	PurchaseLimit Cashbox = "purchase_limit"
)

func (t Cashbox) IsValid() bool {
	switch t {
	case Deposit, Balance, Penalty, PurchaseLimit:
		return true
	default:
		return false
	}
}

func GetCashboxes() []Cashbox {
	return []Cashbox{
		Deposit,
		Balance,
		Penalty,
		PurchaseLimit,
	}
}

func (t Cashbox) String() string {
	return string(t)
}

func FromString(value string) (Cashbox, error) {
	var enum = Cashbox(value)

	if !enum.IsValid() {
		return "", errors.New("invalid cashbox type")
	}

	return enum, nil
}
