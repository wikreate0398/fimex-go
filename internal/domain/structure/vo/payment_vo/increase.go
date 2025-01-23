package payment_vo

type Increase string

func (val Increase) toStr() string {
	return string(val)
}

func (val Increase) IsUp() bool {
	if string(val) == "up" {
		return true
	}
	return false
}
