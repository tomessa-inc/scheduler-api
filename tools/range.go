package tools

import (
	"fmt"
	"scheduler-api/entity"
	"time"
)

func RangeConfiguration(jsonInterface map[string]interface{}) (entity.RangeConfig, error) {
	var rangeConfig entity.RangeConfig
	var err error

	rangeConfig.StartYear, err = GetIntDataFromJSONByKey(jsonInterface, "Range.start.year")
	rangeConfig.StartMonth, err = GetIntDataFromJSONByKey(jsonInterface, "Range.start.month")
	rangeConfig.StartDay, err = GetIntDataFromJSONByKey(jsonInterface, "Range.start.day")
	rangeConfig.EndDay, err = GetIntDataFromJSONByKey(jsonInterface, "Range.end.day")
	rangeConfig.EndYear, err = GetIntDataFromJSONByKey(jsonInterface, "Range.end.year")
	rangeConfig.EndMonth, err = GetIntDataFromJSONByKey(jsonInterface, "Range.end.month")

	fmt.Println("start year")
	fmt.Println(rangeConfig.StartYear)
	return rangeConfig, err
}

func DaysInMonth(m int, year int) int {

	return time.Date(year, time.Month(m)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func DaysOfWeek(m int, d int, year int) time.Weekday {

	return time.Date(year, time.Month(m)+1, d, 0, 0, 0, 0, time.UTC).Weekday()
}

func DaysConfig(rangeConfig entity.RangeConfig, y int, i int) entity.DaysConfig {
	var dayToCheck entity.DaysConfig

	fmt.Println("dagsConfig")
	fmt.Println(rangeConfig)
	if i == rangeConfig.StartMonth {
		dayToCheck.DaysToCheck = rangeConfig.StartDay
	} else {
		dayToCheck.DaysToCheck = 1
	}
	if i == rangeConfig.EndMonth {
		dayToCheck.Days = rangeConfig.EndDay
	} else {
		dayToCheck.Days = DaysInMonth(i, y)
	}

	return dayToCheck
}

/*
func buildRange(c echo.Context, key string) {
	var dayToCheck int

	jsonInterface := GetJSONRawBody(c)

	keyData := GetStringArrayDataFromJSONByKey(jsonInterface, key)
	rangeConfig, err := RangeConfiguration(jsonInterface)

	for y := rangeConfig.StartYear; y <= rangeConfig.EndYear; y++ {
		for i := rangeConfig.StartMonth; i <= rangeConfig.EndMonth; i++ {
			dayToCheck := DaysConfig(rangeConfig, y, i)

			for dayCheck := dayToCheck.DaysToCheck; dayCheck <= dayToCheck.Days; dayCheck++ {
			peopleAmount = keyData.UsherAmount
			dayofWeekMass := DaysOfWeek(i, dayCheck, y)
			if usherGroupData.Day == strings.ToLower(dayofWeekMass.String()) {
				fmt.Printf("\n\n\nstart day\n\n\n")
				week.Day = dayCheck
				week.Hour = usherGroupData.Hour
				week.Minute = usherGroupData.Minute
				week.Month = i
				week.Year = y
				week.UsherGroup = usherGroupData.ID
				weekId, err := m.AddWeek(&week)
				userUsherGroup.UsherGroup = usherGroupData.ID
				userUsherGroup.Number = 0
				usersInUsherGroup, err := m.GetUserUsherGroupByUsherGroup(userUsherGroup)

				if err != nil {

				}

				if len(usersInUsherGroup) < peopleAmount {
					fmt.Printf("\n\n\nless then user amount\n\n\n")

					peopleAmount = usherGroupData.UsherAmount - len(usersInUsherGroup)

					unavailable.UsherGroup = usherGroupData.ID
					m.RemoveUnAvailable(&unavailable)

					for last := 0; last < len(usersInUsherGroup); last++ {
						schedule.UserUsherGroup = usersInUsherGroup[last].ID
						schedule.Week = weekId
						m.AddSchedule(&schedule)
						unavailable.UserUsherGroup = usersInUsherGroup[last].ID
						unavailable.UsherGroup = usherGroupData.ID
						m.AddUnAvailable(&unavailable)
					}

					usersInUsherGroup, err = m.GetUserUsherGroupByUsherGroup(userUsherGroup)
					if err != nil {

					}
					/*
						unavailable.UsherGroup = usherGroupData.ID
						m.RemoveUnAvailable(&unavailable)
						number := uint64(usherGroupData.UsherAmount - len(usersInUsherGroup))
						fmt.Printf("\n\n\nthe number: %d", number)
						userUsherGroup.Number = number
						userUsherGroup.UsherGroup = usherGroupData.ID
						usersInUsherGroupTemp, err := m.GetUserUsherGroupByUsherGroup(userUsherGroup)
						if err != nil {
							panic(err)
						}
						usersInUsherGroup = append(usersInUsherGroup, usersInUsherGroupTemp...)
						//							fmt.Printf("made it here")

						//							os.Exit(3)
*/

/*
				for people := 0; people < peopleAmount; people++ {
					fmt.Printf("\n\n\nstart people\n\n\n")

					fmt.Printf("the list %v", usersInUsherGroup)
					rand.Seed(time.Now().Unix())
					n := rand.Intn(len(usersInUsherGroup))
					schedule.UserUsherGroup = usersInUsherGroup[n].ID
					schedule.Week = weekId
					err := m.AddSchedule(&schedule)
					unavailable.UserUsherGroup = usersInUsherGroup[n].ID
					unavailable.UsherGroup = usherGroupData.ID

					m.AddUnAvailable(&unavailable)
					if err != nil {
						fmt.Printf("error add achedule: %s:", err)
					}

					fmt.Printf("\n\n\n before slice: %v\n\n\n", usersInUsherGroup)
					//	slice := []int{1, 2, 3, 4}
					fmt.Printf("\n\n\nthe number: %d\n\n", n)
					fmt.Printf("\n\n\nthe number remaining: %d\n\n", len(usersInUsherGroup))
					usersInUsherGroup, err = remove(usersInUsherGroup, n)

					fmt.Printf("\n\n\n after slice: %v\n\n\n", usersInUsherGroup)

					//	fmt.Println(slice) // [1 3 4]
					//	RemoveIndex(usersInUsherGroup, n)
					fmt.Printf("\n\n\n end people\n\n\n")
				}
			}
		}
		//	fmt.Printf("made it here")
		//	os.Exit(3)

	}

}
*/
