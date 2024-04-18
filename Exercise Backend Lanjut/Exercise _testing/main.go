package main

import (
	"Exercise_Testing/helper"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Transaksi struct {
	ID        int
	Deskripsi string
	Jumlah    int
	Tanggal   time.Time
}

var transaksiList []Transaksi

func init() {
	transaksiList = []Transaksi{
        {
            ID:        1,
            Deskripsi: "Pembelian barang A",
            Jumlah:    100,
            Tanggal:   time.Now(),
        },
        {
            ID:       2,
            Deskripsi: "Pembelian barang B",
            Jumlah:    150,
            Tanggal:   time.Now().AddDate(0, 0, -1),
        },
    }
}

func AmbilTransaksi() []Transaksi {
	return transaksiList
}

func TambahTransaksi(deskripsi string, jumlah int) (string, error) {
	var (
		result string
		err    error
	)

	lastTransaksi := transaksiList[len(transaksiList)-1]
	lastID := lastTransaksi.ID
	lastID++
	
	if deskripsi == "" || jumlah == 0 {
		err = errors.New("Deskripsi dan jumlah tidak boleh kosong")
	} else {
		transaksi := Transaksi{
			ID:        lastID,
			Deskripsi: deskripsi,
			Jumlah:    jumlah,
			Tanggal:   time.Now(),
		}
		transaksiList = append(transaksiList, transaksi)
		result = "Transaksi ditambahkan dengan id "+strconv.Itoa(lastID)
	}
	if err == nil {
		err = nil
	}

	return result, err
}


func HapusTransaksi(id int) (string, error) {
	var (
		result string
		err    error
	)
	
	if id == 0  {
		err = errors.New("ID tidak boleh kosong")
	} else{
		for i, transaksi := range transaksiList {
			if transaksi.ID == id {
				transaksiList = append(transaksiList[:i], transaksiList[i+1:]...)
				result = "Transaksi dihapus dengan id " + strconv.Itoa(id)
				break
			}
		}
	}

	if err == nil {
		err = nil
	}
	return result, err
}


func main() {
	for {
		helper.ClearScreen()
		transaksiList := AmbilTransaksi()
		for _, transaksi := range transaksiList {
			fmt.Printf("ID: %d\n", transaksi.ID)
			fmt.Printf("Deskripsi: %s\n", transaksi.Deskripsi)
			fmt.Printf("Jumlah: %d\n", transaksi.Jumlah)
			fmt.Printf("Tanggal: %s\n", transaksi.Tanggal.Format("2006-01-02"))
			fmt.Println()
		}
		fmt.Println("Selamat datang di FinanceTrack!")
		fmt.Println("1. Tambah Transaksi")
		fmt.Println("2. hapus Transaksi")
		fmt.Println("3. Exit")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== TambahTransaksi ===")
			fmt.Print("Deskripsi: ")
			des, _ := reader.ReadString('\n')
			des = strings.TrimSpace(des)
			fmt.Printf("Jumlah: ")
			jum, _ := reader.ReadString('\n')
			jum = strings.TrimSpace(jum)
			jumInt, _ := strconv.Atoi(jum)
			msg, err := TambahTransaksi(des, jumInt)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			fmt.Println("=== hapusTransaksi ===")
			fmt.Print("Masukkan ID transaksi yang ingin dihapus: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			idInt, _ := strconv.Atoi(id)
			msg, err := HapusTransaksi(idInt)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			fmt.Println("Terimakasih!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

	}
}