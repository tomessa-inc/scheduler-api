package controller

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"

	//	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	//"scheduler-api/tools"
	"strings"
)

func SetAbsence(c echo.Context) error {
	//var rangeConfig e.RangeConfig
	//	var err error
	var absence e.Absence
	var userUserGroup []e.UserUsherGroup
	//	var usherGroup e.UsherGroup

	buf := new(strings.Builder)
	io.Copy(buf, c.Request().Body)
	json.Unmarshal([]byte(buf.String()), &absence)
	c.Bind(&absence)

	fmt.Println("body")
	fmt.Println(absence)
	fmt.Println("ID")
	fmt.Println("RANGE")

	fmt.Println(absence.ID)
	fmt.Println(absence.Range)

	fmt.Println("yo here")
	fmt.Println(absence)

	/*
		//	var absence e.Absence

		//	jsonBody := make(map[string]interface{})
		fmt.Println("the body")
		fmt.Println(c.Request().Body)
		buf := new(strings.Builder)
		n, err := io.Copy(buf, c.Request().Body)
		fmt.Println(n)
		fmt.Println("the body2")
		fmt.Println(buf.String())
		fmt.Println("the body3")
		fmt.Println([]byte(buf.String()))
		json.Unmarshal([]byte(buf.String()), &absence)
		fmt.Println("yo here")
		fmt.Println(absence)
		c.Bind(&absence)
		fmt.Println("yo herytte")
		fmt.Println(absence)
	*/
	//	pageIndex, err := strconv.ParseUint(c.Param("page-index"), 10, 64)
	//	pageSize, err := strconv.ParseUint(c.Param("page-size"), 10, 64)
	//field := c.Param("field")
	//	fmt.Println("the c")//
	//	fmt.Println(c)
	//	fmt.Println("arrivbe11")
	//	fmt.Println(c.Request().Body)/
	//	n, err := io.Copy(buf, c.Request().Body)
	//	fmt.Println(n)
	//	fmt.Println(err)
	//	fmt.Println("the body2ins")
	//	fmt.Println(buf.String())

	//	order := c.Param("order")
	//	err = c.Bind(&absence)
	fmt.Println(absence.ID)
	//	userUserGroup, err = m.GetUserUsherGroupByUser(absence.ID)
	fmt.Println(userUserGroup)

	//jsonInterface := GetJSONRawBody2(c)
	fmt.Println("arrivbe22")
	/*	rangeConfig, err = tools.RangeConfiguration(jsonInterface)

		if err == nil {
			return nil
		}
		fmt.Println("arrivbed")
		for y := rangeConfig.StartYear; y <= rangeConfig.EndYear; y++ {
			fmt.Println("arrivbed1")
			for i := rangeConfig.StartMonth; i <= rangeConfig.EndMonth; i++ {
				fmt.Println("arrivbed2")
				dayToCheck := tools.DaysConfig(rangeConfig, y, i)
				fmt.Println("days to check")
				fmt.Println(dayToCheck)
				fmt.Println(dayToCheck.Days)
				for dayCheck := dayToCheck.DaysToCheck; dayCheck <= dayToCheck.Days; dayCheck++ {
					fmt.Println("arrivbed3")
					fmt.Println(len(userUserGroup))
					for usherGroupCount := 0; usherGroupCount < len(userUserGroup); usherGroupCount++ {

						fmt.Println("arrivbed")
						dayofWeekMass := tools.DaysOfWeek(i, dayCheck, y)
						fmt.Println(dayofWeekMass)
						fmt.Println(userUserGroup)
						fmt.Println(userUserGroup[usherGroupCount])
						fmt.Println(userUserGroup[usherGroupCount].UsherGroup)
						usherGroup, err = m.GetUsherGroupById(userUserGroup[usherGroupCount].UsherGroup)
						fmt.Println("the day")
						fmt.Println(dayofWeekMass.String())
						fmt.Println("the day check")
						fmt.Println(usherGroup.Day)

						if usherGroup.Day == strings.ToLower(dayofWeekMass.String()) {
							fmt.Println("we have arrived")
						}
						//					fmt.Printf("\n\n\nstart day\n\n\n")
						/*					absense.ID = ""
											absense.
											//					absense
											//					absense.Range					week.Day = d
											week.Hour = usherGroupData.Hour
											week.Minute = usherGroupData.Minute
											week.Month = i
											week.Year = y
											week.UsherGroup = usherGroupData.ID
											weekId, err := m.AddWeek(&week) */
	//	}
	//	}

	//	}
	//}

	//	tools.buildRange(c, "UsherGroup")

	//tools.buildRange(c, "UsherGroup")

	//var absence e.Absence
	err2 := c.Bind(&absence)

	//	if err != nil {//

	//	}
	if err2 != nil {

	}

	m.SetUnAvaiable(absence)

	//usherGroups, err := m.GetUsherGroups(pageIndex, pageSize, field, order)

	//	fmt.Println("dodo")
	//	fmt.Println(schedule.RequestId)

	//list, err := m.GetSchedule(&schedule, pageIndex, pageSize, field, order)

	//usherGroupBytes, err := json.Marshal(list)
	//	usherGroupJson := tools.ConvertByteToJSON(usherGroupBytes)

	return c.JSON(http.StatusOK, "")

}

func GetJSONRawBody2(c echo.Context) map[string]interface{} {
	var absence e.Absence
	var err error
	buf2 := new(strings.Builder)
	fmt.Println("arrivbe11")
	fmt.Println(c.Request().Body)
	n, err := io.Copy(buf2, c.Request().Body)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println("the body2ins")
	fmt.Println(buf2.String())

	jsonBody := make(map[string]interface{})
	fmt.Println("the body")

	fmt.Println("the body2")
	fmt.Println(buf2.String())
	fmt.Println("the body3")
	fmt.Println([]byte(buf2.String()))
	json.Unmarshal([]byte(buf2.String()), &absence)
	fmt.Println("yo here")
	fmt.Println(absence)

	fmt.Println("yo herytte")
	fmt.Println(absence)

	err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	fmt.Println("checking jsonody")
	fmt.Println(jsonBody)
	return jsonBody
}

/*
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
*/
