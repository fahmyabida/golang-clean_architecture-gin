package invoice

import "github.com/fahmyabida/golang-clean_architecture-gin/models"

type Usecase interface {
	GetDataInvoice()(*models.Invoice, error)
}
