package models

type Menu struct {
	Id 				int		`json:"id"`
	IdJenisPesanan	int		`json:"id_jensi_pesanan"`
	Isi				string	`json:"isi"`
	Harga			int 	`json:"harga"`
}
