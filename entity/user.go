package entity

type User struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	Email       string  `db:"email"`
	Description string  `db:"description"`
	Phone       string  `db:"phone"`
	UsherGroup  *string `db:"usher_group"`
	Password    string  `db:"password"`
}

type AuthCheck struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuth struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Description string `db:"description"`
	Phone       string `db:"phone"`
	//UsherGroups []UsherGroups `db:"usher_groups"`
}

type UserWrite struct {
	ID          string        `db:"id"`
	Name        string        `db:"name"`
	Email       string        `db:"email"`
	Description string        `db:"description"`
	Phone       string        `db:"phone"`
	UsherGroups []UsherGroups `db:"usher_groups"`
	// UsherGroup [string] `db:"usher_group"`
}

type UsherGroups struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type UserResetPassword struct {
	ID       string `db:"id"`
	Password string `db:"password"`
}
