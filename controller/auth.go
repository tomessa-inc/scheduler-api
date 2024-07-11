package controller

import (
	"fmt"
	"net/http"
	e "scheduler-api/entity"
	m "scheduler-api/model"

	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	EmptyValue = make([]int, 0)
)

func SignIn(c echo.Context) error {
	var user e.User

	fmt.Println("the c")
	fmt.Println(c)
	fmt.Println("the body")
	fmt.Println(c.Request().Body)

	fmt.Println("user")

	err := c.Bind(&user)
	fmt.Println(user)
	///if err != nil {
	//		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	//	}

	user, err = m.GetUserInfoByEmailAndPassword(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "could not find user listing")
	}

	token := generateToken(user.Email)
	/*
		token, errToken := createToken(user.Email)

		if errToken != nil {

		}

		errVerify := verifyToken(token)

		if errVerify != nil {

		}*/
	retval := map[string]interface{}{"user": user, "token": token}

	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", retval))

	//return c.JSON(http.StatusOK, gallery)
}

func SignOut(c echo.Context) error {
	return c.JSON(http.StatusCreated, e.SetResponse(http.StatusCreated, "ok", EmptyValue))
}

var secretKey = []byte("secret-key")

func generateToken(email string) string {
	token, errToken := createToken(email)

	if errToken != nil {

	}

	errVerify := verifyToken(token)

	if errVerify != nil {

	}

	return token
}

func createToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
