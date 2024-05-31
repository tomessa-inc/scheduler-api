package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"
	"time"

	sg "github.com/Masterminds/squirrel"
)

func AddWeek(week *e.Week) (string, error) {
	//now := time.Now()
	//week.ID = "sdsdfsdfsdf"
	//fmt.Printf("about to insert into db\n")
	const query = `INSERT INTO week (id, hour,minute, day, month, year, usher_group, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	tx, err := db.DB.Begin()

	id := fmt.Sprintf("%d-%d-%d-%d-%d-%s", week.Hour, week.Minute, week.Day, week.Month, week.Year, week.UsherGroup)

	_, err = tx.Exec(query, id, week.Hour, week.Minute, week.Day, week.Month, week.Year, week.UsherGroup, time.Now().Format("20060102150405"))

	if err != nil {
		fmt.Printf("Error in SQL: %s\n", err)
		tx.Rollback()

		return id, err
	}

	tx.Commit()

	return id, err
}

func GetDetails(id string) ([]e.UserAuth, error) {
	var user e.UserAuth
	var userList []e.UserAuth

	userSQL, args, err := sg.Select("user.id, user.name, user.email, user.phone").
		From("week").
		InnerJoin("schedule ON schedule.week = week.id").
		InnerJoin("user_usher_group as uug ON uug.id = schedule.user_usher_group").
		InnerJoin("user_usher_group ON uug.user = user_usher_group.user").
		InnerJoin("user ON user.id = user_usher_group.user").
		//	InnerJoin("user ON user.id = user_usher_group.user").
		Where(sg.Eq{"schedule.week": id}).ToSql()

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

		userList = append(userList, user)
	}

	err = rows.Err()
	return userList, err
}

func GetWeekDetails(id string) (e.Week, error) {
	var week e.Week

	weekSQL, args, err := sg.Select("week.id, week.hour, week.minute, week.day, week.month, week.year, week.usher_group, week.created_at").
		From("week").
		Where(sg.Eq{"week.id": id}).ToSql()

	fmt.Println(weekSQL)
	fmt.Println(args)

	rows, err := db.DB.Queryx(weekSQL, id)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&week)

		if err != nil {
			log.Fatalln(err)
		}

	}

	err = rows.Err()
	return week, err
}
