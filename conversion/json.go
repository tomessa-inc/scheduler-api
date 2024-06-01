package conversion

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
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

func GetIntDataFromJSONByKey(jsonInterface map[string]interface{}, key string) (int, error) {
	//	fmt.Printf("json stirng data interface: %s\n", jsonInterface)
	//	fmt.Printf("json stirng key: %s\n", key)

	jsonBytes := ApplyMarshal(jsonInterface)
	//	fmt.Printf("json stirng data marsh: %s\n", jsonBytes)

	valueBytes := gjson.GetBytes(jsonBytes, key)
	//	fmt.Printf("number starting month: %s\n", valueBytes)
	valueInt, err := ConvertGibsonBytesToInt(valueBytes)
	//	fmt.Printf("json stirng data marsh int: %d\n", valueInt)
	return valueInt, err
}

func GetStringDataFromJSONByKey(jsonInterface map[string]interface{}, key string) string {
	//	fmt.Printf("json stirng data interface: %s\n", jsonInterface)
	//	fmt.Printf("json stirng key: %s\n", key)

	jsonBytes := ApplyMarshal(jsonInterface)
	//	fmt.Printf("json stirng data marsh: %s\n", jsonBytes)

	valueBytes := gjson.GetBytes(jsonBytes, key)
	//	fmt.Printf("number starting month: %s\n", valueBytes)
	valueString := ConvertGibsonBytesToString(valueBytes)
	//	fmt.Printf("json stirng data marsh int: %d\n", valueInt)
	return valueString
}
