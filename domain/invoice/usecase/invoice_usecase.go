package usecase

import (
	"context"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/invoice"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/jenis_pesanan"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/menu"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/order"
	"github.com/fahmyabida/golang-clean_architecture-gin/domain/pemesan"
	"github.com/fahmyabida/golang-clean_architecture-gin/models"
	"time"
)

type invoiceUsecase struct {
	jPesanRepo		jenis_pesanan.Repository
	menuRepo 		menu.Repository
	orderRepo 		order.Repository
	pemesanRepo    	pemesan.Repository
	contextTimeout 	time.Duration
}

func NewInvoiceMenuUsecase(jpR jenis_pesanan.Repository, mR menu.Repository, oR order.Repository, pR pemesan.Repository, timeout time.Duration) invoice.Usecase  {
	return &invoiceUsecase{
		jpR,mR,oR,pR,timeout,
	}
}

func (u *invoiceUsecase) toPayloadInvoiceMenu(pemesan *models.Pemesan, jPesanan *models.JenisPesanan, menu *models.Menu, order *models.Order) (*models.InvoiceMenuPayload, error) {
	tmp := &models.InvoiceMenuPayload{
		NamaPemesan		: pemesan.Nama,
		NoHp			: pemesan.NoHp,
		JenisPesanan	: jPesanan.Nama,
		IsiPesanan		: menu.Isi,
		JumlahPesanan	: order.Total,
		TanggalKirim	: order.TanggalDikirim.String(),
	}
	return tmp,nil
}

func (u *invoiceUsecase) toPayloadInvoicePayment(pemesan *models.Pemesan, jPesanan *models.JenisPesanan, menu *models.Menu, order *models.Order) (*models.InvoicePaymentPayload, error) {
	tmp := &models.InvoicePaymentPayload{
		NamaPemesan		: pemesan.Nama,
		TanggalPesan	: order.TanggalPesan.String(),
		JumlahPesanan	: order.Total,
		HargaPerPorsi	: menu.Harga,
		GrandTotal		: (menu.Harga * order.Total),
	}
	return tmp, nil
}

func (u *invoiceUsecase) GetDataInvoiceMenu(ctx context.Context, idOrder int) (*models.InvoiceMenuPayload, error) {
	order, err 		:= u.orderRepo.GetObjectById(ctx, idOrder); if err!=nil{return nil, err}
	pemesan, err 	:= u.pemesanRepo.GetObjectById(ctx, order.IdPemesan); if err!=nil{return nil, err}
	menu, err 		:= u.menuRepo.GetObjectById(ctx, order.IdMenu); if err!=nil{return nil, err}
	jPesan, err 	:= u.jPesanRepo.GetObjectById(ctx, menu.IdJenisPesanan); if err!=nil{return nil, err}
	res, err		:= u.toPayloadInvoiceMenu(pemesan,jPesan,menu,order); if err!=nil{return nil, err}
	return res, nil
}


func (u *invoiceUsecase) GetDataInvoicePayment(ctx context.Context, idOrder int) (*models.InvoicePaymentPayload, error) {
	order, err 		:= u.orderRepo.GetObjectById(ctx, idOrder); if err!=nil{return nil, err}
	pemesan, err 	:= u.pemesanRepo.GetObjectById(ctx, order.IdPemesan); if err!=nil{return nil, err}
	menu, err 		:= u.menuRepo.GetObjectById(ctx, order.IdMenu); if err!=nil{return nil, err}
	jPesan, err 	:= u.jPesanRepo.GetObjectById(ctx, menu.IdJenisPesanan); if err!=nil{return nil, err}
	res, err		:= u.toPayloadInvoicePayment(pemesan,jPesan,menu,order); if err!=nil{return nil, err}
	return res, nil
}

//+ (menu.Harga * order.Total * 10/100),