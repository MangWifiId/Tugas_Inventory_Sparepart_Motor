package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Membaca input berupa teks dari user
func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)

	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	return strings.TrimSpace(input)
}

// Membaca input teks yang wajib diisi
func readRequired(reader *bufio.Reader, prompt string) string {
	for {
		input := readLine(reader, prompt)

		if input != "" {
			return input
		}

		fmt.Println("Input tidak boleh kosong.")
	}
}

// Membaca input angka
func readInt(reader *bufio.Reader, prompt string) int {
	for {
		input := readLine(reader, prompt)

		angka, err := strconv.Atoi(input)

		if err == nil {
			return angka
		}

		fmt.Println("Masukkan angka yang benar.")
	}
}

// Membuat menu pilihan
func pilihMenu(reader *bufio.Reader, judul string, pilihan []string) string {
	fmt.Println()
	fmt.Println(judul)

	// Menampilkan semua pilihan
	for i, item := range pilihan {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Meminta user memilih
	for {
		nomor := readInt(reader, "Pilih: ")

		// Mengecek apakah pilihan valid
		if nomor >= 1 && nomor <= len(pilihan) {
			return pilihan[nomor-1]
		}

		fmt.Println("Pilihan tidak tersedia.")
	}
}

// Memilih beberapa item dari daftar
func pilihBanyak(reader *bufio.Reader, judul string, pilihan []string) []string {
	var hasil []string

	fmt.Println()
	fmt.Println(judul)

	for i, item := range pilihan {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	for {
		nomor := readInt(reader, "Pilih nomor motor: ")

		if nomor < 1 || nomor > len(pilihan) {
			fmt.Println("Pilihan tidak tersedia.")
			continue
		}

		// Mengambil pilihan motor
		item := pilihan[nomor-1]

		// Mengecek agar motor tidak dipilih dua kali
		sudahAda := false

		for _, x := range hasil {
			if x == item {
				sudahAda = true
				break
			}
		}

		if sudahAda {
			fmt.Println("Motor tersebut sudah dipilih.")
		} else {
			hasil = append(hasil, item)
			fmt.Println("Ditambahkan:", item)
		}

		// Menanyakan apakah ingin menambah motor lagi
		lagi := readLine(reader, "Tambah motor lagi? (y/n): ")

		if strings.ToLower(lagi) != "y" {
			break
		}
	}

	return hasil
}
