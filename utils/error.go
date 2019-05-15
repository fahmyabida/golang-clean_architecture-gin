package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case ErrInternalServerError:
		return http.StatusInternalServerError
	case ErrNotFound:
		return http.StatusNotFound
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func Error(err error, result gin.H) gin.H{
	emptyArray := make([]string,0)
	result = gin.H{
		"data": emptyArray,
		"detail": err.Error(),
	}
	return result
}