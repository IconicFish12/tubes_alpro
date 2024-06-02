package main

import "fmt"

const NMAX int = 1000

type tempat struct {
	nama_tempat        string
	lokasi_tempat      string
	kapasitas_tersedia int
	harga_sewa_perjam  int
}

type tabTempat [NMAX]tempat

func (t tempat) LessThan(other tempat) bool {
	return t.nama_tempat < other.nama_tempat
}

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
func editData(A *tabTempat, n int, input *string, input_2 *int) {

	var data_ubah string

	fmt.Print("Masukan Data yang akan diubah (nama, lokasi, kapasitas, harga) / masukan tidak jika \n tidak ingin mengubah data: ")
	fmt.Scan(&data_ubah)

	if data_ubah == "nama" {
		var x string
		fmt.Print("Masukan nama yang akan diubah: ")
		fmt.Scan(&x)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].nama_tempat = x
			fmt.Println("Data Nama Tempat berhasil diubah")
		} else {
			fmt.Print("Data nama tempat tidak ditemukan!")
		}
	} else if data_ubah == "lokasi" {
		var y string
		fmt.Print("Masukan lokasi yang akan diubah: ")
		fmt.Scan(&y)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].lokasi_tempat = y
			fmt.Println("Data lokasi Tempat berhasil diubah")
		} else {
			fmt.Print("Data lokasi tempat tidak ditemukan!")
		}
	} else if data_ubah == "kapasitas" {
		fmt.Print("Masukan kapasitas yang akan diubah: ")
		fmt.Scan(input_2)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].kapasitas_tersedia = *input_2
			fmt.Println("Data Kapasitas Tempat berhasil diubah")
		} else {
			fmt.Print("Data Kapasitas tempat tidak ditemukan!")
		}
	} else if data_ubah == "harga" {
		fmt.Print("Masukan harga yang akan diubah: ")
		fmt.Scan(input_2)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].harga_sewa_perjam = *input_2
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
func sewaTempat(A tabTempat, n int) {

	var durasi, hasil int
	var input string

	fmt.Println("=================================")
	fmt.Print("Masukan durasi penyewaan tempat acara / jam : ")
	fmt.Scan(&durasi)
	fmt.Print("Masukan Nama Tempat yang ingin disewa : ")
	fmt.Scan(&input)

	result := findData(A, n, &input)

	if result != -1 {
		hasil = A[result].harga_sewa_perjam * durasi
		fmt.Printf("Tempat %s sudah disewa selama %d jam, total harga menjadi %d", A[result].nama_tempat, durasi, hasil)
		fmt.Println("=================================")

	} else {
		fmt.Print("Data Nama Tempat tempat tidak ditemukan!")
	}

}

func ascSort(A *tabTempat, n int) {
	for i := 1; i < n; i++ {
		key := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].LessThan(key) {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}


func main() {
	var data_tempat tabTempat
	var nData, query_2 int
	var status, query string

	fmt.Print("Masukan jumlah data yang akan dimasukkan: ")
	fmt.Scan(&nData)

	// Memasukan Data Tempat Penyewaan
	tambahData(&data_tempat, &nData)

	// Tampilkan Semua Data
	fmt.Println("Data Awal Sebelum ada perubahan")
	printDataTempat(data_tempat, nData)

	// Memasukan status Pengubahan, penghapusan, dan penyewaan
	fmt.Println("Apakah ingin merubah data ?")
	fmt.Print("Masukan (edit) jika ingin mengubah, Masukan (hapus) jika ingin menghapus : ")
	fmt.Scan(&status)
	fmt.Print("Masukan nama tempat untuk keperluan pengubahan data atau penghapusan data, penyewaan : ")
	fmt.Scan(&query)

	if status == "edit" {

		editData(&data_tempat, nData, &query, &query_2)

	} else if status == "hapus" {

		hapusData(data_tempat, nData, query)

	} else if status == "sewa" {

		sewaTempat(data_tempat, nData)

		return
	}

	//Pengurutan data
	ascSort(&data_tempat, nData)

	printDataTempat(data_tempat, nData)

	//urutkan data secara ASC menggunakan insertion
	// fmt.Print("Apakah data mau diurutkan ? (ya/tidak): ")

}
