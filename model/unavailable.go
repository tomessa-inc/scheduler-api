package model

import (
	"fmt"
	"scheduler-api/db"
	e "scheduler-api/entity"
)

func AddUnAvailable(unavaialble *e.Unavaialble) error {
	//now := time.Now()
	//week.ID = "sdsdfsdfsdf"
	//fmt.Printf("about to insert into db\n")
	const query = `INSERT INTO unavailable (user_usher_group,usher_group) VALUES (?, ?)`
	tx, err := db.DB.Begin()

	//id := fmt.Sprintf("%d-%d-%d-%d-%d-%s", week.Hour, week.Minute, week.Day, week.Month, week.Year, week.UsherGroup)

	_, err = tx.Exec(query, unavaialble.UserUsherGroup, unavaialble.UsherGroup)

	if err != nil {
		fmt.Printf("Error in unavaialble  SQL: %s\n", err)
		tx.Rollback()

		return err
	}

	tx.Commit()

	return err
}

func RemoveUnAvailable(unavaialble *e.Unavaialble) error {
	//now := time.Now()
	//week.ID = "sdsdfsdfsdf"
	//fmt.Printf("about to insert into db\n")

	const query = `DELETE FROM unavailable where unavailable.usher_group = ?`
	tx, err := db.DB.Begin()

	//id := fmt.Sprintf("%d-%d-%d-%d-%d-%s", week.Hour, week.Minute, week.Day, week.Month, week.Year, week.UsherGroup)

	_, err = tx.Exec(query, unavaialble.UsherGroup)

	if err != nil {
		fmt.Printf("Error in unavaialble  SQL: %s\n", err)
		tx.Rollback()

		return err
	}

	tx.Commit()

	return err
}
