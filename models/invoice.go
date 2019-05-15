package models

type InvoiceMenuPayload struct {
	NamaPemesan		string		`json:"nama_pemesan"`
	NoHp			string		`json:"no_hp"`
	JenisPesanan	string		`json:"jenis_pesanan"`
	IsiPesanan		string		`json:"isi_pesanan"`
	JumlahPesanan	int			`json:"jumlah_pesanan"`
	TanggalKirim	string		`json:"tanggal_kirim"`
}

type InvoicePaymentPayload struct {
	NamaPemesan		string		`json:"nama_pemesan"`
	//JenisPesanan	string		`json:"jenis_pesanan"`
	//IsiPesanan		string		`json:"isi_pesanan"`
	TanggalPesan	string		`json:"tanggal_pesan"`
	JumlahPesanan	int			`json:"jumlah_pesanan"`
	//TanggalKirim	string		`json:"tanggal_kirim"`
	HargaPerPorsi	int			`json:"harga_per_porsi"`
	GrandTotal		int 		`json:"grand_total"`
}