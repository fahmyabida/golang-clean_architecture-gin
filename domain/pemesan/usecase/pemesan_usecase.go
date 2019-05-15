package usecase

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"time"
)

type pemesanUsecase struct {
pemesanRepo    pemesan.Repository
contextTimeout time.Duration
}


func NewPemesanUsecase(pR pemesan.Repository, timeout time.Duration) pemesan.Usecase{
	return &pemesanUsecase{
		pemesanRepo		: pR,
		contextTimeout	: timeout,
	}
}

func (pU *pemesanUsecase) GetPemesan(ctx context.Context, id int) (*models.Pemesan, error) {
	ctx, cancel := context.WithTimeout(ctx, pU.contextTimeout)
	defer cancel()
	res, err := pU.pemesanRepo.GetObjectById(ctx, id); if err != nil {return nil, err}
	return res, nil
}