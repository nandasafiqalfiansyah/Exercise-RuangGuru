package service

import (
	"Exercise_Data_Structure_and_Design_Pattern/app/model"
)

type TransaksiRepository interface {
	Tambah(transaksi *model.Transaksi)
	Hapus(id int)
	TampilkanSemua() []model.Transaksi
	HitungTotal() float64
}
type TransaksiRepositoryImpl struct {
	transaksiList []model.Transaksi
}

func (r * TransaksiRepositoryImpl) Tambah(transaksi *model.Transaksi) {
	r.transaksiList = append(r.transaksiList, *transaksi)
}


func (r *TransaksiRepositoryImpl) Hapus(id int) {
	for i, t := range r.transaksiList {
		if t.ID == id {
			r.transaksiList = append(r.transaksiList[:i], r.transaksiList[i+1:]...)
			break
		}
	}
}

func (r *TransaksiRepositoryImpl) TampilkanSemua() []model.Transaksi {
	return r.transaksiList
}


func (r *TransaksiRepositoryImpl) HitungTotal() float64 {
	total := 0.0
	for _, t := range r.transaksiList {
		total += t.Jumlah
	}
	return total
}

var RepositoryInstance TransaksiRepository = &TransaksiRepositoryImpl{}

type TransaksiService struct {
	TransaksiRepo TransaksiRepository
}

func (s *TransaksiService) TambahTransaksi(transaksi *model.Transaksi) {
	s.TransaksiRepo.Tambah(transaksi)
}

func (s *TransaksiService) HapusTransaksi(id int) {
	s.TransaksiRepo.Hapus(id)
}

func (s *TransaksiService) TampilkanSemuaTransaksi() []model.Transaksi {
	return s.TransaksiRepo.TampilkanSemua()
}


func (s *TransaksiService) HitungTotalTransaksi() float64 {
	return s.TransaksiRepo.HitungTotal()
}