package main

import (
	"encoding/json"
	"os"
)

// Nama file JSON
const sparepartFile = "data/sparepart.json"
const stokFile = "data/stok.json"
const hargaFile = "data/harga.json"
const lokasiFile = "data/lokasi.json"
const supplierFile = "data/supplier.json"

// ============================================================
// MEMBACA SEMUA DATA JSON
// ============================================================

func readAllData() (
	[]Sparepart,
	[]DataStok,
	[]DataHarga,
	[]LokasiGudang,
	[]DataSupplier,
	error,
) {

	// Membaca sparepart.json
	spareparts, err := readJSONSparepart()

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Membaca stok.json
	stok, err := readJSONStok()

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Membaca harga.json
	harga, err := readJSONHarga()

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Membaca lokasi.json
	lokasi, err := readJSONLokasi()

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	// Membaca supplier.json
	suppliers, err := readJSONSupplier()

	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	return spareparts, stok, harga, lokasi, suppliers, nil
}

// ============================================================
// BACA SPAREPART
// ============================================================

func readJSONSparepart() ([]Sparepart, error) {

	var data []Sparepart

	file, err := os.Open(sparepartFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)

	return data, err
}

// ============================================================
// BACA STOK
// ============================================================

func readJSONStok() ([]DataStok, error) {

	var data []DataStok

	file, err := os.Open(stokFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)

	return data, err
}

// ============================================================
// BACA HARGA
// ============================================================

func readJSONHarga() ([]DataHarga, error) {

	var data []DataHarga

	file, err := os.Open(hargaFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)

	return data, err
}

// ============================================================
// BACA LOKASI
// ============================================================

func readJSONLokasi() ([]LokasiGudang, error) {

	var data []LokasiGudang

	file, err := os.Open(lokasiFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)

	return data, err
}

// ============================================================
// BACA SUPPLIER
// ============================================================

func readJSONSupplier() ([]DataSupplier, error) {

	var data []DataSupplier

	file, err := os.Open(supplierFile)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)

	return data, err
}

// ============================================================
// MENYIMPAN DATA KE JSON
// ============================================================

func saveJSON(filename string, data interface{}) error {

	// Mengubah data Go menjadi JSON
	jsonData, err := json.MarshalIndent(
		data,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	// Membuat atau membuka file
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	// Menulis JSON ke file
	_, err = file.Write(jsonData)

	return err
}
