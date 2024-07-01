package model

import (
	"errors"
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"
	"strings"

	sg "github.com/Masterminds/squirrel"
)

func AddUser(user *e.User) error {
	fmt.Printf("name valueL: %s\n", user.Name)
	//now := time.Now()
	id := strings.ToLower(user.Name)

	if strings.Contains(id, ",") {
		temp := strings.Split(id, ",")
		fmt.Printf("temp valueL: %s\n", temp)
		id = fmt.Sprintf("%s %s", temp[1], temp[0])
		fmt.Printf("id valueL: %s\n", id)

	}
	id = strings.Trim(id, " ")
	id = strings.Replace(id, " ", "-", -1)
	fmt.Printf("id2 valueL: %s\n", id)

	addQuery, args, err := sg.Insert(`user`).
		Columns(`id`, `name`, `email`, `description`, `phone`).
		Values(id, user.Name, user.Email, user.Description, user.Phone).
		ToSql()
	fmt.Println(addQuery)
	fmt.Println(args)

	tx, err := db.DB.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(addQuery, id, user.Name, user.Email, user.Description, user.Phone)
	fmt.Printf("err valueL: %s\n", err)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func UpdateUser(user *e.User) error {
	updateSQL, args, err := sg.Update(`user`).
		Set(`name`, user.Name).
		Set(`email`, user.Email).
		Set(`phone`, user.Phone).
		Set(`description`, user.Description).
		Where(sg.Eq{"id": user.ID}).
		ToSql()

	fmt.Println(updateSQL)
	fmt.Println(args)

	//		SET name = ?, email = ?, phone = ?, description = ? WHERE id = ?`

	tx, err := db.DB.Begin()
	if err != nil {

		return err
	}

	res, err := tx.Exec(updateSQL, user.Name, user.Email, user.Phone, user.Description, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("User has not been affected")
	}

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New("Strange behaviour. Total affected is : " + string(rowsAffected))
	}

	tx.Commit()

	return nil
}

func GetUsers(pageIndex uint64, pageSize uint64, field string, order string) ([]e.User, error) {

	var userList []e.User
	user := e.User{}

	offset := ((pageIndex - 1) * pageSize)

	orderBy := fmt.Sprintf("%s %s", field, order)

	userListSQL, args, err := sg.Select("user.id, user.name, user.email, user.phone").
		From("user").
		OrderBy(orderBy).
		Limit(pageSize).
		Offset(offset).
		ToSql()

	fmt.Println(userListSQL)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(userListSQL)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}

		userList = append(userList, user)
	}
	///	test := fmt.Sprintf("{data: %s}", userList)

	err = rows.Err()

	return userList, err
}

func GetUsersByPrefix(pageIndex uint64, pageSize uint64, field string, order string, prefix string) ([]e.User, error) {

	var userList []e.User
	user := e.User{}

	offset := ((pageIndex - 1) * pageSize)

	orderBy := fmt.Sprintf("%s %s", field, order)

	userListSQL, args, err := sg.Select("user.id, user.name, user.email, user.phone").
		From("user").
		Where(sg.Or{
			sg.Like{"user.name": fmt.Sprint("%", prefix, "%")},
			sg.Like{"user.id": fmt.Sprint("%", prefix, "%")},
		}).
		OrderBy(orderBy).
		Limit(pageSize).
		Offset(offset).
		ToSql()
	fmt.Println(args)
	fmt.Printf("userListSQL: %v\n", userListSQL)

	rows, err := db.DB.Queryx(userListSQL, fmt.Sprint("%", prefix, "%"), fmt.Sprint("%", prefix, "%"))

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("userListSQL: %s\n", user)

		userList = append(userList, user)
	}
	///	test := fmt.Sprintf("{data: %s}", userList)

	err = rows.Err()

	return userList, err
}

