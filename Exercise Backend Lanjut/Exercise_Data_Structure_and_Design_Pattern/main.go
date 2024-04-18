package main

import (
	"Exercise_Data_Structure_and_Design_Pattern/app/model"
	"Exercise_Data_Structure_and_Design_Pattern/app/service"
	"fmt"
	"time"
)


func main() {
	// Inisialisasi layanan Transaksi.
	transaksiService := service.TransaksiService{TransaksiRepo: service.RepositoryInstance}

	// Tambahkan beberapa transaksi.
	transaksi1 := &model.Transaksi{ID: 1, Deskripsi: "Transaksi 1", Jumlah: 100, Tanggal: time.Now()}
	transaksi2 := &model.Transaksi{ID: 2, Deskripsi: "Transaksi 2", Jumlah: 200, Tanggal: time.Now()}
	transaksiService.TambahTransaksi(transaksi1)
	transaksiService.TambahTransaksi(transaksi2)

	// Tampilkan semua transaksi.
	fmt.Println("---------------------------------")
	fmt.Println("Semua Transaksi:")

	for _, t := range transaksiService.TampilkanSemuaTransaksi() {
        fmt.Printf("ID: %d\n", t.ID)
        fmt.Printf("Deskripsi: %s\n", t.Deskripsi)
        fmt.Printf("Jumlah: %.2f\n", t.Jumlah)
        fmt.Printf("Tanggal: %s\n", t.Tanggal.Format("02 Jan 2006 15:04:05"))
        fmt.Println("")
    }

	// Hitung total jumlah transaksi.
	total := transaksiService.HitungTotalTransaksi()
	fmt.Printf("Total Jumlah Transaksi: %.2f\n", total)

	// Hapus transaksi dengan ID 1.
	fmt.Println("---------------------------------")
	transaksiService.HapusTransaksi(1)

	// Tampilkan semua transaksi setelah dihapus.
	fmt.Println("Semua Transaksi Setelah dihapus:")
	for _, t := range transaksiService.TampilkanSemuaTransaksi() {
        fmt.Printf("ID: %d\n", t.ID)
        fmt.Printf("Deskripsi: %s\n", t.Deskripsi)
        fmt.Printf("Jumlah: %.2f\n", t.Jumlah)
        fmt.Printf("Tanggal: %s\n", t.Tanggal.Format("02 Jan 2006 15:04:05"))
        fmt.Println("")
    }
}