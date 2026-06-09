package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

const MAX_RESEP = 100
const MAX_BAHAN = 100
const MAX_LANGKAH = 100

type Resep struct {
	Judul string //done
	Kategori string // done
	BahanUtama string // done
	Bahan [MAX_BAHAN]string // done
	JumlahBahan int //done
	Langkah [MAX_LANGKAH]string // done
	JumlahLangkah int // done
	Durasi int // done
	JumlahDicari int // done
}

type DataResep struct {
	Resep [MAX_RESEP]Resep
	JumlahResep int
}

// Variabel Global 
var daftarResep DataResep
var reader = bufio.NewReader(os.Stdin)

func main() {
	var pilihan int
	tampilanLuculucuan()

	for {
		pilihan = menu()

		switch pilihan {

		case 1:
			tambahResep()

		case 2:
			tampilResep()

		case 3:
			cariResepMenu()

		case 4:
			editResep()

		case 5:
			hapusResep()

		case 6:
			urutkanResep()

		case 7:
			statistik()

		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi Resepku!😍")
			return
		}
	}
}

// INPUT 
func inputString(prompt string) string { //buat nerima input string yang bisa pake spasi
	var teks string

	fmt.Print(prompt)
	teks, _ = reader.ReadString('\n')

	return strings.TrimSpace(teks)
}

// MENU
func menu() int {
	var pilihan int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         🍳 RESEPKU APP 🍳")
	fmt.Println("========================================")
	fmt.Println(" 1. ➕ Tambah Resep")
	fmt.Println(" 2. 📖 Tampilkan Resep")
	fmt.Println(" 3. 🔍 Cari Resep")
	fmt.Println(" 4. ✏️  Edit Resep")
	fmt.Println(" 5. 🗑️  Hapus Resep")
	fmt.Println(" 6. ↕️  Urutkan Resep")
	fmt.Println(" 7. 📊 Statistik")
	fmt.Println(" 0. 🚪 Keluar")
	fmt.Println("========================================")

	fmt.Print("👉 Pilih menu: ")
	fmt.Scan(&pilihan)

	reader.ReadString('\n')
	return pilihan
}

// CRUD RESEP
func tambahResep() {
	var r Resep
	var i int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         ➕ TAMBAH RESEP")
	fmt.Println("========================================")

	if daftarResep.JumlahResep >= MAX_RESEP {
		fmt.Println("❌ Maaf, jumlah resep sudah mencapai batas maksimum.")
		return
	}

	r.Judul = inputString("Judul Resep: ")
	r.Kategori = inputString("Kategori: ")
	r.BahanUtama = inputString("Bahan Utama: ")

	fmt.Println()
	fmt.Println("📝 Tambahkan Bahan")
	fmt.Println("(Kosongkan input lalu tekan Enter jika sudah selesai)")
	fmt.Println()

	for i = 0; i < MAX_BAHAN; i++ {

		r.Bahan[i] = inputString(
			fmt.Sprintf("Bahan %d: ", i+1),
		)

		if r.Bahan[i] == "" {
			break
		}

		r.JumlahBahan++
	}

	fmt.Printf("✅ %d bahan berhasil ditambahkan\n", r.JumlahBahan)

	fmt.Println()

	fmt.Println()
	fmt.Println("👨‍🍳 Tambahkan Langkah Memasak")
	fmt.Println("(Kosongkan input lalu tekan Enter jika sudah selesai)")
	fmt.Println()

	for i = 0; i < MAX_LANGKAH; i++ {

		r.Langkah[i] = inputString(
			fmt.Sprintf("Langkah %d: ", i+1),
		)

		if r.Langkah[i] == "" {
			break
		}

		r.JumlahLangkah++
	}

	fmt.Printf("✅ %d langkah berhasil ditambahkan\n", r.JumlahLangkah)

	fmt.Print("Durasi Memasak (menit): ")
	fmt.Scan(&r.Durasi)

	reader.ReadString('\n')

	daftarResep.Resep[daftarResep.JumlahResep] = r
	daftarResep.JumlahResep++

	fmt.Println()
	fmt.Println("✅ Resep berhasil ditambahkan!")
}

