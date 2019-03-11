package repository

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/me/golang-clean_architecture-gin_gonic/mahasiswa"
	"github.com/me/golang-clean_architecture-gin_gonic/models"
	"github.com/sirupsen/logrus"
)

type mysqlMahasiswaRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) mahasiswa.Repository{
	return &mysqlMahasiswaRepository{Conn}
}

func (m *mysqlMahasiswaRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Mahasiswa, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Mahasiswa, 0)
	for rows.Next() {
		row := new(models.Mahasiswa)
		err = rows.Scan(
			&row.Uuid,
			&row.Nama,
			&row.Nim,
			&row.Ipk,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (m *mysqlMahasiswaRepository) GetByUUID(ctx context.Context, uuid uuid.UUID) (*models.Mahasiswa, error){
	query := `SELECT uuid, nim, nama_mahasiswa, ipk FROM mahasiswa WHERE uuid=?`
	list, err := m.fetch(ctx, query, uuid)
	if err != nil {
		return nil, err
	}
	mahasiswa := &models.Mahasiswa{}
	if len(list) > 0 {
		mahasiswa = list[0]
	} else {
		return nil, models.ErrNotFound
	}
	return mahasiswa, nil
}