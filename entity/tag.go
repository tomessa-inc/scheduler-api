package entity

type Tag struct {
	Id      int
	TagName string `db:"tag_name"`
}
