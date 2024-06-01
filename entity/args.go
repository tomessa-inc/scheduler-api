package entity

type JsonType struct {
	Array []string
}
type LV struct {
	Label string
	Value string
}

type JsonColumn[T any] struct {
	v *T
}
