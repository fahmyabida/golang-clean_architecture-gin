package http

import (
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type httpInvoiceHandler struct {
	iUsecase invoice.Usecase
}

func NewInvoiceHttpHandler(e *gin.Engine, iU invoice.Usecase) {
	handler := &httpInvoiceHandler{
		iUsecase	: iU,
	}
	e.GET("/invoice-menu", handler.GetByIdOrder)
}


func (h *httpInvoiceHandler) GetByIdOrder(c *gin.Context){
	c.Request.Body.
	idParams := c.Request.URL.Query().Get("id")
	id, err := strconv.Atoi(idParams); if err!=nil {
		c.JSON(http.StatusBadRequest, utils.Error(errors.New("Bad Request"), gin.H{})); return}
	ctx := c.Request.Context(); if ctx == nil {ctx = context.Background()}
	pemesan, err := a.pUsecase.GetPemesan(ctx, id); if err != nil {
		c.JSON(http.StatusOK, models.ErrorQuery(err, gin.H{})); return
	}
	c.JSON(http.StatusOK, pemesan); return
}
