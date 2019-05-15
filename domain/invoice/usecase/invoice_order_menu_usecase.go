package usecase

import (
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type invoiceOrderMenuUsecase struct {

}

func NewInvoiceOrderMenuUsecase() invoice.Usecase  {
	return &invoiceOrderMenuUsecase{}
}

func (u *invoiceOrderMenuUsecase) GetDataInvoice() (*models.InvoiceOrder, error) {
	panic("implement me")
}
