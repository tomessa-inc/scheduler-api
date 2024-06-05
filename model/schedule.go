package model

import (
	"fmt"
	"log"
	"scheduler-api/db"
	e "scheduler-api/entity"
	"strings"

	"time"

	sg "github.com/Masterminds/squirrel"
)

func AddSchedule(schedule *e.Schedule) error {
	//now := time.Now()
	//week.ID = "sdsdfsdfsdf"
	//fmt.Printf("about to insert into db\n")
	id := fmt.Sprintf("%s-%s", schedule.Week, schedule.UserUsherGroup)
	createdAt := time.Now().Format("20060102150405")

	addIinsert, args, err := sg.Insert(`schedule`).
		Columns(`id`, `week`, `user_usher_group`, `created_at`).
		Values(id, schedule.Week, schedule.UserUsherGroup, createdAt).
		ToSql()

	fmt.Println(args)
	fmt.Println("time")
	fmt.Println(time.Now().Format("20060102150405"))
	tx, err := db.DB.Begin()

	//id := fmt.Sprintf("%s-%s", schedule.Week, schedule.UserUsherGroup)

	_, err = tx.Exec(addIinsert, id, schedule.Week, schedule.UserUsherGroup, createdAt)

	if err != nil {
		fmt.Printf("Error in SQL: %s\n", err)
		tx.Rollback()

		return err
	}

	tx.Commit()

	return err
}

func InsertStuff(schedule e.ScheduleUser) (string, error) {
	const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

	// To encrypt the StringToEncrypt

	const query = `INSERT INTO schedule_week_user_token (schedule, week, user, token) VALUES (?, ?, ?, ?)`
	//	fmt.Println("time")
	//	fmt.Println(time.Now().Format("20060102150405"))
	tx, err := db.DB.Begin()

	id := fmt.Sprintf("%s+%s+%s", schedule.ScheduleId, schedule.UserId, schedule.Week)

	fmt.Println("the id")
	fmt.Println(id)
	encText, err := Encrypt(id, MySecret)

	fmt.Printf("Variable schedule in SQL: %s\n", schedule.ScheduleId)
	fmt.Printf("Variable week in SQL: %s\n", schedule.Week)
	fmt.Printf("Variable userid in SQL: %s\n", schedule.UserId)

	fmt.Printf("Variable token in SQL: %s\n", encText)

	_, err = tx.Exec(query, schedule.ScheduleId, schedule.Week, schedule.UserId, encText)

	if err != nil {
		fmt.Printf("Error in SQL: %s\n", err)
		tx.Rollback()

		return "nil", err
	}

	tx.Commit()

	return encText, err
}

func SetUnAvailable(schedule *e.SetUnAvailable) error {
	deleteQuery, args, err := sg.Delete("").
		From("schedule").
		Where(sg.Eq{"id": schedule.ID}).
		ToSql()

	fmt.Println(deleteQuery)
	fmt.Println(args)

	tx, err := db.DB.Begin()

	_, err = tx.Exec(deleteQuery, schedule.ID)

	if err != nil {
		fmt.Printf("Error in SQL: %s\n", err)
		tx.Rollback()

		return err
	}

	tx.Commit()

	return err
}
func SetAvailable(schedule *e.SetAvailable) (string, error) {

	id := fmt.Sprintf("%s-%s", schedule.Week, schedule.UserUsherGroup)

	createdAt := time.Now().Format("20060102150405")

	insertQuery, args, err := sg.Insert(`schedule`).Columns(`id`, `week`, `user_usher_group`, `created_at`).
		Values(id, schedule.Week, schedule.UserUsherGroup, createdAt).
		ToSql()

	fmt.Println(insertQuery)
	fmt.Println(args)

	tx, err := db.DB.Begin()

	_, err = tx.Exec(insertQuery, id, schedule.Week, schedule.UserUsherGroup, createdAt)

	if err != nil {
		fmt.Printf("Error in SQL: %s\n", err)
		tx.Rollback()

		return "", err
	}

	tx.Commit()

	return id, err
}

