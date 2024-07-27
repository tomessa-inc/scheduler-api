package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"

	sg "github.com/Masterminds/squirrel"
	sg2 "github.com/n-r-w/squirrel"
)

func AddUserUsherGroup(userId string, usherGroupId string) error {
	fmt.Printf("arrived here")
	fmt.Printf(userId)
	// /	test := map[string]interface{}{"label": usherGroup.Name, "value": usherGroup.ID}

	//now := time.Now()
	id := fmt.Sprintf("%s-%s", userId, usherGroupId)

	fmt.Printf("id2 valueL: %s\n", id)

	const query = `INSERT INTO user_usher_group (id, user,usher_group) VALUES (?, ?, ?)`
	tx, err := db.DB.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id, userId, usherGroupId)
	fmt.Printf("err valueL: %s\n", err)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func DeleteUserUsherGroupByUser(userId string, usherGroupId string) error {
	fmt.Printf("arrived here")
	fmt.Printf(userId)
	// /	test := map[string]interface{}{"label": usherGroup.Name, "value": usherGroup.ID}

	//now := time.Now()
	id := fmt.Sprintf("%s-%s", userId, usherGroupId)

	fmt.Printf("id2 valueL: %s\n", id)

	const query = `DELETE FROM user_usher_group WHERE user_usher_group.user = ? AND user_usher_group.usher_group = ?`
	fmt.Printf("delete sql: %s\n", query)
	tx, err := db.DB.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(query, userId, usherGroupId)
	fmt.Printf("err valueL: %s\n", err)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func GetUserUsherGroup() ([]e.Gallery, error) {

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

func GetUserUsherGroupByUser(user string) ([]e.UserUsherGroup, error) {
	var userUsherGroupList []e.UserUsherGroup
	fmt.Println(userUsherGroupList)
	userUsherGroup := e.UserUsherGroup{}
	userUsherGroupSQL, args, err := sg.Select("user_usher_group.id, user_usher_group.user, user_usher_group.usher_group").
		From("user_usher_group").
		Where(sg.Eq{"user_usher_group.user": user}).ToSql()
	fmt.Println("the query")
	fmt.Println(userUsherGroupSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(userUsherGroupSQL, user)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&userUsherGroup)

		if err != nil {
			log.Fatalln(err)
		}

		userUsherGroupList = append(userUsherGroupList, userUsherGroup)
	}

	err = rows.Err()
	return userUsherGroupList, err
}

func GetUserUsherGroupByUsherGroup(userUsherGroup e.UserUsherGroup) ([]e.UserUsherGroup, error) {
	var userUsherGroupList []e.UserUsherGroup
	var number uint64 = 50
	fmt.Println(userUsherGroupList)
	//userUsherGroup := e.UserUsherGroup{}

	if userUsherGroup.Number > 0 {
		number = userUsherGroup.Number
	}

	userUsherGroupSQL, args, err := sg2.Select("user_usher_group.id").
		From("user_usher_group").
		LeftJoin("unavailable ON unavailable.usher_group = user_usher_group.usher_group").
		Where(sg2.And{
			sg2.Eq{"user_usher_group.usher_group": userUsherGroup.UsherGroup},
			sg2.NotIn("user_usher_group.id", sg2.Select("unavailable.user_usher_group").From("unavailable").Where(sg2.Eq{"unavailable.usher_group": userUsherGroup.UsherGroup})),
		}).
		GroupBy("user_usher_group.id").
		Offset(0).
		Limit(number).ToSql()
	//	if userUsherGroup.Number > 0 {
	//		number = userUsherGroup.Number
	//	} else {
	//	}

	//	} else {
	//		userUsherGroupSQLBuilder.Limit(userUsherGroup.Number)
	//	}
	//.ToSql()

	fmt.Println(userUsherGroupSQL)
	fmt.Println(args)
	fmt.Println(err)

	fmt.Printf("the num %d", number)
	rows, err := db.DB.Queryx(userUsherGroupSQL, userUsherGroup.UsherGroup, userUsherGroup.UsherGroup)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&userUsherGroup)

		if err != nil {
			log.Fatalln(err)
		}

		userUsherGroupList = append(userUsherGroupList, userUsherGroup)
	}

	err = rows.Err()
	return userUsherGroupList, err
}
