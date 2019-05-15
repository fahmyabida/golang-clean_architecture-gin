package http

import (
	"context"
	"errors"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/me/golang-clean_architecture-gin_gonic/models"
	"net/http"
	"strconv"
)

type httpPemesanHandler struct {
	pUsecase pemesan.Usecase
}

func NewPemesanHttpHandler(e *gin.Engine, pU pemesan.Usecase) {
	handler := &httpPemesanHandler{
		pUsecase	: pU,
	}
	e.GET("/pemesan", handler.GetById)
}

func (a *httpPemesanHandler) GetById(c *gin.Context){
	idParams := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idParams); if err!=nil {
		c.JSON(http.StatusBadRequest, utils.Error(errors.New("Bad Request"), gin.H{})); return}
	ctx := c.Request.Context(); if ctx == nil {ctx = context.Background()}
	pemesan, err := a.pUsecase.GetPemesan(ctx, id); if err != nil {
		c.JSON(http.StatusOK, models.ErrorQuery(err, gin.H{})); return
	}
	c.JSON(http.StatusOK, pemesan); return
}
