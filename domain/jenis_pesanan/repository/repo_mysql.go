package repository

import (
	"context"
	"database/sql"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/jenis_pesanan"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/sirupsen/logrus"
)

type jenisPesananRepo struct {
	Conn *sql.DB
}


func NewJenisPesananRepository(Conn *sql.DB) jenis_pesanan.Repository{
	return &jenisPesananRepo{Conn}
}

func (m *jenisPesananRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.JenisPesanan, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.JenisPesanan, 0)
	for rows.Next() {
		row := new(models.JenisPesanan)
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

func (r *jenisPesananRepo) GetObjectById(ctx context.Context, id int) (*models.JenisPesanan, error) {
	query := `SELECT
				id,
				nama
			FROM jenis_pesanan WHERE id=?`
	list, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	obj := &models.JenisPesanan{}
	if len(list) > 0 {
		obj = list[0]
	} else {return nil, utils.ErrNotFound}
	return obj, nil
}
