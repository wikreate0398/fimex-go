package consumers

type Consumer interface {
	Handle()
	ToStruct(result []byte)
}
