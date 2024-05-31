package entity

type Category struct {
	Id           int
	CategoryName string `db:"category_name"`
}
