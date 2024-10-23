package main

import (
	"fmt"
)

// Mendefinisikan tipe bentukan untuk data nasabah
type nasabah struct {
	idNasabah   int
	namaNasabah string
	namaBank    string
	noRekening  string
}

// Mendefinisikan tipe array untuk menyimpan data nasabah dengan kapasitas 2022
type tabNasabah [2022]nasabah

// Prosedur untuk menambahkan nasabah baru
func addNasabah(T *tabNasabah, N *int) {
	if *N >= 2022 {
		fmt.Println("Data penuh")
		return
	}
	var nasabahBaru nasabah
	fmt.Print("Masukkan ID Nasabah: ")
	fmt.Scan(&nasabahBaru.idNasabah)
	fmt.Print("Masukkan Nama Nasabah: ")
	fmt.Scan(&nasabahBaru.namaNasabah)
	fmt.Print("Masukkan Nama Bank: ")
	fmt.Scan(&nasabahBaru.namaBank)
	fmt.Print("Masukkan Nomor Rekening: ")
	fmt.Scan(&nasabahBaru.noRekening)
	T[*N] = nasabahBaru
	(*N)++
}

// Prosedur untuk mencetak data nasabah berdasarkan nama bank
func cetak(T tabNasabah, N int, X string) {
	for i := 0; i < N; i++ {
		if T[i].namaBank == X {
			fmt.Printf("ID Nasabah: %d, Nama Nasabah: %s, Nama Bank: %s, Nomor Rekening: %s\n",
				T[i].idNasabah, T[i].namaNasabah, T[i].namaBank, T[i].noRekening)
		}
	}
}

func main() {
	var nasabahList tabNasabah
	var jumlahNasabah int

	for {
		var pilihan int
		fmt.Println("Menu:")
		fmt.Println("1. Tambah Nasabah")
		fmt.Println("2. Cetak Nasabah Berdasarkan Bank")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			addNasabah(&nasabahList, &jumlahNasabah)
		case 2:
			var namaBank string
			fmt.Print("Masukkan nama bank: ")
			fmt.Scan(&namaBank)
			cetak(nasabahList, jumlahNasabah, namaBank)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
