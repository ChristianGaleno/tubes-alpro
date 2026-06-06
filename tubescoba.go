package main

import "fmt"

type Kandidat struct {
	NoUrut int
	Nama   string
	Suara  int
}

var dataKandidat [100]Kandidat
var totalKandidat int

func main() {
	var pilihan int
	for {
		fmt.Println("\n=== MENU E-VOTING ===")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Urutkan & Tampilkan")
		fmt.Println("3. Vote (Input Suara)")
		fmt.Println("4. Lihat Statistik")
		fmt.Println("5. Cari Kandidat")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 0 {
			break
		}

		if pilihan == 1 {
			tambahKandidat()
		} else if pilihan == 2 {
			menuUrutkan()
		} else if pilihan == 3 {
			inputSuara()
		} else if pilihan == 4 {
			lihatStatistik()
		} else if pilihan == 5 {
			CariKandidat()
		}
	}
}

func CariKandidat() {
	var noUrut int
	var mid, low, high int

	if totalKandidat == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}

	fmt.Print("Masukkan nomor urut kandidat yang ingin dicari: ")
	fmt.Scan(&noUrut)
	mid = totalKandidat / 2
	low = 0
	high = totalKandidat - 1

	for low <= high {
		mid = (low + high) / 2
		if dataKandidat[mid].NoUrut == noUrut {
			fmt.Printf("Kandidat ditemukan: No Urut: %d | Nama: %s | Suara: %d\n",
				dataKandidat[mid].NoUrut,
				dataKandidat[mid].Nama,
				dataKandidat[mid].Suara)
			return
		} else if dataKandidat[mid].NoUrut < noUrut {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Kandidat tidak ditemukan.")
}

func tambahKandidat() {
	if totalKandidat >= 100 {
		fmt.Println("Data kandidat sudah penuh!")
		return
	}

	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&dataKandidat[totalKandidat].Nama)
	dataKandidat[totalKandidat].NoUrut = totalKandidat + 1
	totalKandidat++
	fmt.Println("Sip, data masuk!")
}

func menuUrutkan() {
	var pilihanUrutan int

	if totalKandidat == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}

	fmt.Println("\n--- Pilih Urutan ---")
	fmt.Println("1. Ascending (kecil ke besar)")
	fmt.Println("2. Descending (besar ke kecil)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilihanUrutan)

	if pilihanUrutan == 1 {
		urutkanKandidat(true)
		fmt.Println("Data sudah diurutkan ascending.")
	} else if pilihanUrutan == 2 {
		urutkanKandidat(false)
		fmt.Println("Data sudah diurutkan descending.")
	} else {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	tampilkanKandidat()
}

func urutkanKandidat(ascending bool) {
	for i := 0; i < totalKandidat-1; i++ {
		idxTerpilih := i

		for j := i + 1; j < totalKandidat; j++ {
			if ascending {
				if dataKandidat[j].NoUrut < dataKandidat[idxTerpilih].NoUrut {
					idxTerpilih = j
				}
			} else {
				if dataKandidat[j].NoUrut > dataKandidat[idxTerpilih].NoUrut {
					idxTerpilih = j
				}
			}
		}

		if idxTerpilih != i {
			dataKandidat[i], dataKandidat[idxTerpilih] = dataKandidat[idxTerpilih], dataKandidat[i]
		}
	}
}

func tampilkanKandidat() {
	fmt.Println("\n--- Daftar Kandidat ---")
	for i := 0; i < totalKandidat; i++ {
		fmt.Printf("No Urut: %d | Nama: %s | Suara: %d\n",
			dataKandidat[i].NoUrut,
			dataKandidat[i].Nama,
			dataKandidat[i].Suara)
	}
}

func inputSuara() {
	var nomor int
	var jumlahSuara int

	if totalKandidat == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}
	tampilkanKandidat()
	fmt.Print("Masukkan Nomor Urut Kandidat yang mau di-vote: ")
	fmt.Scan(&nomor)
	fmt.Print("Masukkan jumlah suara yang ingin ditambahkan: ")
	fmt.Scan(&jumlahSuara)

	if jumlahSuara <= 0 {
		fmt.Println("Jumlah suara harus lebih dari 0.")
		return
	}

	for i := 0; i < totalKandidat; i++ {
		if dataKandidat[i].NoUrut == nomor {
			dataKandidat[i].Suara += jumlahSuara
			fmt.Println("Berhasil vote!")
			return
		}
	}
	fmt.Println("Kandidat nggak ketemu!")
}

func lihatStatistik() {
	fmt.Println("--- Statistik Perolehan Suara ---")
	for i := 0; i < totalKandidat; i++ {
		fmt.Printf("%s: %d suara\n", dataKandidat[i].Nama, dataKandidat[i].Suara)
	}
}
