package invoice

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type Usecase interface {
	GetInvoicePayment(ctx context.Context, idOrder int)(*models.InvoicePaymentPayload, error)
}
