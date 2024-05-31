package error

import "net/http"

func StatusUnprocessableEntity(err) {
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
}

///if err != nil {
//		return c.JSON(http.StatusUnprocessableEntity, err.Error())
//	}
