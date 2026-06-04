package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const MAX_RESEP = 100
const MAX_BAHAN = 100

type Resep struct {
	Judul string //done
	Kategori string // done
	BahanUtama string // done
	Bahan [MAX_BAHAN]string // done
	JumlahBahan int //done
	Langkah string // done
	Durasi int // done
	JumlahDicari int // done
}

type DataResep struct {
	Resep [MAX_RESEP]Resep
	JumlahResep int
}

var daftarResep DataResep
var reader = bufio.NewReader(os.Stdin)

func main() {
	var pilihan int

	for {
		pilihan = menu()

		switch pilihan {

		case 1:
			tambahResep()

		case 2:
			tampilResep()

		case 3:
			cariResep()

		case 4:
			editResep()

		case 5:
			hapusResep()

		case 6:
			fmt.Println("Urutkan Resep")

		case 7:
			fmt.Println("Statistik")

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi Resepku!😍")
			return
		}
	}
}

func inputString(prompt string) string { //buat nerima input string yang bisa pake spasi
	var teks string

	fmt.Print(prompt)
	teks, _ = reader.ReadString('\n')

	return strings.TrimSpace(teks)
}

func menu() int {
	var pilihan int

	fmt.Println()
	fmt.Println("===== RESEPKU =====")
	fmt.Println("1. Tambah Resep")
	fmt.Println("2. Tampilkan Resep")
	fmt.Println("3. Cari Resep")
	fmt.Println("4. Edit Resep")
	fmt.Println("5. Hapus Resep")
	fmt.Println("6. Urutkan Resep")
	fmt.Println("7. Statistik")
	fmt.Println("0. Keluar")

	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	reader.ReadString('\n') 
	return pilihan
}

func tambahResep() {
	var r Resep
	var i int

	if daftarResep.JumlahResep >= MAX_RESEP {
		fmt.Println("Maaf, jumlah resep sudah mencapai batas maksimum.")
		return
	}

	r.Judul = inputString("Judul Resep: ")
	r.Kategori = inputString("Kategori: ")
	r.BahanUtama = inputString("Bahan Utama: ")

	fmt.Print("Jumlah Bahan: ")
	fmt.Scan(&r.JumlahBahan)

	reader.ReadString('\n')

	for i = 0; i < r.JumlahBahan; i++ {
		r.Bahan[i] = inputString(fmt.Sprintf("Bahan %d: ", i+1))
	}
	
	r.Langkah = inputString("Langkah Memasak: ")

	fmt.Print("Durasi Memasak (menit): ")
	fmt.Scan(&r.Durasi)

	reader.ReadString('\n')

	daftarResep.Resep[daftarResep.JumlahResep] = r
	daftarResep.JumlahResep++

	fmt.Println("Resep berhasil ditambahkan!")
}

func tampilResep() {
	var i, j int

	if daftarResep.JumlahResep == 0 {
		fmt.Println("Belum ada resep yang tersimpan.")
		return
	}

	fmt.Println("\n===== DAFTAR RESEP =====")

	for i = 0; i < daftarResep.JumlahResep; i++ {
		fmt.Println("\nResep ke-", i+1)
		fmt.Println("Judul       :", daftarResep.Resep[i].Judul)
		fmt.Println("Kategori    :", daftarResep.Resep[i].Kategori)
		fmt.Println("Bahan Utama :", daftarResep.Resep[i].BahanUtama)

		fmt.Println("Daftar Bahan:")
		for j = 0; j < daftarResep.Resep[i].JumlahBahan; j++ {
			fmt.Printf("%d. %s\n", j+1, daftarResep.Resep[i].Bahan[j])
		}

		fmt.Println("Langkah     :", daftarResep.Resep[i].Langkah)
		fmt.Println("Durasi      :", daftarResep.Resep[i].Durasi, "menit")
	}
}

func cariResep() {
	var judul string
	var idx int
	var j int

	judul = inputString("Masukkan judul resep: ")

	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("Resep tidak ditemukan.")
	} else {
		fmt.Println("\n===== RESEP DITEMUKAN =====")

		fmt.Println("Judul       :", daftarResep.Resep[idx].Judul)
		fmt.Println("Kategori    :", daftarResep.Resep[idx].Kategori)
		fmt.Println("Bahan Utama :", daftarResep.Resep[idx].BahanUtama)

		fmt.Println("Daftar Bahan:")
		for j = 0; j < daftarResep.Resep[idx].JumlahBahan; j++ {
			fmt.Printf("%d. %s\n", j+1, daftarResep.Resep[idx].Bahan[j])
		}

		fmt.Println("Langkah     :", daftarResep.Resep[idx].Langkah)
		fmt.Println("Durasi      :", daftarResep.Resep[idx].Durasi, "menit")

		daftarResep.Resep[idx].JumlahDicari++
	}
}

func editResep() {
	var judul string
	var idx int

	judul = inputString("Masukkan judul resep yang ingin diedit: ")

	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("Resep tidak ditemukan.")
		return
	}

	fmt.Println("\n===== EDIT RESEP =====")

	daftarResep.Resep[idx].Judul = inputString("Judul Baru: ")
	daftarResep.Resep[idx].Kategori = inputString("Kategori Baru: ")
	daftarResep.Resep[idx].BahanUtama = inputString("Bahan Utama Baru: ")

	fmt.Println("Resep berhasil diubah!")
}

func hapusResep() {
	var judul string
	var idx, i int

	judul = inputString("Masukkan judul resep yang ingin dihapus: ")
	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("Resep tidak ditemukan.")
		return
	}

	for i = idx; i < daftarResep.JumlahResep-1; i++ {
		daftarResep.Resep[i] = daftarResep.Resep[i+1]
	}

	daftarResep.JumlahResep-- //untuk ngurangin jumlah resep setelah dihapus
	fmt.Println("Resep berhasil dihapus!")
}

func cariResepSequential(judul string) int {
	var i int

	for i = 0; i < daftarResep.JumlahResep; i++ {
		if strings.EqualFold(daftarResep.Resep[i].Judul, judul) { // cari resep biar gk usah ngeliat gede kecil huruf (string.EqualFold)
			return i
		}
	}
	return -1
}



