package main

import (
	"bufio"
	"fmt"
	"strings"
)

// ============================================================
// MENCARI STOK
// ============================================================

func findStok(
	stok []DataStok,
	kode string,
) DataStok {

	for _, item := range stok {

		if strings.EqualFold(
			item.KodeBarang,
			kode,
		) {
			return item
		}
	}

	return DataStok{}
}

// ============================================================
// MENCARI HARGA
// ============================================================

func findHarga(
	harga []DataHarga,
	kode string,
) DataHarga {

	for _, item := range harga {

		if strings.EqualFold(
			item.KodeBarang,
			kode,
		) {
			return item
		}
	}

	return DataHarga{}
}

// ============================================================
// MENCARI LOKASI
// ============================================================

func findLokasi(
	lokasi []LokasiGudang,
	kode string,
) LokasiGudang {

	for _, item := range lokasi {

		if strings.EqualFold(
			item.KodeBarang,
			kode,
		) {
			return item
		}
	}

	return LokasiGudang{}
}

// ============================================================
// MENCARI SUPPLIER
// ============================================================

func findSupplier(
	suppliers []DataSupplier,
	kode string,
) DataSupplier {

	for _, item := range suppliers {

		if strings.EqualFold(
			item.KodeBarang,
			kode,
		) {
			return item
		}
	}

	return DataSupplier{}
}

// ============================================================
// CEK KODE BARANG
// ============================================================

func kodeSudahAda(
	spareparts []Sparepart,
	kode string,
) bool {

	for _, item := range spareparts {

		if strings.EqualFold(
			item.KodeBarang,
			kode,
		) {
			return true
		}
	}

	return false
}

// ============================================================
// MENCARI ID BERIKUTNYA
// ============================================================

func nextID(
	spareparts []Sparepart,
) int {

	max := 0

	for _, item := range spareparts {

		if item.No > max {
			max = item.No
		}
	}

	return max + 1
}

// ============================================================
// MENCARI BARANG BERDASARKAN KODE
// ============================================================

func cariBarang(
	reader *bufio.Reader,
	spareparts []Sparepart,
	stok []DataStok,
	harga []DataHarga,
	lokasi []LokasiGudang,
	suppliers []DataSupplier,
) {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("            CARI BARANG")
	fmt.Println("======================================")

	// Meminta kode barang dari user
	kode := readRequired(
		reader,
		"Masukkan Kode Barang: ",
	)

	// Mencari barang
	for _, item := range spareparts {

		// Membandingkan kode tanpa memperhatikan huruf besar/kecil
		if strings.EqualFold(
			item.KodeBarang,
			kode,
		) {

			// Mencari data tambahan berdasarkan kode barang
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

			// ==========================================
			// MENAMPILKAN HASIL PENCARIAN
			// ==========================================

			fmt.Println()
			fmt.Println("======================================")
			fmt.Println("          DATA BARANG")
			fmt.Println("======================================")

			fmt.Println(
				"Kode Barang :",
				item.KodeBarang,
			)

			fmt.Println(
				"Nama        :",
				item.Identitas.Nama,
			)

			fmt.Println(
				"Kategori    :",
				item.Identitas.Kategori,
			)

			fmt.Println(
				"Merek       :",
				item.Identitas.Merek,
			)

			fmt.Println(
				"Tipe Motor  :",
				strings.Join(
					item.Spesifikasi.TipeMotor,
					", ",
				),
			)

			fmt.Println(
				"Volume      :",
				item.Spesifikasi.Ukuran.Volume,
			)

			fmt.Println(
				"Viskositas  :",
				item.Spesifikasi.Ukuran.Viskositas,
			)

			fmt.Println(
				"Stok        :",
				dataStok.Stok.Jumlah,
			)

			fmt.Println(
				"Satuan      :",
				dataStok.Stok.Satuan,
			)

			fmt.Println(
				"Minimum     :",
				dataStok.Stok.Minimum,
			)

			fmt.Println(
				"Harga Beli  :",
				formatRupiah(dataHarga.Harga.Beli),
			)

			fmt.Println(
				"Harga Jual  :",
				formatRupiah(dataHarga.Harga.Jual),
			)

			fmt.Println(
				"Gudang      :",
				dataLokasi.LokasiGudang.Gudang,
			)

			fmt.Println(
				"Rak         :",
				dataLokasi.LokasiGudang.Rak.Kode,
			)

			fmt.Println(
				"Lantai      :",
				dataLokasi.LokasiGudang.Rak.Lantai,
			)

			fmt.Println(
				"Supplier    :",
				dataSupplier.Supplier.Nama,
			)

			fmt.Println(
				"Telepon     :",
				dataSupplier.Supplier.Kontak.Telepon,
			)

			fmt.Println(
				"Email       :",
				dataSupplier.Supplier.Kontak.Email,
			)

			fmt.Println("======================================")

			// Barang sudah ditemukan
			return
		}
	}

	// Kalau barang tidak ditemukan
	fmt.Println()
	fmt.Println(
		"Barang dengan kode",
		kode,
		"tidak ditemukan.",
	)
}
