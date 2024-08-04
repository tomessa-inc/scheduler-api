package entity

type Absence struct {
	ID    string `db:"id"`
	Range Range
}
type Range struct {
	Start Start
	End   End
}
