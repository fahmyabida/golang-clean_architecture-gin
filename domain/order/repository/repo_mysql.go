package repository

import (
	"context"
	"database/sql"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/order"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"github.com/fahmyabida/golang-clean_architecture-gin/utils"
	"github.com/sirupsen/logrus"
)

type orderRepo struct {
	Conn *sql.DB
}


func NewOrderRepository(Conn *sql.DB) order.Repository{
	return &orderRepo{Conn}
}

func (m *orderRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Order, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Order, 0)
	for rows.Next() {
		row := new(models.Order)
		err = rows.Scan(
			&row.Id,
			&row.IdMenu,
			&row.IdPemesan,
			&row.Total,
			&row.TanggalPesan,
			&row.TanggalDikirim,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (r *orderRepo) GetObjectById(ctx context.Context, id int) (*models.Order, error) {
	query := `SELECT
				id,
				id_menu,
				id_pemesan,
				total,
				tanggal_pesan,
				tanggal_dikirim
			FROM order_pesanan WHERE id=?`
	list, err := r.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}
	obj := &models.Order{}
	if len(list) > 0 {
		obj = list[0]
	} else {return nil, utils.ErrNotFound}
	return obj, nil
}