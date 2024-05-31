package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"

	sg "github.com/Masterminds/squirrel"
)

func GetMainGallery() ([]e.Gallery, error) {

	var galleryList []e.Gallery
	gallery := e.Gallery{}
	gallerySQL, args, err := sg.Select("gallery.id, gallery.image_name, gallery.gallery_name, gallery.image_name, gallery.slug, category.category_name, tag.tag_name").
		From("gallery").
		LeftJoin("category ON gallery.category = category.id").
		LeftJoin("tag ON tag.id = gallery.tag").
		Where("gallery.main_featured = 1").
		ToSql()

	//	fmt.Println(gallerySQL)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(gallerySQL)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&gallery)

		if err != nil {
			log.Fatalln(err)
		}

		galleryList = append(galleryList, gallery)
	}

	err = rows.Err()
	return galleryList, err
}

func GetGalleryByCategory(category string) ([]e.Gallery, error) {
	var galleryList []e.Gallery
	fmt.Println(category)
	gallery := e.Gallery{}
	gallerySQL, args, err := sg.Select("gallery.id, gallery.image_name, gallery.gallery_name, gallery.image_name, gallery.slug, category.category_name, tag.tag_name").
		From("gallery").
		LeftJoin("category ON gallery.category = category.id").
		LeftJoin("tag ON tag.id = gallery.tag").
		Where(sg.Eq{"gallery.category": category, "gallery.featured": 1}).ToSql()

	fmt.Println(gallerySQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(gallerySQL, category, 1)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&gallery)

		if err != nil {
			log.Fatalln(err)
		}

		galleryList = append(galleryList, gallery)
	}

	err = rows.Err()
	return galleryList, err
}

func GetGalleryByTag(tag string) ([]e.Gallery, error) {
	var galleryList []e.Gallery
	gallery := e.Gallery{}
	gallerySQL, args, err := sg.Select("gallery.id, gallery.image_name, gallery.gallery_name, gallery.image_name, gallery.slug, category.category_name, tag.tag_name").
		From("gallery").
		LeftJoin("category ON gallery.category = category.id").
		LeftJoin("tag ON tag.id = gallery.tag").
		Where(sg.Eq{"gallery.tag": tag, "gallery.featured": 1}).ToSql()

	//	fmt.Println(gallerySQL)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(gallerySQL, tag, 1)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&gallery)

		if err != nil {
			log.Fatalln(err)
		}

		galleryList = append(galleryList, gallery)
	}

	err = rows.Err()
	return galleryList, err
}

/*
func GetMainGallery2(c echo.Context) error {

	//gallery := &entity.Gallery{}

	//	return c.JSON(http.StatusOK, gallery)
	//    return gallery

	//    return c.JSON(http.StatusOK, u)
	/*
		galleries := []entity.Gallery{}
		bd, err := db.GetDB()
		if err != nil {
			panic(err.Error())
		}

		rows, err := bd.Query("SELECT id, image_name, name, description, created, slug, featured FROM gallery")

		if err != nil {
			panic(err.Error())
		}
		fmt.Print(fmt.Sprintf("%#v", rows))
		for rows.Next() {
			// In each step, scan one row
			var gallery entity.Gallery

			err = rows.Scan(&gallery.Id, &gallery.Name, &gallery.Image_Name, &gallery.Description)

			// and append it to the array
			galleries = append(galleries, gallery)
		}

		return galleries
		//	_, err = bd.Exec("DELETE FROM video_games W
		//
		//	HERE id = ?", id)

		// defer the close till after the main function has finished
		// executing
		//	return err
}
*/
