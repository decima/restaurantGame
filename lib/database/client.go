package database

type Client interface {
	Disconnect() error
}

type ID string

type Criterion struct {
	Field string
	Value interface{}
}
type Sort struct {
	Field string
	Order int
}

type Limit struct {
	Count  int
	Offset int
}
