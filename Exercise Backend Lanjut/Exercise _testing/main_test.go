package main_test

import (
	main "Exercise_Testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)


var _ = Describe("Main", func() {

	// Testing AmbilTransaksi
	Describe("Test Main", func() {
		It("Test AmbilTransaksi", func() {
			transaksiList := main.AmbilTransaksi()
			Expect(len(transaksiList)).To(Equal(2))
		})
	})

	// Testing TambahTransaksi
	Describe("Test TambahTransaksi", func() {
		When("Deskripsi dan jumlah tidak boleh kosong", func() {
			It("should return error", func() {
				result, err := main.TambahTransaksi("", 0)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Deskripsi dan jumlah tidak boleh kosong"))
				Expect(result).To(Equal(""))
			})
		})

		When("TambahTransaksi", func() {
			It("should add transaction with id 3", func() {
				result, err := main.TambahTransaksi("Pembelian barang A", 3)
				Expect(err).To(BeNil())
				Expect(result).To(Equal("Transaksi ditambahkan dengan id 3"))
			})
			It("should add transaction with id 4", func() {
				result, err := main.TambahTransaksi("Pembelian barang B", 4)
				Expect(err).To(BeNil())
				Expect(result).To(Equal("Transaksi ditambahkan dengan id 4"))
			})
		})
	})

	// testing HapusTransaksi
	Describe("Test HapusTransaksi", func() {
		It("should delete transaction with id 0", func() {
			result, err := main.HapusTransaksi(0)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("ID tidak boleh kosong"))
			Expect(result).To(Equal(""))
		})
		It("should delete transaction with id 2", func() {
			result, err := main.HapusTransaksi(2)
			Expect(err).To(BeNil())
			Expect(result).To(Equal("Transaksi dihapus dengan id 2"))
		})
		It("should delete transaction with id 1", func() {
			result, err := main.HapusTransaksi(1)
			Expect(err).To(BeNil())
			Expect(result).To(Equal("Transaksi dihapus dengan id 1"))
		})
	})
})