package main

import "fmt"

const NMAX int = 1000

type Tempat struct {
	nama_tempat        string
	lokasi_tempat      string
	kapasitas_tersedia int
	harga_sewa_perjam  int
}

type TabTempat [NMAX]Tempat

func (t Tempat) LessThan(other Tempat) bool {
	return t.nama_tempat < other.nama_tempat
}

// Penambahan Data
func tambahData(A *TabTempat, n *int) {
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

func printDataTempat(A TabTempat, n int) {
	fmt.Println("=================================")
	fmt.Println("Data tempat penyewaan:")
	for i := 0; i < n; i++ {
		fmt.Printf("Nama Tempat: %s, Lokasi Tempat: %s, Kapasitas Tersedia: %d, Harga Sewa Per Jam: %d\n",
			A[i].nama_tempat, A[i].lokasi_tempat, A[i].kapasitas_tersedia, A[i].harga_sewa_perjam)
	}
	fmt.Println("=================================")
}

// Pencarian data menggunakan Linear Search
func findData(A TabTempat, n int, input string) int {
	for i := 0; i < n; i++ {
		if A[i].nama_tempat == input {
			return i // Mengembalikan indeks jika ditemukan
		}
	}
	return -1
}

// Proses Pengubahan data tempat penyewaan acara
func editData(A *TabTempat, n int, input string) {
	var data_ubah string

	fmt.Print("Masukan Data yang akan diubah (nama, lokasi, kapasitas, harga) / masukan tidak jika tidak ingin mengubah data: ")
	fmt.Scan(&data_ubah)

	switch data_ubah {
	case "nama":
		var x string
		fmt.Print("Masukan nama yang akan diubah: ")
		fmt.Scan(&x)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].nama_tempat = x
			fmt.Println("Data Nama Tempat berhasil diubah")
		} else {
			fmt.Println("Data nama tempat tidak ditemukan!")
		}
	case "lokasi":
		var y string
		fmt.Print("Masukan lokasi yang akan diubah: ")
		fmt.Scan(&y)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].lokasi_tempat = y
			fmt.Println("Data lokasi Tempat berhasil diubah")
		} else {
			fmt.Println("Data lokasi tempat tidak ditemukan!")
		}
	case "kapasitas":
		var kapasitas int
		fmt.Print("Masukan kapasitas yang akan diubah: ")
		fmt.Scan(&kapasitas)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].kapasitas_tersedia = kapasitas
			fmt.Println("Data Kapasitas Tempat berhasil diubah")
		} else {
			fmt.Println("Data Kapasitas tempat tidak ditemukan!")
		}
	case "harga":
		var harga int
		fmt.Print("Masukan harga yang akan diubah: ")
		fmt.Scan(&harga)
		result := findData(*A, n, input)
		if result != -1 {
			(*A)[result].harga_sewa_perjam = harga
			fmt.Println("Data Harga Sewa Tempat berhasil diubah")
		} else {
			fmt.Println("Data Harga Sewa tempat tidak ditemukan!")
		}
	default:
		fmt.Println("Data tidak valid atau tidak ada perubahan yang dilakukan.")
	}
}

// Proses penghapusan data tempat penyewaan acara
func hapusData(A *TabTempat, n *int, input string) {
	index := findData(*A, *n, input)
	if index != -1 {
		for i := index; i < *n-1; i++ {
			(*A)[i] = (*A)[i+1]
		}
		*n--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

// Proses Penyewaan tempat penyewaan acara
func sewaTempat(A TabTempat, n int) {
	var durasi int
	var input string

	fmt.Println("=================================")
	fmt.Print("Masukan durasi penyewaan tempat acara (jam): ")
	fmt.Scan(&durasi)
	fmt.Print("Masukan Nama Tempat yang ingin disewa: ")
	fmt.Scan(&input)

	result := findData(A, n, input)

	if result != -1 {
		totalHarga := A[result].harga_sewa_perjam * durasi
		fmt.Printf("Tempat %s sudah disewa selama %d jam, total harga menjadi %d\n", A[result].nama_tempat, durasi, totalHarga)
		fmt.Println("=================================")
	} else {
		fmt.Println("Data Nama Tempat tidak ditemukan!")
	}
}

func ascSort(A *TabTempat, n int) {
	for i := 1; i < n; i++ {
		key := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].nama_tempat > key.nama_tempat {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}

func main() {
	var data_tempat TabTempat
	var nData int

	fmt.Print("Masukan jumlah data yang akan dimasukkan: ")
	fmt.Scan(&nData)

	// Memasukan Data Tempat Penyewaan
	tambahData(&data_tempat, &nData)

	// Tampilkan Semua Data
	fmt.Println("Data Awal Sebelum ada perubahan")
	printDataTempat(data_tempat, nData)

	var status, query string

	// Memasukan status Pengubahan, penghapusan, dan penyewaan
	fmt.Print("Apakah ingin merubah data? (edit/hapus/sewa): ")
	fmt.Scan(&status)
	fmt.Print("Masukan nama tempat untuk keperluan pengubahan data atau penghapusan data, penyewaan: ")
	fmt.Scan(&query)

	switch status {
	case "edit":
		editData(&data_tempat, nData, query)
	case "hapus":
		hapusData(&data_tempat, &nData, query)
	case "sewa":
		sewaTempat(data_tempat, nData)
	default:
		fmt.Println("Perintah tidak dikenali")
	}

	// Pengurutan data
	ascSort(&data_tempat, nData)

	// Tampilkan Data yang Sudah Diurutkan
	printDataTempat(data_tempat, nData)
}