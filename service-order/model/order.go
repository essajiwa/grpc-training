package model

type Order struct {
	ID       int64
	Invoice  string
	Products []Product
}