func tampilResep() {
	var i, j int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         📖 DAFTAR RESEP")
	fmt.Println("========================================")

	if daftarResep.JumlahResep == 0 {
		fmt.Println("📭 Belum ada resep yang tersimpan.")
		return
	}

	for i = 0; i < daftarResep.JumlahResep; i++ {

		fmt.Println()
		fmt.Println("========================================")
		fmt.Printf("📖 RESEP #%d\n", i+1)
		fmt.Println("========================================")

		fmt.Println("🍽️  Judul       :", daftarResep.Resep[i].Judul)
		fmt.Println("📂  Kategori    :", daftarResep.Resep[i].Kategori)
		fmt.Println("🥘  Bahan Utama :", daftarResep.Resep[i].BahanUtama)

		fmt.Println()
		fmt.Println("📝 Daftar Bahan:")

		for j = 0; j < daftarResep.Resep[i].JumlahBahan; j++ {
			fmt.Printf("   %d. %s\n",
				j+1,
				daftarResep.Resep[i].Bahan[j],
			)
		}

		fmt.Println()
		fmt.Println("👨‍🍳 Langkah Memasak:")

		for j = 0; j < daftarResep.Resep[i].JumlahLangkah; j++ {
			fmt.Printf("   %d. %s\n",
				j+1,
				daftarResep.Resep[i].Langkah[j],
			)
		}

		fmt.Println()
		fmt.Println("⏱️  Durasi :", daftarResep.Resep[i].Durasi, "menit")
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("📚 Total Resep :", daftarResep.JumlahResep)
	fmt.Println("========================================")
}

func editResep() {
	var judul string
	var idx int

	judul = inputString("Masukkan judul resep yang ingin diedit: ")

	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("⚠️ Resep tidak ditemukan.")
		return
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         ✏️ EDIT RESEP")
	fmt.Println("========================================")

	daftarResep.Resep[idx].Judul = inputString("Judul Baru: ")
	daftarResep.Resep[idx].Kategori = inputString("Kategori Baru: ")
	daftarResep.Resep[idx].BahanUtama = inputString("Bahan Utama Baru: ")

	fmt.Println()
	fmt.Println("📝 Edit Daftar Bahan")
	fmt.Println("(Kosongkan input lalu tekan Enter jika sudah selesai)")
	fmt.Println()

	daftarResep.Resep[idx].JumlahBahan = 0 // ini buat reset jadi pas edit bener2 kehapus yang sebelumnya

	for i := 0; i < MAX_BAHAN; i++ {

		bahan := inputString(
			fmt.Sprintf("Bahan %d: ", i+1),
		)

		if bahan == "" {
			break
		}

		daftarResep.Resep[idx].Bahan[i] = bahan
		daftarResep.Resep[idx].JumlahBahan++
	}

	fmt.Println()
	fmt.Println("👨‍🍳 Edit Langkah Memasak")
	fmt.Println("(Kosongkan input lalu tekan Enter jika sudah selesai)")
	fmt.Println()

	daftarResep.Resep[idx].JumlahLangkah = 0 // reset langkah lama supaya terganti dengan langkah baru

	for i := 0; i < MAX_LANGKAH; i++ {

		langkah := inputString(
			fmt.Sprintf("Langkah %d: ", i+1),
		)

		if langkah == "" {
			break
		}

		daftarResep.Resep[idx].Langkah[i] = langkah
		daftarResep.Resep[idx].JumlahLangkah++
	}

	fmt.Println()

	fmt.Print("⏱️ Durasi Baru (menit): ")
	fmt.Scan(&daftarResep.Resep[idx].Durasi)

	reader.ReadString('\n') // membersihkan enter setelah input angka

	fmt.Println()
	fmt.Println("✅ Resep berhasil diubah!")
}

func hapusResep() {
	var judul string
	var idx, i int

	judul = inputString("Masukkan judul resep yang ingin dihapus: ")
	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("⚠️ Resep tidak ditemukan.")
		return
	}

	for i = idx; i < daftarResep.JumlahResep-1; i++ {
		daftarResep.Resep[i] = daftarResep.Resep[i+1]
	}

	daftarResep.JumlahResep-- //untuk ngurangin jumlah resep setelah dihapus
	fmt.Println("Resep berhasil dihapus!")
}
// END CRUD RESEP

// MENU SEARCHING START

