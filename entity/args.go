package entity

type JsonType struct {
	Array []string
}
type LV struct {
	Label string
	Value string
}

type JsonColumn[T any] struct {
	v *T
}

type RangeConfig struct {
	StartYear  int
	StartMonth int
	StartDay   int
	EndYear    int
	EndMonth   int
	EndDay     int
}

type DaysConfig struct {
	DaysToCheck int
	Days        int
}

//startYear, err := tools.GetIntDataFromJSONByKey(jsonInterface, "Range.start.year")
//startMonth, err := tools.GetIntDataFromJSONByKey(jsonInterface, "Range.start.month")
//startDay, err := tools.GetIntDataFromJSONByKey(jsonInterface, "Range.start.day")
//endDay, err := tools.GetIntDataFromJSONByKey(jsonInterface, "Range.end.day")
//endYear, err := tools.GetIntDataFromJSONByKey(jsonInterface, "Range.end.year")
//endMonth, err := tools.GetIntDataFromJSONByKey(jsonInterface, "Range.end.month")
