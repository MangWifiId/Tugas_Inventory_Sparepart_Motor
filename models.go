package main

// ============================================================
// STRUCT SPAREPART
// ============================================================

// Struct utama sparepart
type Sparepart struct {
	No          int         `json:"no"`
	KodeBarang  string      `json:"kode_barang"`
	Identitas   Identitas   `json:"identitas"`
	Spesifikasi Spesifikasi `json:"spesifikasi"`
}

// Struct Identitas berada di dalam Sparepart
type Identitas struct {
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
	Merek    string `json:"merek"`
}

// Struct Spesifikasi berada di dalam Sparepart
type Spesifikasi struct {
	TipeMotor []string `json:"tipe_motor"`
	Ukuran    Ukuran   `json:"ukuran"`
}

// Struct Ukuran berada di dalam Spesifikasi
type Ukuran struct {
	Volume     string `json:"volume"`
	Viskositas string `json:"viskositas"`
}

// ============================================================
// STRUCT STOK
// ============================================================

type DataStok struct {
	KodeBarang string   `json:"kode_barang"`
	Stok       StokInfo `json:"stok"`
}

// Informasi stok
type StokInfo struct {
	Jumlah  int    `json:"jumlah"`
	Satuan  string `json:"satuan"`
	Minimum int    `json:"minimum"`
}

// ============================================================
// STRUCT HARGA
// ============================================================

type DataHarga struct {
	KodeBarang string    `json:"kode_barang"`
	Harga      HargaInfo `json:"harga"`
}

// Informasi harga
type HargaInfo struct {
	Beli int `json:"beli"`
	Jual int `json:"jual"`
}

// ============================================================
// STRUCT LOKASI GUDANG
// ============================================================

type LokasiGudang struct {
	KodeBarang   string           `json:"kode_barang"`
	LokasiGudang LokasiGudangInfo `json:"lokasi_gudang"`
}

// Informasi gudang
type LokasiGudangInfo struct {
	Gudang string `json:"gudang"`
	Rak    Rak    `json:"rak"`
}

// Informasi rak
type Rak struct {
	Kode   string `json:"kode"`
	Lantai int    `json:"lantai"`
}

// ============================================================
// STRUCT SUPPLIER
// ============================================================

type DataSupplier struct {
	KodeBarang string       `json:"kode_barang"`
	Supplier   SupplierInfo `json:"supplier"`
}

// Informasi supplier
type SupplierInfo struct {
	Nama   string     `json:"nama"`
	Kontak KontakInfo `json:"kontak"`
}

// Informasi kontak supplier
type KontakInfo struct {
	Telepon string `json:"telepon"`
	Email   string `json:"email"`
}

// Data yang diterima API saat menambah barang
type TambahBarangRequest struct {
	Identitas   Identitas        `json:"identitas"`
	Spesifikasi Spesifikasi      `json:"spesifikasi"`
	Stok        StokInfo         `json:"stok"`
	Harga       HargaInfo        `json:"harga"`
	Lokasi      LokasiGudangInfo `json:"lokasi_gudang"`
	Supplier    SupplierInfo     `json:"supplier"`
}
