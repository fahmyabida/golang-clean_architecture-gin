package http

import (
	"context"
	"errors"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/gin-gonic/gin"
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
	e.GET("/pemesan", handler.GetPemesan)
}

func (a *httpPemesanHandler) GetPemesan(c *gin.Context){
	idParams := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idParams); if err!=nil {
		c.JSON(http.StatusBadRequest, utils.Error(errors.New("Bad Request"), gin.H{})); return}
	ctx := c.Request.Context(); if ctx == nil {ctx = context.Background()}
	pemesan, err := a.pUsecase.GetPemesan(ctx, id); if err != nil {
		c.JSON(http.StatusOK, utils.Error(err, gin.H{})); return
	}
	c.JSON(http.StatusOK, utils.ResponseObject(pemesan, gin.H{})); return
}
