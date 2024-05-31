package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

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

type User struct {
	ID          string  `db:"id"`
	Name        string  `db:"name"`
	Email       string  `db:"email"`
	Description string  `db:"description"`
	Phone       string  `db:"phone"`
	UsherGroup  *string `db:"usher_group"`
	Password    string  `db:"password"`
}

type UserAuth struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Email       string `db:"email"`
	Description string `db:"description"`
	Phone       string `db:"phone"`
	//UsherGroups []UsherGroups `db:"usher_groups"`
}

type UserWrite struct {
	ID          string        `db:"id"`
	Name        string        `db:"name"`
	Email       string        `db:"email"`
	Description string        `db:"description"`
	Phone       string        `db:"phone"`
	UsherGroups []UsherGroups `db:"usher_groups"`
	// UsherGroup [string] `db:"usher_group"`
}

type UsherGroups struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type UserResetPassword struct {
	ID       string `db:"id"`
	Password string `db:"password"`
}
