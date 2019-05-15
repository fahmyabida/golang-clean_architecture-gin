package pemesan

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type Usecase interface {
	GetPemesan(ctx context.Context, id int) (*models.Pemesan, error)
}