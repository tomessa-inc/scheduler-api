package entity

type Schedule struct {
	ID             string `db:"id"`
	UserUsherGroup string `db:"user_usher_group"`
	Week           string `db:"week"`
	CreatedAt      int    `db:"created_at"`
}

type GetSchedule struct {
	UserId    string `db:"userId"`
	Type      string `json:"type"`
	RequestId string `json:"requestId"`
}

type SetUnAvailable struct {
	ID string `db:"id"`
}

type SetAvailable struct {
	Week           string `db:"week"`
	UserUsherGroup string `db:"user_usher_group"`
}

type ScheduleList struct {
	ID   string `db:"id"`
	User string `db:"user"`
	Mass string `db:"mass"`
	Week string `db:"week"`
}

type ScheduleUser struct {
	ScheduleId string `db:"scheduleId"`
	UserId     string `db:"userId"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	Week       string `db:"week"`
}

type UsersBasedOnSchedule struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Description string `db:"description"`
	Phone       string `db:"phone"`
	//UsherGroups []UsherGroups `db:"usher_groups"`
	UsherGroup      *string `db:"usher_group"`
	Password        string  `db:"password"`
	UserUsherGroups []UserUsherGroups
}

type ScheduleTest struct {
	ID              string            `db:"id"`
	UserUsherGroup  string            `db:"user_usher_group"`
	Week            string            `db:"week"`
	CreatedAt       int               `db:"created_at"`
	UserUsherGroups []UserUsherGroups `db:"usher_groups"`
}

type UserUsherGroups struct {
	UserUsherGroups []UserUsherGroups `db:"usher_groups"`
}
