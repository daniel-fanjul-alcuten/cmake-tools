package model

// any UnexpectedToken or Empty
type Item interface {
	String() string
	ItemString() string
	Equal(Item) bool
}
