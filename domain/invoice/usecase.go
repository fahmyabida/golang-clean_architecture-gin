package invoice

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type Usecase interface {
	GetDataInvoiceMenu(ctx context.Context, idOrder int)(*models.InvoiceMenuPayload, error)
	GetDataInvoicePayment(ctx context.Context, idOrder int)(*models.InvoicePaymentPayload, error)
}
