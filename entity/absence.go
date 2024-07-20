package entity

type Absence struct {
	ID   string `db:"id"`
	User string `db:"user"`
	Week string `db:"week"`
}
