package main

import (
	"fmt"
	"os"
	"strconv"
)

type Classmate struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	//get args
	args := os.Args
	if len(args) != 2 {
		fmt.Println("masukkan go run biodata.go <no_absen>")
		return
	}

	//convert args to int
	no_absen, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("masukkan angka")
		return
	}

	//contoh data
	dataClassmate := []Classmate{
		{
			Nama:      "Ade",
			Alamat:    "Jalan Medan Merdeka Barat",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin belajar bahasa pemrograman baru",
		},
		{
			Nama:      "Rizky",
			Alamat:    "Jalan Medan Merdeka Barat",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Tertarik dengan bahasa pemrograman baru",
		},
	}

	// Mencari data teman berdasarkan absen
	if no_absen <= len(dataClassmate) {
		teman := dataClassmate[no_absen-1]
		// Menampilkan data teman
		fmt.Println("Nama:", teman.Nama)
		fmt.Println("Alamat:", teman.Alamat)
		fmt.Println("Pekerjaan:", teman.Pekerjaan)
		fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
	} else {
		fmt.Println("Nomor absen tidak ditemukan")
	}

}
