package entity

type UserUsherGroup struct {
	ID         string `db:"id"`
	User       string `db:"user"`
	UsherGroup string `db:"usher_group"`
	Number     uint64 `uint64:"usher_group"`
}
