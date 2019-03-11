package mahasiswa

import (
	"context"
	"github.com/google/uuid"
	"github.com/me/golang-clean_architecture-gin_gonic/models"
)

// Usecase represent the article's usecases
type Usecase interface {
	//Fetch(ctx context.Context, cursor string, num int64) ([]*models.Mahasiswa, string, error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Mahasiswa, error)
	//Update(ctx context.Context, ar *models.Mahasiswa) error
	//GetByTitle(ctx context.Context, title string) (*models.Mahasiswa, error)
	//Store(context.Context, *models.Mahasiswa) error
	//Delete(ctx context.Context, id int64) error
}