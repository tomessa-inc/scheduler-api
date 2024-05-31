package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"scheduler-api/conversion"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetSchedule(c echo.Context) error {
	pageIndex, err := strconv.ParseUint(c.Param("page-index"), 10, 64)
	pageSize, err := strconv.ParseUint(c.Param("page-size"), 10, 64)
	field := c.Param("field")
	order := c.Param("order")
	var schedule e.GetSchedule
	err2 := c.Bind(&schedule)

	if err != nil {

	}
	if err2 != nil {

	}
	//usherGroups, err := m.GetUsherGroups(pageIndex, pageSize, field, order)

	fmt.Println("dodo")
	fmt.Println(schedule.RequestId)

	list, err := m.GetSchedule(&schedule, pageIndex, pageSize, field, order)

	usherGroupBytes, err := json.Marshal(list)
	usherGroupJson := conversion.ConvertStructToJSON(usherGroupBytes)

	return c.JSON(http.StatusOK, usherGroupJson)

}

func SetUnAvailable(c echo.Context) error {
	var schedule e.SetUnAvailable
	err := c.Bind(&schedule)
	var test e.ScheduleUser
	if err != nil {

	}

	id := schedule.ID
	fmt.Println("the id")
	fmt.Println(id)

	users, err := m.GetUsersByScheduleId(id)

	//scheduleid + week.id+nameiD

	for i := 0; i < len(users); i++ {
		test = users[i]
		fmt.Println("the users")

		fmt.Println(test.ScheduleId)
		fmt.Println(test.UserId)
		fmt.Println(test.Email)
		fmt.Println(test.Name)
		fmt.Println(test.Week)
		token, err := m.InsertStuff(test)

		if err != nil {

		}
		fmt.Println("the token")
		fmt.Println(token)

		//	m.SetUnAvailable(&schedule)
		m.PrepareEmail(test, "unavaialble", token)
	}

	//	userIds := funk.Map(users, func(u e.User) string {

	//		return u.Id
	//	}).([]string)

	///var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	// This should be in an env file in production

	// Hashing the password with the default cost of 10

	fmt.Println(err) // nil means it is a match

	//err2 := m.SetUnAvailable(&schedule)

	//m.GetUserUsherGroupByUser()

	if err != nil {

	}

	//m.PrepareEmail(users)

	//	usherGroupBytes, err := json.Marshal(list)
	//	usherGroupJson := ConvertStructToJSON(usherGroupBytes)
	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//	return c.JSON(http.StatusOK, usherGroupJson)

}

func SetAvailable(c echo.Context) error {
	var available e.SetAvailable
	err := c.Bind(&available)

	id, err := m.SetAvailable(&available)

	var test e.ScheduleUser
	if err != nil {

	}

	users, err := m.GetUsersByScheduleId(id)

	//scheduleid + week.id+nameiD

	for i := 0; i < len(users); i++ {
		test = users[i]
		fmt.Println("the users")

		fmt.Println(test.ScheduleId)
		fmt.Println(test.UserId)
		fmt.Println(test.Email)
		fmt.Println(test.Name)
		fmt.Println(test.Week)

		//	m.SetUnAvailable(&schedule)
		m.PrepareEmail(test, "available")
	}

	//	userIds := funk.Map(users, func(u e.User) string {

	//		return u.Id
	//	}).([]string)

	///var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
	// This should be in an env file in production

	// Hashing the password with the default cost of 10

	fmt.Println(err) // nil means it is a match

	//err2 := m.SetUnAvailable(&schedule)

	//m.GetUserUsherGroupByUser()

	if err != nil {

	}

	//m.PrepareEmail(users)

	//	usherGroupBytes, err := json.Marshal(list)
	//	usherGroupJson := ConvertStructToJSON(usherGroupBytes)
	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//	return c.JSON(http.StatusOK, usherGroupJson)

}