func GetUserById(id string) (e.User, error) {
	var user e.User

	userSQL, args, err := sg.Select("user.id, user.name, user.description, user.email, user.phone, (SELECT CAST(CONCAT('[',GROUP_CONCAT(JSON_OBJECT('label', `usher_group`.`name`, 'value', `usher_group`.`id`)),']') as JSON) FROM usher_group LEFT JOIN user_usher_group ON user_usher_group.usher_group = usher_group.id where user_usher_group.user = user.id) as usher_group").
		From("user").
		Where(sg.Eq{"user.id": id}).ToSql()

	fmt.Println(userSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(userSQL, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}

	}

	err = rows.Err()
	return user, err
}

func GetUsersByUsherGroupId(category string) ([]e.Gallery, error) {
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

func GetUserInfoByEmailAndPassword(user e.User) (e.User, error) {
	var userRes e.User

	userSQL, args, err := sg.Select("user.id, user.name, user.description, user.email, user.phone, (SELECT CAST(CONCAT('[',GROUP_CONCAT(JSON_OBJECT('usher_group_id', `usher_group`.`id`)),']') as JSON) FROM usher_group LEFT JOIN user_usher_group ON user_usher_group.usher_group = usher_group.id where user_usher_group.user = user.id) as usher_group").
		From("user").
		Where(sg.Eq{"user.email": user.Email, "user.password": user.Password}).ToSql()
	//		Where(sg.Eq{"user.id": userRes.ID}).ToSql()

	//	userSQL, args, err := sg.Select("user.id, user.name, user.description, user.email, user.phone").
	//		From("user").
	//		Where(sg.Eq{"user.email": user.Email, "user.password": user.Password}).ToSql()

	fmt.Println("userSQL")
	fmt.Println(userSQL)
	fmt.Println("args")
	fmt.Println(args)

	rows, err := db.DB.Queryx(userSQL, user.Email, user.Password)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&userRes)

		if err != nil {
			log.Fatalln(err)
		}

	}

	err = rows.Err()
	return userRes, err
}

func ResetPassword(user e.UserResetPassword) error {

	tx, err := db.DB.Begin()
	if err != nil {

		return err
	}
	resetQuery, args, err := sg.Update(`user`).
		Set(`password`, user.Password).
		Where(sg.Eq{"id": user.ID}).ToSql()

	fmt.Println(resetQuery)
	fmt.Println(args)

	res, err := tx.Exec(resetQuery, user.Password, user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return errors.New("User has not been affected")
	}

	if rowsAffected > 1 {
		tx.Rollback()
		return errors.New("Strange behaviour. Total affected is : " + string(rowsAffected))
	}

	tx.Commit()

	return nil
	/*
	   var user e.User

	   const query = `UPDATE user SET password = ? WHERE id = ?`

	   tx, err := db.DB.Begin()
	   if err != nil {

	   		return err
	   	}

	   res, err := tx.Exec(query, user.Password, user.ID)

	   	if err != nil {
	   		tx.Rollback()
	   		return err
	   	}

	   rowsAffected, err := res.RowsAffected()

	   	if err != nil {
	   		tx.Rollback()
	   		return err
	   	}

	   	if rowsAffected == 0 {
	   		tx.Rollback()
	   		return errors.New("User has not been affected")
	   	}

	   	if rowsAffected > 1 {
	   		tx.Rollback()
	   		return errors.New("Strange behaviour. Total affected is : " + string(rowsAffected))
	   	}

	   tx.Commit()

	   return nil
	*/
}

func GetUserByToken(id string) (e.User, error) {
	var user e.User

	userSQL, args, err := sg.Select("user.id, user.name").
		From("user").
		Where(sg.Eq{"user.password_token": id}).ToSql()

	fmt.Println(userSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(userSQL, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}
	}

	err = rows.Err()
	return user, err
}

func GetUserAuthByToken(id string) (e.UserAuth, error) {
	var user e.UserAuth

	userSQL, args, err := sg.Select("user.id, user.name, user.email, user.phone, schedule_week_user_token.week").
		From("user").
		InnerJoin("schedule_week_user_token ON schedule_week_user_token.user = user.id").
		Where(sg.Eq{"schedule_week_user_token.token": id}).ToSql()

	fmt.Println(userSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(userSQL, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}
	}

	err = rows.Err()
	return user, err
}
