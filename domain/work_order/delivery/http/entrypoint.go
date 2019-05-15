package http

import (
	"context"
	"errors"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/work_order"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type httpWorkOrderHandler struct {
	woUsecase 		work_order.Usecase
}

func NewWorkOrderHttpHandler(e *gin.Engine, woU work_order.Usecase) {
	handler := &httpWorkOrderHandler{
		woUsecase	: woU,
	}
	e.GET("/work-order", handler.GetWorkOrder)
}


func (h *httpWorkOrderHandler) GetWorkOrder(c *gin.Context){
	idParams := c.Request.URL.Query().Get("id")
	idOrder, err := strconv.Atoi(idParams); if err!=nil {
		c.JSON(http.StatusBadRequest, utils.Error(errors.New("Bad Request"), gin.H{})); return}
	ctx := c.Request.Context(); if ctx == nil {ctx = context.Background()}
	pemesan, err := h.woUsecase.GetWorkOrder(ctx, idOrder); if err != nil {
		c.JSON(http.StatusOK, utils.Error(err, gin.H{})); return
	}
	c.JSON(http.StatusOK, pemesan); return
}