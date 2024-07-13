package conversion

import (
	"encoding/json"
	"fmt"
	"io"
	e "scheduler-api/entity"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

var (
	EmptyValue = make([]int, 0)
)

func ConvertArrayBytesToString(bytesData []byte) (map[string]string, error) {

	//	fmt.Printf("json stirng data marsh string hwerw: %s\n", stringData)
	//	fmt.Printf("json stirng data marsh string raw: %s\n", stringData.Raw)
	//stringData := map[string]string([]byte(bytesData))
	makeString := make(map[string]string)
	err := json.Unmarshal(bytesData, &makeString)
	//	intData, err := strconv.Atoi(stringData)
	//intNumberMonth, err := strconv.Atoi(startMonth)
	//	fmt.Printf("json stirng data marsh int: %d\n", intData)
	return makeString, err
}

func ConvertGibsonBytesToInt(stringData gjson.Result) (int, error) {
	//	fmt.Printf("json stirng data marsh string hwerw: %s\n", stringData)
	//	fmt.Printf("json stirng data marsh string raw: %s\n", stringData.Raw)
	intData, err := strconv.Atoi(string([]byte(stringData.Raw)))
	//	intData, err := strconv.Atoi(stringData)
	//intNumberMonth, err := strconv.Atoi(startMonth)
	//	fmt.Printf("json stirng data marsh int: %d\n", intData)
	return intData, err
}

func ConvertJSONToString(jsonInterface interface{}) (string, error) {
	jsonBytes, err := json.Marshal(jsonInterface)
	stringData := string([]byte(jsonBytes))
	//valueInt := ConvertGibsonBytesToString(jsonBytes)
	//	fmt.Printf("json stirng data marsh int: %d\n", valueInt)
	return stringData, err
	/*
		if err := json.Unmarshal([]byte(str), &i); err != nil {
			fmt.Println("ugh: ", err)
		}

		fmt.Println("info: ", i)
		fmt.Println("currency: ", i.Data.Currency)

		//	fmt.Printf("json stirng data marsh string hwerw: %s\n", stringData)
		//	fmt.Printf("json stirng data marsh string raw: %s\n", stringData.Raw)
		intData, err := strconv.Atoi(string([]byte(stringData.Raw)))
		//	intData, err := strconv.Atoi(stringData)
		//intNumberMonth, err := strconv.Atoi(startMonth)
		//	fmt.Printf("json stirng data marsh int: %d\n", intData)
		return intData, err */
}

func ConvertGibsonBytesToString(bytesData gjson.Result) string {
	//	fmt.Printf("json stirng data marsh string hwerw: %s\n", stringData)
	//	fmt.Printf("json stirng data marsh string raw: %s\n", stringData.Raw)
	stringData := string([]byte(bytesData.Raw))
	//	intData, err := strconv.Atoi(stringData)
	//intNumberMonth, err := strconv.Atoi(startMonth)
	//	fmt.Printf("json stirng data marsh int: %d\n", intData)
	return stringData
}

func ConvertBytesToStringArray(bytesData gjson.Result) []string {
	//var testtt []e.UsherGroupKV
	var stringArray []string
	//	fmt.Printf("json stirng data marsh string hwerw: %s\n", stringData)
	//	fmt.Printf("json stirng data marsh string raw: %s\n", stringData.Raw)

	//	bytesData.IsArray()
	fmt.Printf("bytes  %d", bytesData.IsArray())
	//fmt.Printf("bytes  %d", bytesData.ForEach())

	bytesData.ForEach(func(key, value gjson.Result) bool {
		stringArray = append(stringArray, value.String())
		//		println(value.String())
		return true // keep iterating
	})

	return stringArray
}

func ConvertByteToJSON(bytesData []byte) map[string]interface{} {
	var jsonMap map[string]interface{}

	stringData := string(bytesData[:])

	stringDataReady := fmt.Sprintf("{\"data\": %s}", stringData)

	json.Unmarshal([]byte(stringDataReady), &jsonMap)

	return jsonMap

}

func ConvertIntToString(number int) string {
	intNumberMonth := strconv.Itoa(number)

	return intNumberMonth
}

func GetStringArrayDataFromJSONByKey(jsonInterface map[string]interface{}, key string) []string {
	//	fmt.Printf("json stirng data interface: %s\n", jsonInterface)
	//	fmt.Printf("json stirng key: %s\n", key)

	jsonBytes := ApplyMarshal(jsonInterface)
	//	fmt.Printf("json stirng data marsh: %s\n", jsonBytes)

	valueBytes := gjson.GetBytes(jsonBytes, key)
	fmt.Printf("number starting month woot: %s\n", valueBytes)
	//fmt.Printf("number starting month num: %d\n", len(valueBytes))

	var valueString []string = ConvertBytesToStringArray(valueBytes)
	//fmt.Printf("json stirng data array string: %v\n", valueString)

	return valueString
}

func GetJSONRawBody2(c echo.Context) {
	fmt.Println("start of addweek2")
	my_data := echo.Map{}

	if err := c.Bind(&my_data); err != nil {
		//	return err
	} else {

		start := fmt.Sprintf("%v", my_data["start"])
		fmt.Printf("jsonDatajson data: %s\n", start)
		//		useremail := fmt.Sprintf("%v", my_data["end"])
	}

	//	fmt.Printf("jsonDatajson data: %s\n", start)
	/*	var week e.CreateWeek
		//jsonBody := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&week)
		if err != nil {

			log.Error("empty json body")
			return nil
		}
		fmt.Printf("jsonDatajson data: %s\n", week)
		//fmt.Printf("end") */
	//return err
}

func UnwantedJSONHandler(c echo.Context) error {
	b, _ := io.ReadAll(c.Request().Body)
	fmt.Printf("dispay b: %v\n", b)
	var week e.CreateWeek
	answer := json.Unmarshal(b, &week)

	if answer != nil {
		fmt.Printf("jsonData to display: %s\n", &answer)
		return answer
	}
	//	log.Println(err.Error())
	return answer
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func DifferenceLV(changes []e.LV, source []e.UserUsherGroup) []string {
	var sourceCheck bool = true
	var diff []string
	//fmt.Printf("check out changes  %v", changes)

	//	fmt.Printf("check out source %v", source)
	fmt.Printf("check out number sources value %d\n", len(source))

	for i := 0; i < len(source); i++ {
		fmt.Printf("begin of outter loop\n")
		//		fmt.Printf("check out sources value %s", source[i].UsherGroup)
		//	fmt.Printf("check out chanages value %s\n", changes[i].Value)
		sourceCheck = true
		fmt.Printf("check out number changes  value %d\n", len(changes))
		for j := 0; j < len(changes); j++ {
			fmt.Printf("begin of inner loop\n")
			//sourceCheck := source[j].UsherGroup
			fmt.Printf("check out chanages value %s\n", changes[j].Value)
			fmt.Printf("check out sources value %s\n", source[i].UsherGroup)
			if changes[j].Value == source[i].UsherGroup {

				sourceCheck = false
			}
			//		fmt.Printf("check out sources value %s", changes[j].Value)
			//	err2 := m.AddUserUsherGroup(user.ID, userLV[i].Value)
			//	fmt.Println(err2)
			fmt.Printf("end of loop\n")
		}
		if sourceCheck {
			fmt.Printf("check out sources value %s\n", source[i].UsherGroup)
			diff = append(diff, source[i].UsherGroup)
		}

	}

	return diff

}
