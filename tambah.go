package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Fungsi untuk menambahkan barang baru
func tambahBarang(
	reader *bufio.Reader,
	spareparts []Sparepart,
	stok []DataStok,
	harga []DataHarga,
	lokasi []LokasiGudang,
	suppliers []DataSupplier,
) ([]Sparepart, []DataStok, []DataHarga, []LokasiGudang, []DataSupplier) {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("          TAMBAH BARANG")
	fmt.Println("======================================")

	// Membuat kode barang secara otomatis
	id := nextID(spareparts)

	kodeBarang := fmt.Sprintf("BRG-%03d", id)

	fmt.Println("Kode Barang :", kodeBarang)

	// ==============================
	// DATA IDENTITAS BARANG
	// ==============================

	nama := readRequired(reader, "Nama Barang : ")

	// Pilihan kategori
	kategoriPilihan := []string{
		"Oli",
		"Kelistrikan",
		"Rem",
		"Mesin",
	}

	kategori := pilihMenu(
		reader,
		"Pilih Kategori:",
		kategoriPilihan,
	)

	// Pilihan merek
	merekPilihan := []string{
		"AHM",
		"NGK",
		"Yamaha",
		"Suzuki",
	}

	merek := pilihMenu(
		reader,
		"Pilih Merek:",
		merekPilihan,
	)

	// ==============================
	// DATA SPESIFIKASI
	// ==============================

	tipeMotorPilihan := []string{
		"Beat",
		"Vario",
		"Scoopy",
		"PCX",
		"ADV",
	}

	tipeMotor := pilihBanyak(
		reader,
		"Pilih Tipe Motor:",
		tipeMotorPilihan,
	)

	volume := readRequired(reader, "Volume : ")
	viskositas := readRequired(reader, "Viskositas : ")

	// ==============================
	// MEMBUAT DATA SPAREPART
	// ==============================

	dataSparepart := Sparepart{
		No:         id,
		KodeBarang: kodeBarang,

		Identitas: Identitas{
			Nama:     nama,
			Kategori: kategori,
			Merek:    merek,
		},

		Spesifikasi: Spesifikasi{
			TipeMotor: tipeMotor,

			Ukuran: Ukuran{
				Volume:     volume,
				Viskositas: viskositas,
			},
		},
	}

	// ==============================
	// DATA STOK
	// ==============================

	jumlahStok := readInt(reader, "Jumlah Stok : ")

	satuanPilihan := []string{
		"Pcs",
		"Botol",
		"Set",
		"Box",
	}

	satuan := pilihMenu(
		reader,
		"Pilih Satuan:",
		satuanPilihan,
	)

	stokMinimum := readInt(reader, "Stok Minimum : ")

	dataStok := DataStok{
		KodeBarang: kodeBarang,

		Stok: StokInfo{
			Jumlah:  jumlahStok,
			Satuan:  satuan,
			Minimum: stokMinimum,
		},
	}

	// ==============================
	// DATA HARGA
	// ==============================

	hargaBeli := readInt(reader, "Harga Beli : ")
	hargaJual := readInt(reader, "Harga Jual : ")

	dataHarga := DataHarga{
		KodeBarang: kodeBarang,

		Harga: HargaInfo{
			Beli: hargaBeli,
			Jual: hargaJual,
		},
	}

	// ==============================
	// DATA LOKASI GUDANG
	// ==============================

	gudangPilihan := []string{
		"Gudang A",
		"Gudang B",
		"Gudang C",
	}

	gudang := pilihMenu(
		reader,
		"Pilih Gudang:",
		gudangPilihan,
	)

	rakPilihan := []string{
		"A-01",
		"A-02",
		"A-03",
		"B-01",
		"B-02",
		"B-03",
	}

	rakKode := pilihMenu(
		reader,
		"Pilih Rak:",
		rakPilihan,
	)

	lantai := readInt(reader, "Lantai Gudang : ")

	dataLokasi := LokasiGudang{
		KodeBarang: kodeBarang,

		LokasiGudang: LokasiGudangInfo{
			Gudang: gudang,

			Rak: Rak{
				Kode:   rakKode,
				Lantai: lantai,
			},
		},
	}

	// ==============================
	// DATA SUPPLIER
	// ==============================

	supplierPilihan := []string{
		"Astra Honda",
		"NGK Indonesia",
		"Yamaha Indonesia",
		"Suzuki Indonesia",
	}

	supplierNama := pilihMenu(
		reader,
		"Pilih Supplier:",
		supplierPilihan,
	)

	// Kontak supplier tetap bisa dimasukkan manual
	telepon := readRequired(reader, "Nomor Telepon Supplier : ")
	email := readRequired(reader, "Email Supplier : ")

	dataSupplier := DataSupplier{
		KodeBarang: kodeBarang,

		Supplier: SupplierInfo{
			Nama: supplierNama,

			Kontak: KontakInfo{
				Telepon: telepon,
				Email:   email,
			},
		},
	}

	// ==============================
	// MENAMBAHKAN DATA KE SLICE
	// ==============================

	spareparts = append(spareparts, dataSparepart)
	stok = append(stok, dataStok)
	harga = append(harga, dataHarga)
	lokasi = append(lokasi, dataLokasi)
	suppliers = append(suppliers, dataSupplier)

	// ==============================
	// MENYIMPAN KE FILE JSON
	// ==============================

	if err := saveJSON(sparepartFile, spareparts); err != nil {
		fmt.Println("Gagal menyimpan sparepart:", err)
		return spareparts, stok, harga, lokasi, suppliers
	}

	if err := saveJSON(stokFile, stok); err != nil {
		fmt.Println("Gagal menyimpan stok:", err)
		return spareparts, stok, harga, lokasi, suppliers
	}

	if err := saveJSON(hargaFile, harga); err != nil {
		fmt.Println("Gagal menyimpan harga:", err)
		return spareparts, stok, harga, lokasi, suppliers
	}

	if err := saveJSON(lokasiFile, lokasi); err != nil {
		fmt.Println("Gagal menyimpan lokasi:", err)
		return spareparts, stok, harga, lokasi, suppliers
	}

	if err := saveJSON(supplierFile, suppliers); err != nil {
		fmt.Println("Gagal menyimpan supplier:", err)
		return spareparts, stok, harga, lokasi, suppliers
	}

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("     BARANG BERHASIL DITAMBAHKAN")
	fmt.Println("======================================")
	fmt.Println("Kode Barang :", kodeBarang)
	fmt.Println("Nama        :", nama)
	fmt.Println("Kategori    :", kategori)
	fmt.Println("Merek       :", merek)
	fmt.Println("Tipe Motor  :", strings.Join(tipeMotor, ", "))
	fmt.Println("Stok        :", jumlahStok, satuan)
	fmt.Println("Gudang      :", gudang)
	fmt.Println("Rak         :", rakKode)
	fmt.Println("Supplier    :", supplierNama)
	fmt.Println("======================================")

	return spareparts, stok, harga, lokasi, suppliers
}
