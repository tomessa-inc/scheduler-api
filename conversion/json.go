package conversion

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetJSONRawBody(c echo.Context) map[string]interface{} {

	jsonBody := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	return jsonBody
}

func ApplyMarshal(jsonInterface map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(jsonInterface)
	if err != nil {
		return nil
	}

	return jsonBytes
}
