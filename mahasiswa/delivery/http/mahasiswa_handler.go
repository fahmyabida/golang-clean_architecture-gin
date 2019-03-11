package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/me/golang-clean_architecture-gin_gonic/mahasiswa"
	"github.com/me/golang-clean_architecture-gin_gonic/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

// HttpArticleHandler  represent the httphandler for article
type HttpMahasiswaHandler struct {
	MhsUsecase mahasiswa.Usecase
}

func NewMahasiswaHttpHandler(e *gin.Engine, mu mahasiswa.Usecase) {
	handler := &HttpMahasiswaHandler{
		MhsUsecase: mu,
	}
	e.GET("/mahasiswa", handler.GetByUUID)
}

func (a *HttpMahasiswaHandler) GetByUUID(c *gin.Context){
	stringUuidMhs := c.Request.URL.Query().Get("uuid")
	uuidMhs := uuid.MustParse(stringUuidMhs)
	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}
	mhs, err := a.MhsUsecase.GetByUUID(ctx, uuidMhs)
	if err != nil {
		c.JSON(getStatusCode(err), models.ErrorQuery(err, gin.H{}))
		return
	}
	c.JSON(http.StatusOK, mhs)
	return
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

