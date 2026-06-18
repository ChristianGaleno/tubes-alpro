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
		fmt.Println("6. Hapus Kandidat")
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
		}else if pilihan == 6 {
			hapusKandidat()
		}
	}
}

func hapusKandidat(){
	var kandidat int
	var removed bool
	var i int
	fmt.Println("Kandidat yang ingin dihapus: ")
	fmt.Scan(&kandidat)
	removed = true

	for i < totalKandidat && removed {
		if dataKandidat[i].NoUrut == kandidat{
			for j:=i;j<totalKandidat-1;j++{
				dataKandidat[j] = dataKandidat[j+1]
				dataKandidat[j].NoUrut -= 1
			}
			removed = false
			totalKandidat -= 1
		}
		i++
	}

	fmt.Printf("Data Kandidat: %d Berhasil dihapus", kandidat)

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
	
	var data [100]Kandidat
	for i := 0; i < totalKandidat; i++ {
		data[i] = dataKandidat[i]
	}
	for i := 0; i < totalKandidat-1; i++ {
		idxMin := i
		for j := i + 1; j < totalKandidat; j++ {
			if data[j].NoUrut < data[idxMin].NoUrut {
				idxMin = j
			}
		}
		if idxMin != i {
			data[i], data[idxMin] = data[idxMin], data[i]
		}
	}

	low = 0
	high = totalKandidat - 1

	for low <= high {
		mid = (low + high) / 2
		if data[mid].NoUrut == noUrut {
			fmt.Printf("Kandidat ditemukan: No Urut: %d | Nama: %s | Suara: %d\n",
				data[mid].NoUrut,
				data[mid].Nama,
				data[mid].Suara)
			return
		} else if data[mid].NoUrut < noUrut {
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
	var totalSuara int
	
	for i := 0; i < totalKandidat; i++ {
		totalSuara += dataKandidat[i].Suara
	} 

	fmt.Println("--- Statistik Perolehan Suara ---")
	for i := 0; i < totalKandidat; i++ {
		var persentase float64
		if totalSuara > 0 {
			persentase = float64(dataKandidat[i].Suara) / float64(totalSuara) * 100
		}
		fmt.Printf("%s: %d suara | persentase: %.2f \n", dataKandidat[i].Nama, dataKandidat[i].Suara, persentase)
	}
}
