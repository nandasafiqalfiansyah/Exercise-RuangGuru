# Exercise Testing

# FinanceTrack - Package main

FinanceTrack adalah sebuah program sederhana untuk melacak transaksi keuangan. Package `main` dalam program ini menyediakan fitur utama untuk menambah dan menghapus transaksi, serta menampilkan daftar transaksi yang telah ada.

## Fitur Utama

### 1. AmbilTransaksi

- Fungsi `AmbilTransaksi` digunakan untuk mengambil daftar transaksi yang sudah ada.

### 2. TambahTransaksi

- Fungsi `TambahTransaksi(deskripsi string, jumlah int) (string, error)` digunakan untuk menambahkan transaksi baru.
- Input:
  - `deskripsi`: Deskripsi dari transaksi yang ingin ditambahkan.
  - `jumlah`: Jumlah uang yang terlibat dalam transaksi.
- Output:
  - `string`: Pesan sukses jika transaksi berhasil ditambahkan.
  - `error`: Error jika terjadi kesalahan, misalnya deskripsi atau jumlah kosong.

### 3. HapusTransaksi

- Fungsi `HapusTransaksi(id int) (string, error)` digunakan untuk menghapus transaksi berdasarkan ID.
- Input:
  - `id`: ID dari transaksi yang ingin dihapus.
- Output:
  - `string`: Pesan sukses jika transaksi berhasil dihapus.
  - `error`: Error jika terjadi kesalahan, misalnya ID tidak valid atau kosong.

### 4. Main

- Fungsi `main` merupakan fungsi utama yang menjalankan program.
- Program akan menampilkan daftar transaksi yang sudah ada dan meminta input dari pengguna untuk menambah, menghapus, atau keluar dari program.

## Cara Menggunakan

1. Clone repository ini ke dalam komputer Anda.
2. Pastikan Anda memiliki Go compiler yang terinstal.
3. Buka terminal atau command prompt.
4. Masuk ke direktori tempat Anda menyimpan repository ini.
5. Ketik `go run main.go` dan tekan Enter untuk menjalankan program.

## Catatan

Pastikan `go mod tidy` untuk mengimpor paket-paket yang diperlukan sebelum menggunakan program ini .
