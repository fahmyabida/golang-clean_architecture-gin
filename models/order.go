package models

type order struct {
	Id				int		`json:"id"`
	Id_pemesan		int	    `json:"id_pemesan"`
	Jenis_pesanan	string	`json:"jenis_pesanan"`
	Jumlah			int		`json:"jumlah"`

}
