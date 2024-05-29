package main

import "fmt"

const NMAX int = 1000

type tempat struct {
	nama_tempat        string
	lokasi_tempat      string
	kapasitas_tersedia int
	harga_sewa_perjam  int
}

type pesanan struct {
	nama_user     string
	nama_tempat   string
	lokasi_tempat string
	kapasitas     int
	durasi_sewa   int
	total_harga   float64
}

type tabTempat [NMAX]tempat
type tabPesanan [NMAX]pesanan

// Penambahan Data
func tambahData(A *tabTempat, n *int) {
	if *n > NMAX {
		*n = NMAX
	}

	fmt.Println("=================================")
	fmt.Println("Masukan data tempat penyewaan : ")
	for i := 0; i < *n; i++ {
		fmt.Println("Data tempat ke-", i+1)
		fmt.Print("Nama Tempat: ")
		fmt.Scan(&A[i].nama_tempat)
		fmt.Print("Lokasi Tempat: ")
		fmt.Scan(&A[i].lokasi_tempat)
		fmt.Print("Kapasitas Tersedia: ")
		fmt.Scan(&A[i].kapasitas_tersedia)
		fmt.Print("Harga Sewa Per Jam: ")
		fmt.Scan(&A[i].harga_sewa_perjam)
	}
	fmt.Println("=================================")
}

func printDataTempat(A tabTempat, n int) {
	fmt.Println("=================================")
	fmt.Println("Data tempat penyewaan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Nama Tempat: %s, Lokasi Tempat: %s, Kapasitas Tersedia: %d, Harga Sewa Per Jam: %d\n",
			A[i].nama_tempat, A[i].lokasi_tempat, A[i].kapasitas_tersedia, A[i].harga_sewa_perjam)
	}
	fmt.Println("=================================")
}

// Pencarian data menggunakan Binary Search
func findData(A tabTempat, n int, input *string) int {
	for i := 0; i < n; i++ {
		if A[i].nama_tempat == *input {
			return i // Mengembalikan indeks jika ditemukan
		}
	}
	return -1
}

// Proses Pengubahan data tempat penyewaan acara
func editData(A tabTempat, n int, input *string, input_2 *int) {

	// variable data_ubah adalah variable yang merepresentasikan data mana yang kan diubah
	var data_ubah string

	fmt.Print("Masukan Data yang akan diubah (nama, lokasi, kapasitas, harga) / masukan tidak jika \n tidak ingin mengubah data: ")
	fmt.Scan(&data_ubah)

	if data_ubah == "nama" {
		// variable x adalah data untuk mengubah nama dan lokasi
		var x string

		fmt.Print("Masukan nama yang akan diubah: ")
		fmt.Scan(&x)

		result := findData(A, n, input)

		if result != -1 {
			A[result].nama_tempat = x
			fmt.Println("Data Nama Tempat berhasil diubah")
		} else {
			fmt.Print("Data nama tempat tidak ditemukan!")
		}

	} else if data_ubah == "lokasi" {
		// variable x adalah data untuk mengubah nama dan lokasi
		var y string

		fmt.Print("Masukan lokasi yang akan diubah: ")
		fmt.Scan(&y)

		result := findData(A, n, input)

		if result != -1 {
			A[result].lokasi_tempat = y
			fmt.Println("Data lokasi Tempat berhasil diubah")
			
		} else {
			fmt.Print("Data lokasi tempat tidak ditemukan!")
		}

	} else if data_ubah == "kapasistas" {

		fmt.Print("Masukan kapasitas yang akan diubah: ")
		fmt.Scan(&*input_2)

		result := findData(A, n, input)

		if result != -1 {
			A[result].kapasitas_tersedia = *input_2
			fmt.Println("Data Kapasitas Tempat berhasil diubah")
		} else {
			fmt.Print("Data Kapasitas tempat tidak ditemukan!")
		}

	} else if data_ubah == "harga" {

		fmt.Print("Masukan harga yang akan diubah: ")
		fmt.Scan(&*input_2)

		result := findData(A, n, input)

		if result != -1 {
			A[result].harga_sewa_perjam = *input_2
			fmt.Println("Data Harga Sewa Tempat berhasil diubah")
		} else {
			fmt.Print("Data Harga Sewa tempat tidak ditemukan!")
		}

	}

}

// Proses penghapusan data tempat penyewaan acara
func hapusData(A tabTempat, n int, input string) {
	// var x string

	// kondisi := findData(A, n, x)

	// if kondisi != -1 {
	// 	fmt.Print("ADA")
	// } else {
	// 	fmt.Print("TIDAK ADA")
	// }

}

// Proses Penyewaan temppat penyewaan acara
func sewaTempat(B *tabPesanan) {

}

func main() {
	var data_tempat tabTempat
	var data_pesan tabPesanan
	var nData, query_2 int
	var status, query string

	fmt.Print("Masukan jumlah data yang akan dimasukkan: ")
	fmt.Scan(&nData)

	// Memasukan Data Tempat Penyewaan
	tambahData(&data_tempat, &nData)

	// Tampilkan Semua Data
	printDataTempat(data_tempat, nData)

	// Memasukan status Pengubahan, penghapusan, dan penyewaan
	fmt.Println("Apakah ingin merubah data ?")
	fmt.Print("Masukan (edit) jika ingin mengubah, Masukan (hapus) jika ingin menghapus : ")
	fmt.Scan(&status)
	fmt.Print("Masukan id untuk keperluan pengubahan dan penghapusan data : ")
	fmt.Scan(&query)

	if status == "edit" {

		editData(data_tempat, nData, &query, &query_2)

	} else if status == "hapus" {

		hapusData(data_tempat, nData, query)


	} else if status == "sewa" {

		sewaTempat(&data_pesan)

	}

	//urutkan data secara ASC menggunakan insertion
	fmt.Print("Apakah data mau diurutkan ? (ya/tidak): ")
	
}
