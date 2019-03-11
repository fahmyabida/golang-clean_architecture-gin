package models

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Given Param is not valid")
)

func ErrorQuery(err error, result gin.H) gin.H{
	result = gin.H{
		"message": "error",
		"detail": err.Error(),
	}
	return result
}