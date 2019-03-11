package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/me/golang-clean_architecture-gin_gonic/mahasiswa"
	"github.com/me/golang-clean_architecture-gin_gonic/models"
	"time"
)

type mahasiswaUsecase struct {
mahasiswaRepo    mahasiswa.Repository
contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of article.Usecase interface
func NewArticleUsecase(mhsRepo mahasiswa.Repository, timeout time.Duration) mahasiswa.Usecase {
	return &mahasiswaUsecase{
		mahasiswaRepo:    mhsRepo,
		contextTimeout: timeout,
	}
}

func (mhsU *mahasiswaUsecase) GetByUUID(c context.Context, uuid uuid.UUID) (*models.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(c, mhsU.contextTimeout)
	defer cancel()
	res, err := mhsU.mahasiswaRepo.GetByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return res, nil
}