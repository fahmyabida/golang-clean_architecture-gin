package http

import (
	"context"
	"errors"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type httpInvoiceHandler struct {
	iUsecase 		invoice.Usecase
}

func NewInvoiceHttpHandler(e *gin.Engine, iU invoice.Usecase) {
	handler := &httpInvoiceHandler{
		iUsecase: iU,
	}
	e.GET("/invoice-payment", handler.GetInvoiceOrderPayment)
}

func (h *httpInvoiceHandler) GetInvoiceOrderPayment(c *gin.Context){
	idParams := c.Request.URL.Query().Get("id")
	idOrder, err := strconv.Atoi(idParams); if err!=nil {
		c.JSON(http.StatusBadRequest, utils.Error(errors.New("Bad Request"), gin.H{})); return}
	ctx := c.Request.Context(); if ctx == nil {ctx = context.Background()}
	invoice, err := h.iUsecase.GetInvoicePayment(ctx, idOrder); if err != nil {
		c.JSON(http.StatusOK, utils.Error(err, gin.H{})); return
	}
	c.JSON(http.StatusOK, invoice); return
}
