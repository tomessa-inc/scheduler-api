package tools

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/tidwall/gjson"
	"io"
	e "scheduler-api/entity"
	"strings"
)

func GetJSONRawBody(reader io.Reader) map[string]interface{} {
	var absence e.Absence
	var err error
	buf2 := new(strings.Builder)
	fmt.Println("arrivbe11")
	fmt.Println(reader)
	n, err := io.Copy(buf2, reader)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println("the body2ins")
	fmt.Println(buf2.String())

	jsonBody := make(map[string]interface{})
	fmt.Println("the body")
	fmt.Println(reader)
	//	buf := new(strings.Builder)
	io.Copy(buf2, reader)
	fmt.Println("the body2")
	fmt.Println(buf2.String())
	fmt.Println("the body3")
	fmt.Println([]byte(buf2.String()))
	json.Unmarshal([]byte(buf2.String()), &absence)
	fmt.Println("yo here")
	fmt.Println(absence)

	fmt.Println("yo herytte")
	fmt.Println(absence)

	err = json.NewDecoder(reader).Decode(&jsonBody)
	if err != nil {

		log.Error("empty json body")
		return nil
	}

	fmt.Println("checking jsonody")
	fmt.Println(jsonBody)
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
