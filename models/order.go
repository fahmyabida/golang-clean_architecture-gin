package models

import "time"

type Order struct {
	Id				int			`json:"id"`
	IdMenu			int			`json:"id_menu"`
	IdPemesan		int	    	`json:"id_pemesan"`
	Total			int			`json:"total"`
	TanggalPesan	time.Time	`json:"tanggal_pesan"`
	TanggalDikirim	time.Time	`json:"tanggal_dikirim"`
}
