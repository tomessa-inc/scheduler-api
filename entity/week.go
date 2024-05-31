package entity

type Week struct {
	ID         string `db:"id"`
	Hour       int    `db:"hour"`
	Minute     int    `db:"minute"`
	Day        int    `db:"day"`
	Month      int    `db:"month"`
	Year       int    `db:"year"`
	UsherGroup string `db:"usher_group"`
	CreatedAt  int    `db:"created_at"`
}

type Analytics struct {
	Visitors  int `json:"visitors"`
	PageViews int `json:"page_views"`
}

type CreateWeek struct {
	Begin  []Start
	Finish []End
	Mass   []MassTime
}

// Error implements error.
func (CreateWeek) Error() string {
	panic("unimplemented")
}

type Start struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type End struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}
type MassTime struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
	Day    int `json:"day"`
	Month  int `json:"month"`
	Year   int `json:"year"`
}
