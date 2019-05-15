package repository

import (
	"context"
	"database/sql"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/menu"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/sirupsen/logrus"
)

type menuRepo struct {
	Conn *sql.DB
}


func NewMenuRepository(Conn *sql.DB) menu.Repository{
	return &menuRepo{Conn}
}

func (m *menuRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Menu, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Menu, 0)
	for rows.Next() {
		row := new(models.Menu)
		err = rows.Scan(
			&row.Id,
			&row.IdJenisPesanan,
			&row.Isi,
			&row.Harga,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (r *menuRepo) GetObjectById(ctx context.Context, id int) (*models.Menu, error) {
	query := `SELECT
				id,
				id_jenis_pesanan,
				isi,
				harga
			FROM menu WHERE id=?`
	list, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	obj := &models.Menu{}
	if len(list) > 0 {
		obj = list[0]
	} else {return nil, utils.ErrNotFound}
	return obj, nil
}
