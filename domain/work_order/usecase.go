package work_order

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type Usecase interface {
	GetWorkOrder(ctx context.Context, idOrder int)(*models.WorkOrderPayload, error)
}
