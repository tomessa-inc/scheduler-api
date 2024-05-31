package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsherGroups(c echo.Context) error {

	pageIndex, err := strconv.ParseUint(c.Param("page-index"), 10, 64)
	pageSize, err := strconv.ParseUint(c.Param("page-size"), 10, 64)
	field := c.Param("field")
	order := c.Param("order")

	fmt.Printf("pageIndex %d\n", pageIndex)
	fmt.Printf("pageSize %d\n", pageSize)

	fmt.Printf("first\n\n")

	usherGroups, err := m.GetUsherGroups(pageIndex, pageSize, field, order)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find usher group listing")
	}

	usherGroupBytes, err := json.Marshal(usherGroups)
	usherGroupJson := ConvertStructToJSON(usherGroupBytes)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find gallery listing")
	}

	return c.JSON(http.StatusOK, usherGroupJson)
}

func GetUsherGroupsById(c echo.Context) error {
	id := c.Param("id")

	userGroup, err := m.GetUsherGroupById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find usher group")
	}

	return c.JSON(http.StatusOK, userGroup)
}

func GetUserGroupLabelValue(c echo.Context) error {
	var usherGroupKeyValueArray []e.UsherGroupKV
	var usherGroupKeyValue e.UsherGroupKV

	usherGroupList, err := m.GetUsherGroupsKV()

	//fmt.Printf("\n\n\n1%v1n\n\n", usherGroupList)

	//v := reflect.ValueOf(usherGroupList)

	for i := 0; i < len(usherGroupList); i++ {
		usherGroup := usherGroupList[i]

		test := map[string]interface{}{"label": usherGroup.Name, "value": usherGroup.ID}

		b, err := json.Marshal(test)
		json.Unmarshal([]byte(b), &usherGroupKeyValue)

		if err == nil {

		}

		usherGroupKeyValueArray = append(usherGroupKeyValueArray, usherGroupKeyValue)
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find usher group")
	}

	return c.JSON(http.StatusOK, usherGroupKeyValueArray)
}

func AddUsherGroup(c echo.Context) error {
	var usherGroup e.UsherGroup
	err := c.Bind(&usherGroup)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = m.AddUsherGroup(&usherGroup)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not create new usher group")
	}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))

	//return c.JSON(http.StatusOK, gallery)
}
