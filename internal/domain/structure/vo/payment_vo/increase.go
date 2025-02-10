package payment_vo

type Increase string

func (val Increase) toStr() string {
	return string(val)
}

func (val Increase) IsUp() bool {
	return string(val) == "up"
}
