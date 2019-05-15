package models

type WorkOrderPayload struct {
	NamaPemesan			string		`json:"nama_pemesan"`
	NoHp				string		`json:"no_hp"`
	JenisPesanan		string		`json:"jenis_pesanan"`
	IsiPesanan			string		`json:"isi_pesanan"`
	JumlahPesanan		int			`json:"jumlah_pesanan"`
	TanggalKirim		string		`json:"tanggal_kirim"`
	TanggalPersiapan	string		`json:"tanggal_persiapan"`
}
