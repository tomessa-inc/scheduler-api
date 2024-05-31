package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"scheduler-api/conversion"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	//	fmt.Printf("echo value: %v\n", c)
	var user e.User
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = m.AddUser(&user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user listing")
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))
}

func UpdateUser(c echo.Context) error {

	var user e.User
	var userLV []e.LV
	//var userUsherGroup e.UserUsherGroup

	err := c.Bind(&user)

	if err != nil {

	}
	err = m.UpdateUser(&user)

	userUsherGroup, err := m.GetUserUsherGroupByUser(user.ID)

	//fmt.Printf("check out err var %s", err)

	var usherGroup string = *user.UsherGroup

	json.Unmarshal([]byte(usherGroup), &userLV)

	//	out, err := json.Marshal(userLV)
	diff := conversion.differenceLV(userLV, userUsherGroup)

	for i := 0; i < len(diff); i++ {
		err2 := m.DeleteUserUsherGroupByUser(user.ID, diff[i])
		fmt.Println(err2)
	}

	//fmt.Printf("check out diff  %s", diff)

	for i := 0; i < len(userLV); i++ {
		err2 := m.AddUserUsherGroup(user.ID, userLV[i].Value)
		fmt.Println(err2)
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))
}

func getMyString(items []LV) (string, error) {

	var buffer bytes.Buffer
	var err error
	var b []byte

	for _, item := range items {
		b, err = json.Marshal(item)
		if err != nil {
			return "", err
		}

		buffer.WriteString(string(b) + ",")
	}

	s := strings.TrimSpace(buffer.String())
	// trim last comma
	s = s[:len(s)-1]

	return s, nil
}

func GetUsers(c echo.Context) error {
	pageIndex, err := strconv.ParseUint(c.Param("page-index"), 10, 64)
	pageSize, err := strconv.ParseUint(c.Param("page-size"), 10, 64)
	field := c.Param("field")
	order := c.Param("order")

	users, err := m.GetUsers(pageIndex, pageSize, field, order)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user listing")
	}

	usersBytes, err := json.Marshal(users)
	uersJson := ConvertStructToJSON(usersBytes)

	return c.JSON(http.StatusOK, uersJson)
}

func GetUsersByPrefix(c echo.Context) error {
	pageIndex, err := strconv.ParseUint(c.Param("page-index"), 10, 64)
	pageSize, err := strconv.ParseUint(c.Param("page-size"), 10, 64)
	field := c.Param("field")
	order := c.Param("order")
	prefix := c.Param("prefix")

	users, err := m.GetUsersByPrefix(pageIndex, pageSize, field, order, prefix)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user listing")
	}

	usersBytes, err := json.Marshal(users)
	uersJson := ConvertStructToJSON(usersBytes)

	return c.JSON(http.StatusOK, uersJson)
}

func GetUserById(c echo.Context) error {

	userId := c.Param("id")
	user, err := m.GetUserById(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user")
	}
	//json.Marshal()
	//user.UsherGroup = json.Unmarshal(user.UsherGroup)
	return c.JSON(http.StatusOK, user)
}

func GetUserByToken(c echo.Context) error {
	fmt.Printf("echo value: %v\n", c)
	//	var user e.User
	//	err := c.Bind(&user)
	token := c.Param("token")

	//i

	user, err := m.GetUserByToken(token)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserAuthByToken(c echo.Context) error {
	fmt.Printf("echo value: %v\n", c)
	//	var user e.User
	//	err := c.Bind(&user)
	token := c.Param("token")

	//i

	user, err := m.GetUserAuthByToken(token)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	tokenGenerated := generateToken(user.Email)

	retval := map[string]interface{}{"user": user, "token": tokenGenerated}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", retval))
	// return c.JSON(http.StatusOK, user)
}

func ResetPassword(c echo.Context) error {
	fmt.Printf("echo value: %v\n", c)
	var user e.UserResetPassword
	err := c.Bind(&user)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err2 := m.ResetPassword(user)

	if err2 != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
