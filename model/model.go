package model

// any UnexpectedToken or Empty
type Item interface {
	ItemString() string
	Equal(Item) bool
}
