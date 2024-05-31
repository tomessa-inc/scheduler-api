package entity

type Gallery struct {
	Id          int
	ImageName   string `db:"image_name"`
	GalleryName string `db:"gallery_name"`
	Description string
	Created     string
	Slug        string
	Featured    string
	Category
	Tag
}
