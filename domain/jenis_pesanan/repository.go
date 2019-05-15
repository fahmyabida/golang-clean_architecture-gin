package jenis_pesanan

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
)

type Repository interface {
	GetObjectById(ctx context.Context, id int) (*models.JenisPesanan, error)
}
