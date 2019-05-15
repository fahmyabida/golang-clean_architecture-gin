package usecase

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/jenis_pesanan"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/menu"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/order"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/work_order"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"time"
)

type workOrderUsecase struct {
	jPesanRepo		jenis_pesanan.Repository
	menuRepo 		menu.Repository
	orderRepo 		order.Repository
	pemesanRepo    	pemesan.Repository
	contextTimeout 	time.Duration
}

func NewWorkOrderUsecase(jpR jenis_pesanan.Repository, mR menu.Repository, oR order.Repository, pR pemesan.Repository, timeout time.Duration) work_order.Usecase{
	return &workOrderUsecase{
		jpR,mR,oR,pR,timeout,
	}
}

func (u *workOrderUsecase) toPayloadInvoiceMenu(pemesan *models.Pemesan, jPesanan *models.JenisPesanan, menu *models.Menu, order *models.Order) (*models.WorkOrderPayload, error) {
	tmp := &models.WorkOrderPayload{
		NamaPemesan			: pemesan.Nama,
		NoHp				: pemesan.NoHp,
		JenisPesanan		: jPesanan.Nama,
		IsiPesanan			: menu.Isi,
		JumlahPesanan		: order.Total,
		TanggalKirim		: order.TanggalDikirim.String(),
		TanggalPersiapan	: order.TanggalDikirim.AddDate(0, 0, -2).String(),
	}
	return tmp,nil
}

func (u *workOrderUsecase) GetWorkOrder(ctx context.Context, idOrder int) (*models.WorkOrderPayload, error) {
	order, err 		:= u.orderRepo.GetObjectById(ctx, idOrder); if err!=nil{return nil, err}
	pemesan, err 	:= u.pemesanRepo.GetObjectById(ctx, order.IdPemesan); if err!=nil{return nil, err}
	menu, err 		:= u.menuRepo.GetObjectById(ctx, order.IdMenu); if err!=nil{return nil, err}
	jPesan, err 	:= u.jPesanRepo.GetObjectById(ctx, menu.IdJenisPesanan); if err!=nil{return nil, err}
	res, err		:= u.toPayloadInvoiceMenu(pemesan,jPesan,menu,order); if err!=nil{return nil, err}
	return res, nil
}

