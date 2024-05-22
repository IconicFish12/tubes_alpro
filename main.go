package main

import "fmt"

const NMAX int = 1000

type tempat struct {
	id_tempat          int
	nama_tempat        string
	lokasi_tempat      string
	kapasitas_tersedia int
	kapasitas_maksimum int
	harga_sewa_perjam  int
}

type pesanan struct {
	id_pemesanan  int
	nama_user     string
	nama_tempat   string
	lokasi_tempat string
	kapasitas     int
	total_harga   float64
}

type tabTempat [NMAX]tempat
type tabPesanan [NMAX]pesanan

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
		fmt.Print("Kapasitas Maksimum: ")
		fmt.Scan(&A[i].kapasitas_maksimum)
		fmt.Print("Harga Sewa Per Jam: ")
		fmt.Scan(&A[i].harga_sewa_perjam)
		A[i].id_tempat = i + 1
	}
	fmt.Println("=================================")
}

func printData(A tabTempat, n int) {
	fmt.Println("=================================")
	fmt.Println("Data tempat penyewaan:")
	for i := 0; i < n; i++ {
		fmt.Printf("ID Tempat: %d, Nama Tempat: %s, Lokasi Tempat: %s, Kapasitas Tersedia: %d, Kapasitas Maksimum: %d, Harga Sewa Per Jam: %d\n",
			A[i].id_tempat, A[i].nama_tempat, A[i].lokasi_tempat, A[i].kapasitas_tersedia, A[i].kapasitas_maksimum, A[i].harga_sewa_perjam)
	}
	fmt.Println("=================================")
}

// Pencarian data menggunakan Binary Search
func findData(A tabTempat, n int, input string) int {
	return -1
}


// Proses Pengubahan data tempat penyewaan acara
func editData(A tabTempat, n int, idData int, input string) {

	// kondisi := findData(A, n, input)

}


// Proses penghapusan data tempat penyewaan acara
func hapusData(A tabTempat, n int, idData int, input string) {

}


// Proses Penyewaan temppat penyewaan acara
func sewaTempat() {

}

func main() {
	var data_tempat tabTempat
	// var pemesanan tabPesanan
	var nData, id_data, query_2 int
	var status, query string

	fmt.Print("Masukan jumlah data yang akan dimasukkan: ")
	fmt.Scan(&nData)

	// Memasukan Data Tempat Penyewaan
	tambahData(&data_tempat, &nData)

	// Tampilkan Semua Data
	printData(data_tempat, nData)

	// Memasukan status Pengubahan dan penghapusan data
	fmt.Println("Apakah ingin merubah data ?")
	fmt.Print("Masukan (edit) jika ingin mengubah, Masukan (hapus) jika ingin menghapus : ")
	fmt.Scan(&status)
	fmt.Print("Masukan id, query 1 berupa string(nama dan lokasi), query 2 berupa integer(kapasitas) : ")
	fmt.Scan(&id_data, &query, &query_2)



	if status == "edit" {


	} else if status == "hapus" {

	} else if status == "sewa"  {

	}
}
