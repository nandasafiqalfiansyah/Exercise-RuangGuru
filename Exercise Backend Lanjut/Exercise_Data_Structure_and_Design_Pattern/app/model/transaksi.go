package model

import "time"

type Transaksi struct {
	ID        int
	Deskripsi string
	Jumlah    float64
	Tanggal   time.Time
}


