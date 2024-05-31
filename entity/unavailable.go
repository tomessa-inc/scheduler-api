package entity

type Unavaialble struct {
	ID             string `db:"id"`
	UserUsherGroup string `db:"user_usher_group"`
	UsherGroup     string `db:"usher_group"`
}