func GetSchedule(schedule *e.GetSchedule, pageIndex uint64, pageSize uint64, field string, order string) ([]e.ScheduleList, error) {

	//	SELECT user_usher_group.*, usher_group.*, schedule.*
	//	FROM user_usher_group
	//	LEFT JOIN usher_group  ON user_usher_group.usher_group = usher_group.id
	//	LEFT JOIN schedule ON schedule.user_usher_group = user_usher_group.id
	//	where user_usher_group.user = 'tom-cruickshank'

	//	SELECT usher_group.name as mass,
	//	(SELECT CAST(CONCAT('[',GROUP_CONCAT(JSON_OBJECT('user', `user_inner`.`name`)),']') as JSON)
	//	FROM user as user_inner
	//	LEFT JOIN user_usher_group as user_usher_group_inner ON user_usher_group_inner.user = user_inner.id
	//	LEFT JOIN schedule as schedule_inner ON schedule_inner.user_usher_group =  user_usher_group_inner.id
	//	WHERE schedule_inner.week = schedule.week) as users, schedule.week
	//	FROM schedule
	//	LEFT JOIN user_usher_group ON user_usher_group.id = schedule.user_usher_group
	//	LEFT JOIN usher_group ON usher_group.id = user_usher_group.usher_group/
	//	GROUP BY users
	const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"

	offset := ((pageIndex - 1) * pageSize)

	orderBy := fmt.Sprintf("%s %s", field, order)

	var scheduleList []e.ScheduleList
	var sqlWhere sg.SelectBuilder
	var userInfo []interface{}
	schedule2 := e.ScheduleList{}
	//var test string

	if schedule.RequestId != "" {
		test, err := Decrypt(schedule.RequestId, MySecret)
		fmt.Println(err)
		fmt.Println("test")
		fmt.Println(test)

		check := strings.Split(test, "+")
		userInfo = append(userInfo, check[1])
	}

	//userInfo = schedule.UserId

	scheduleSQL := sg.Select("usher_group.name as mass, (SELECT CAST(CONCAT('[',GROUP_CONCAT(JSON_OBJECT('user', `user_inner`.`name`)),']') as JSON) FROM user as user_inner LEFT JOIN user_usher_group as user_usher_group_inner ON user_usher_group_inner.user = user_inner.id  LEFT JOIN schedule as schedule_inner ON schedule_inner.user_usher_group =  user_usher_group_inner.id 	WHERE schedule_inner.week = schedule.week) as user, schedule.week, schedule.id").
		From("schedule").
		LeftJoin("user_usher_group ON user_usher_group.id = schedule.user_usher_group").
		LeftJoin("usher_group ON usher_group.id = user_usher_group.usher_group").
		GroupBy("schedule.week").
		OrderBy(orderBy).
		Limit(pageSize).
		Offset(offset)

	switch schedule.Type {
	case "solo":
		userInfo = append(userInfo, schedule.UserId)

		sqlWhere = scheduleSQL.Where(sg.Eq{"user_usher_group.user": userInfo})
		break
	case "group":
		usherGroup, err := GetUserUsherGroupByUser(schedule.UserId)

		if err != nil {

		}
		for _, u := range usherGroup {
			userInfo = append(userInfo, u.UsherGroup)
		}
		sqlWhere = scheduleSQL.Where(sg.Eq{"user_usher_group.usher_group": userInfo})

		break
	default:
		sqlWhere = scheduleSQL
		break
	}
	//	sqlWhere = scheduleSQL.Where("user_usher_group.user = 'tom-cruickshank'")

	sql, args, err := sqlWhere.ToSql()
	//Where(sg.Eq{"user_usher_group.user": schedule.UserId}).ToSql()

	fmt.Println(sql)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(sql, userInfo...)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.StructScan(&schedule2)

		if err != nil {
			log.Fatalln(err)
		}

		scheduleList = append(scheduleList, schedule2)
	}

	err = rows.Err()

	return scheduleList, err
}

func GetUsersByScheduleId(id string) ([]e.ScheduleUser, error) {

	//userInfo = schedule.UserId
	var user e.ScheduleUser
	var userList []e.ScheduleUser

	scheduleSQL, args, err := sg.Select("schedule.id as scheduleId, schedule.week, user.id as userId, user.email, user.name").
		From("schedule").
		LeftJoin("user_usher_group ON schedule.user_usher_group = user_usher_group.id").
		LeftJoin("user_usher_group as uug ON user_usher_group.usher_group = uug.usher_group").
		LeftJoin("user ON user.id = uug.user").
		Where(sg.Eq{"schedule.id": id}).ToSql()

	fmt.Println(scheduleSQL)
	fmt.Println(args)

	//	fmt.Println("here")
	rows, err := db.DB.Queryx(scheduleSQL, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	err = rows.Err()

	for rows.Next() {
		err := rows.StructScan(&user)

		if err != nil {
			log.Fatalln(err)
		}

		userList = append(userList, user)
	}

	return userList, err
}