func cariResepMenu() {
	var pilihan int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         🔍 CARI RESEP")
	fmt.Println("========================================")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")

	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	reader.ReadString('\n')

	switch pilihan {

	case 1:
		cariResepSequentialMenu()

	case 2:
		cariResepBinaryMenu()

	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func cariResepSequentialMenu() {
	var judul string
	var idx, j int

	judul = inputString("Masukkan judul resep: ")

	idx = cariResepSequential(judul)

	if idx == -1 {

		fmt.Println("⚠️ Resep tidak ditemukan.")

	} else {

		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("       🎯 RESEP DITEMUKAN")
		fmt.Println("========================================")

		fmt.Println("🍽️  Judul       :", daftarResep.Resep[idx].Judul)
		fmt.Println("📂  Kategori    :", daftarResep.Resep[idx].Kategori)
		fmt.Println("🥘  Bahan Utama :", daftarResep.Resep[idx].BahanUtama)

		fmt.Println()
		fmt.Println("📝 Daftar Bahan:")

		for j = 0; j < daftarResep.Resep[idx].JumlahBahan; j++ {
			fmt.Printf("   %d. %s\n",
				j+1,
				daftarResep.Resep[idx].Bahan[j],
			)
		}

		fmt.Println()
		fmt.Println("👨‍🍳 Langkah Memasak:")

		for j = 0; j < daftarResep.Resep[idx].JumlahLangkah; j++ {
			fmt.Printf("   %d. %s\n",
				j+1,
				daftarResep.Resep[idx].Langkah[j],
			)
		}

		fmt.Println()
		fmt.Println("⏱️  Durasi :", daftarResep.Resep[idx].Durasi, "menit")

		daftarResep.Resep[idx].JumlahDicari++
	}
}

func cariResepBinaryMenu() {
	var judul string
	var idx, j int

	selectionSortJudul(true)

	judul = inputString("Masukkan judul resep: ")

	idx = cariResepBinary(judul)

	if idx == -1 {

		fmt.Println("⚠️ Resep tidak ditemukan.")

	} else {

		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("       🎯 RESEP DITEMUKAN")
		fmt.Println("========================================")

		fmt.Println("🍽️  Judul       :", daftarResep.Resep[idx].Judul)
		fmt.Println("📂  Kategori    :", daftarResep.Resep[idx].Kategori)
		fmt.Println("🥘  Bahan Utama :", daftarResep.Resep[idx].BahanUtama)

		fmt.Println()
		fmt.Println("📝 Daftar Bahan:")

		for j = 0; j < daftarResep.Resep[idx].JumlahBahan; j++ {
			fmt.Printf(
				"   %d. %s\n",
				j+1,
				daftarResep.Resep[idx].Bahan[j],
			)
		}

		fmt.Println()
		fmt.Println("👨‍🍳 Langkah Memasak:")

		for j = 0; j < daftarResep.Resep[idx].JumlahLangkah; j++ {
			fmt.Printf(
				"   %d. %s\n",
				j+1,
				daftarResep.Resep[idx].Langkah[j],
			)
		}

		fmt.Println()
		fmt.Println("⏱️  Durasi :", daftarResep.Resep[idx].Durasi, "menit")

		daftarResep.Resep[idx].JumlahDicari++
	}
}

// JANGAN DI APA APAIN, (GUAJUGA GATAU KENAPA INI BISA WORK HEHEHHEHEHE "if it works don't touch it") MALAS BACA DOKUMENTASI, JADI GUA BUAT SENDIRI, KALO ADA YANG GAK JELAS BISA TANYA AJA
func cariResepSequential(judul string) int {
	var i int

	for i = 0; i < daftarResep.JumlahResep; i++ {
		if strings.EqualFold(daftarResep.Resep[i].Judul, judul) { // cari resep biar gk usah ngeliat gede kecil huruf (string.EqualFold)
			return i
		}
	}
	return -1
}

func cariResepBinary(judul string) int {
	var kiri, kanan, tengah int

	kiri = 0
	kanan = daftarResep.JumlahResep - 1

	for kiri <= kanan {

		tengah = (kiri + kanan) / 2

		if strings.EqualFold(daftarResep.Resep[tengah].Judul, judul) {
			return tengah
		}

		if strings.ToLower(judul) <
			strings.ToLower(daftarResep.Resep[tengah].Judul) {

			kanan = tengah - 1

		} else {

			kiri = tengah + 1
		}
	}

	return -1
}

// MENU SEARCHING END

// MENU SORTING START
func urutkanResep() {
	var pilihan int

	fmt.Println("\n===== URUTKAN RESEP =====")
	fmt.Println("1. Judul Ascending (Selection Sort)")
	fmt.Println("2. Judul Descending (Selection Sort)")
	fmt.Println("3. Durasi Ascending (Insertion Sort)")
	fmt.Println("4. Durasi Descending (Insertion Sort)")

	fmt.Print("Pilih: ")
	fmt.Scan(&pilihan)

	reader.ReadString('\n')

	switch pilihan {

	case 1:
		selectionSortJudul(true)
		fmt.Println("Resep berhasil diurutkan berdasarkan judul A-Z")

	case 2:
		selectionSortJudul(false)
		fmt.Println("Resep berhasil diurutkan berdasarkan judul Z-A")

	case 3:
		insertionSortDurasi(true)
		fmt.Println("Resep berhasil diurutkan berdasarkan durasi terkecil")

	case 4:
		insertionSortDurasi(false)
		fmt.Println("Resep berhasil diurutkan berdasarkan durasi terbesar")

	default:
		fmt.Println("Pilihan tidak valid")
	}
}

func selectionSortJudul(ascending bool) {
	var i, j, idx int
	var temp Resep

	for i = 0; i < daftarResep.JumlahResep-1; i++ {
		idx = i

		for j = i + 1; j < daftarResep.JumlahResep; j++ {

			if ascending {
				if strings.ToLower(daftarResep.Resep[j].Judul) <
					strings.ToLower(daftarResep.Resep[idx].Judul) {
					idx = j
				}
			} else {
				if strings.ToLower(daftarResep.Resep[j].Judul) >
					strings.ToLower(daftarResep.Resep[idx].Judul) {
					idx = j
				}
			}
		}

		temp = daftarResep.Resep[i]
		daftarResep.Resep[i] = daftarResep.Resep[idx]
		daftarResep.Resep[idx] = temp
	}
}

func insertionSortDurasi(ascending bool) {
	var i, j int
	var temp Resep

	for i = 1; i < daftarResep.JumlahResep; i++ {
		temp = daftarResep.Resep[i]
		j = i - 1

		if ascending {

			for j >= 0 && daftarResep.Resep[j].Durasi > temp.Durasi {
				daftarResep.Resep[j+1] = daftarResep.Resep[j]
				j--
			}

		} else {

			for j >= 0 && daftarResep.Resep[j].Durasi < temp.Durasi {
				daftarResep.Resep[j+1] = daftarResep.Resep[j]
				j--
			}

		}

		daftarResep.Resep[j+1] = temp
	}
}
// MENU SORTING END

// MENU STATISTIK START

func statistik() {
	var i, totalDurasi int
	var idxTercepat, idxTerlama, idxPopuler int
	var rataRata float64

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         📊 STATISTIK RESEP")
	fmt.Println("========================================")

	if daftarResep.JumlahResep == 0 {
		fmt.Println("Belum ada resep yang tersimpan.")
		return
	}

	for i = 0; i < daftarResep.JumlahResep; i++ {

		totalDurasi += daftarResep.Resep[i].Durasi

		if daftarResep.Resep[i].Durasi <
			daftarResep.Resep[idxTercepat].Durasi {
			idxTercepat = i
		}

		if daftarResep.Resep[i].Durasi >
			daftarResep.Resep[idxTerlama].Durasi {
			idxTerlama = i
		}

		if daftarResep.Resep[i].JumlahDicari >
			daftarResep.Resep[idxPopuler].JumlahDicari {
			idxPopuler = i
		}
	}

	rataRata = float64(totalDurasi) / float64(daftarResep.JumlahResep)

	fmt.Println("\n===== STATISTIK =====")
	fmt.Println("Jumlah Resep        :", daftarResep.JumlahResep)
	fmt.Printf("Rata-rata Durasi    : %.2f menit\n", rataRata)

	fmt.Printf(
		"Resep Tercepat      : %s (%d menit)\n",
		daftarResep.Resep[idxTercepat].Judul,
		daftarResep.Resep[idxTercepat].Durasi,
	)

	fmt.Printf(
		"Resep Terlama       : %s (%d menit)\n",
		daftarResep.Resep[idxTerlama].Judul,
		daftarResep.Resep[idxTerlama].Durasi,
	)

	fmt.Printf(
		"Resep Paling Dicari : %s (%d kali)\n",
		daftarResep.Resep[idxPopuler].Judul,
		daftarResep.Resep[idxPopuler].JumlahDicari,
	)
}
// MENU STATISTIK END

// BUAT LUCU LUCUAN AJA PAN WALAUPUN BERANTAKAN HEHEHHEHEHE
func tampilanLuculucuan() {
	fmt.Println("==================================================")
	fmt.Println("         🍳 SELAMAT DATANG DI RESEPKU 🍳")
	fmt.Println("==================================================")
	fmt.Println(" Temukan, Kelola, dan Simpan Resep Favoritmu")
	fmt.Println("==================================================")
}
