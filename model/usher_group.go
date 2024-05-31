package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"
	"strings"

	sg "github.com/Masterminds/squirrel"
)

func AddUsherGroup(usherGroup *e.UsherGroup) error {
	//now := time.Now()
	id := strings.ToLower(usherGroup.Name)
	id = strings.Trim(id, " ")
	id = strings.Replace(id, " ", "-", -1)
	fmt.Printf("id2 valueL: %s\n", id)

	const query = `INSERT INTO usher_group (id, name, description, day, hour, minute, usher_amount) VALUES (?, ?, ?, ?, ?, ?, ?)`
	tx, err := db.DB.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(query, id, usherGroup.Name, usherGroup.Description, usherGroup.Day, usherGroup.Hour, usherGroup.Minute, usherGroup.UsherAmount)
	fmt.Printf("err valueL: %s\n", err)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func GetUsherGroupsKV() ([]e.UsherGroup, error) {

	var usherGroupList []e.UsherGroup
	usherGroup := e.UsherGroup{}

	usherGroupSQL, args, err := sg.Select("usher_group.id, usher_group.name, usher_group.description, usher_group.hour, usher_group.minute, usher_group.day, usher_group.usher_amount").
		From("usher_group").
		ToSql()

	if args == nil {

	}

	rows, err := db.DB.Queryx(usherGroupSQL)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&usherGroup)

		if err != nil {
			fmt.Println("error with  scan")

			log.Fatalln(err)
		}

		usherGroupList = append(usherGroupList, usherGroup)
	}

	err = rows.Err()
	return usherGroupList, err
}

func GetUsherGroups(pageIndex uint64, pageSize uint64, field string, order string) ([]e.UsherGroup, error) {

	var usherGroupList []e.UsherGroup
	usherGroup := e.UsherGroup{}
	offset := ((pageIndex - 1) * pageSize)

	orderBy := fmt.Sprintf("%s %s", field, order)

	usherGroupSQL, args, err := sg.Select("usher_group.id, usher_group.name, usher_group.description, usher_group.hour, usher_group.minute, usher_group.day, usher_group.usher_amount").
		From("usher_group").
		OrderBy(orderBy).
		Limit(pageSize).
		Offset(offset).
		ToSql()

	fmt.Println(usherGroupSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(usherGroupSQL)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&usherGroup)

		if err != nil {
			fmt.Println("erri with  scan")

			log.Fatalln(err)
		}

		usherGroupList = append(usherGroupList, usherGroup)
	}

	err = rows.Err()
	return usherGroupList, err
}

func GetUsherGroupById(id string) (e.UsherGroup, error) {
	fmt.Println(id)
	usherGroup := e.UsherGroup{}
	usherGroupSQL, args, err := sg.Select("usher_group.id, usher_group.name, usher_group.description, usher_group.day, usher_group.day, usher_group.hour, usher_group.minute, usher_group.usher_amount").
		From("usher_group").
		Where(sg.Eq{"id": id}).ToSql()

	fmt.Println(usherGroupSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(usherGroupSQL, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&usherGroup)

		if err != nil {
			log.Fatalln(err)
		}
	}

	err = rows.Err()

	return usherGroup, err
}
