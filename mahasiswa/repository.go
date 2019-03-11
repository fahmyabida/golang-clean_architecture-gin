package mahasiswa

import (
	"context"
	"github.com/google/uuid"
	"github.com/me/golang-clean_architecture-gin_gonic/models"
)

// Repository represent the Mahasiswa's repository contract
type Repository interface {
	//Fetch(ctx gin.Context, cursor string, num int64) (res []*models.Mahasiswa, nextCursor string, err error)
	GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Mahasiswa, error)
	//GetByTitle(ctx gin.Context, title string) (*models.Mahasiswa, error)
	//Update(ctx gin.Context, ar *models.Mahasiswa) error
	//Store(ctx gin.Context, a *models.Mahasiswa) error
	//Delete(ctx gin.Context, id int64) error
}