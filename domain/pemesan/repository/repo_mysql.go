package repository

import (
	"context"
	"database/sql"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/sirupsen/logrus"
)

type pemesanRepo struct {
	Conn *sql.DB
}


func NewMysqlArticleRepository(Conn *sql.DB) pemesan.Repository{
	return &pemesanRepo{Conn}
}

func (m *pemesanRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Pemesan, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Pemesan, 0)
	for rows.Next() {
		row := new(models.Pemesan)
		err = rows.Scan(
			&row.Id,
			&row.Nama,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (r *pemesanRepo) GetObjectById(ctx context.Context, id int) (*models.Pemesan, error) {
	query := `SELECT
				id,
				nama,
			FROM pemesan WHERE id=?`
	list, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	nilaiMk := &models.Pemesan{}
	if len(list) > 0 {
		nilaiMk = list[0]
	} else {return nil, utils.ErrNotFound}
	return nilaiMk, nil
}