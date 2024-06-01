package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func (j *JsonColumn[T]) Scan(src any) any {
	if src == nil {
		j.v = nil
		return nil
	}
	j.v = new(T)
	fmt.Printf("hello %s", src)
	fmt.Println(reflect.TypeOf(src))
	fmt.Printf("stuff %v", json.Unmarshal(src.([]byte), j.v))
	json.Unmarshal(src.([]byte), j.v)

	return src //json.Unmarshal(src.([]byte), j.v)
}

func (j *JsonColumn[T]) Value() (driver.Value, error) {
	raw, err := json.Marshal(j.v)
	return raw, err
}

func (j *JsonColumn[T]) Get() *T {
	return j.v
}

func (u *UsherGroups) Scan(value interface{}) any {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	//json.Unmarshal(b, &u)

	fmt.Printf("hello %v", value)
	return json.Unmarshal(b, &u)
}
