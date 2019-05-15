package usecase

import (
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type invoiceOrderPaymentUsecase struct {

}

func NewInvoiceOrderPaymenUsecase() invoice.Usecase  {
	return &invoiceOrderPaymentUsecase{}
}

func (u *invoiceOrderPaymentUsecase) GetDataInvoice() (*models.InvoicePayment, error) {
	panic("implement me")
}
