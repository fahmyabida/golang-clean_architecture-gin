package models

import "gopkg.in/guregu/null.v3/zero"

type Mahasiswa struct {
	Uuid	string 			`json:"uuid"`
	Nim 	string 			`json:"nim"`
	Nama 	string			`json:"nama_mahasiswa"`
	Ipk 	zero.Float 		`json:"ipk"`
}
