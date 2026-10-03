package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// ============================================================
// MENAMPILKAN INVENTORY
// ============================================================

func showInventory(
	spareparts []Sparepart,
	stok []DataStok,
	harga []DataHarga,
	lokasi []LokasiGudang,
	suppliers []DataSupplier,
) {

	fmt.Println("\n================ DATA INVENTORY SPAREPART ================")

	// Membuat tabel agar output lebih rapi
	w := tabwriter.NewWriter(
		os.Stdout,
		0,
		0,
		2,
		' ',
		0,
	)

	// Header tabel
	fmt.Fprintln(
		w,
		"No\tKode\tNama\tKategori\tMerek\tTipe Motor\tStok\tSatuan\tMin\tHarga Beli\tHarga Jual\tGudang\tRak\tSupplier",
	)

	// Garis tabel
	fmt.Fprintln(
		w,
		"--\t----\t----\t--------\t-----\t----------\t----\t------\t---\t----------\t----------\t------\t---\t--------",
	)

	// Perulangan data sparepart
	for _, item := range spareparts {

		// Mengambil data berdasarkan kode barang
		dataStok := findStok(
			stok,
			item.KodeBarang,
		)

		dataHarga := findHarga(
			harga,
			item.KodeBarang,
		)

		dataLokasi := findLokasi(
			lokasi,
			item.KodeBarang,
		)

		dataSupplier := findSupplier(
			suppliers,
			item.KodeBarang,
		)

		// Menampilkan data
		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\t%s\t%d\t%s\t%d\t%s\t%s\t%s\t%s\t%s\n",

			item.No,

			item.KodeBarang,

			item.Identitas.Nama,

			item.Identitas.Kategori,

			item.Identitas.Merek,

			strings.Join(
				item.Spesifikasi.TipeMotor,
				", ",
			),

			dataStok.Stok.Jumlah,

			dataStok.Stok.Satuan,

			dataStok.Stok.Minimum,

			formatRupiah(
				dataHarga.Harga.Beli,
			),

			formatRupiah(
				dataHarga.Harga.Jual,
			),

			dataLokasi.LokasiGudang.Gudang,

			dataLokasi.LokasiGudang.Rak.Kode,

			dataSupplier.Supplier.Nama,
		)
	}

	// Menampilkan tabel
	w.Flush()

	// Total data
	fmt.Println("\nTotal data:", len(spareparts))
}

// ============================================================
// FORMAT RUPIAH
// ============================================================

func formatRupiah(n int) string {

	// Mengubah angka menjadi string
	s := fmt.Sprintf("%d", n)

	var hasil string

	// Menghitung dari belakang
	for i, angka := range s {

		if i > 0 && (len(s)-i)%3 == 0 {
			hasil += "."
		}

		hasil += string(angka)
	}

	return "Rp " + hasil
}