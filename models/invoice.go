package models

type InvoicePaymentPayload struct {
	NamaPemesan		string		`json:"nama_pemesan"`
	TanggalPesan	string		`json:"tanggal_pesan"`
	JumlahPesanan	int			`json:"jumlah_pesanan"`
	HargaPerPorsi	int			`json:"harga_per_porsi"`
	GrandTotal		int 		`json:"grand_total"`
}