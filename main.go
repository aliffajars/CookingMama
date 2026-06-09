package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX_RESEP = 100
const MAX_BAHAN = 100

type Resep struct {
	Judul        string            //done
	Kategori     string            // done
	BahanUtama   string            // done
	Bahan        [MAX_BAHAN]string // done
	JumlahBahan  int               //done
	Langkah      string            // done
	Durasi       int               // done
	JumlahDicari int               // done
}

type DataResep struct {
	Resep       [MAX_RESEP]Resep
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
	var step int = 1
	var inputLangkah, inputDurasi string
	var d int
	var langkahStr string

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         ➕ TAMBAH RESEP")
	fmt.Println("========================================")

	if daftarResep.JumlahResep >= MAX_RESEP {
		fmt.Println("Maaf, jumlah resep sudah mencapai batas maksimum.")
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

	fmt.Printf(
		"✅ %d bahan berhasil ditambahkan\n",
		r.JumlahBahan,
	)

	fmt.Println()

	// kalau gak "." program tetep jalan, jadi berhentiin make "."
	fmt.Println("Langkah Memasak (Ketik . pada langkah untuk selesai):")
	r.Durasi = 0

	// While loop di langkah masak nya dan durasi nya
	inputLangkah = inputString(fmt.Sprintf("Langkah %d: ", step))
	for inputLangkah != "." {

		inputDurasi = inputString(fmt.Sprintf("Durasi langkah %d (menit, enter/. jika lewati): ", step))

		var barisLangkah string
		if inputDurasi == "" || inputDurasi == "." {
			barisLangkah = fmt.Sprintf("%d. %s", step, inputLangkah)
		} else {
			d = 0
			fmt.Sscanf(inputDurasi, "%d", &d)
			r.Durasi += d
			barisLangkah = fmt.Sprintf("%d. %s (%d menit)", step, inputLangkah, d)
		}

		if step == 1 {
			langkahStr = barisLangkah
		} else {
			langkahStr += "\n" + barisLangkah
		}

		step++
		inputLangkah = inputString(fmt.Sprintf("Langkah %d: ", step))
	}
	r.Langkah = langkahStr

	daftarResep.Resep[daftarResep.JumlahResep] = r
	daftarResep.JumlahResep++

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
		fmt.Println("📝 Daftar Bahan")
		fmt.Println("----------------------------------------")

		for j = 0; j < daftarResep.Resep[i].JumlahBahan; j++ {
			fmt.Printf(
				"• %s\n",
				daftarResep.Resep[i].Bahan[j],
			)
		}

		fmt.Println()
		fmt.Println("👨‍🍳 Langkah Memasak")
		fmt.Println("----------------------------------------")
		fmt.Println(daftarResep.Resep[i].Langkah)

		fmt.Println()
		fmt.Printf(
			"⏱️ Total Durasi : %d menit\n",
			daftarResep.Resep[i].Durasi,
		)
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("📚 Total Resep :", daftarResep.JumlahResep)
	fmt.Println("========================================")
}

// TAMPILKAN JUDUL SAJA (Fungsi Bantuan)
func tampilDaftarJudul() {
	var i int
	if daftarResep.JumlahResep > 0 {

		fmt.Println("📋 Daftar Judul Resep Tersedia:")
		for i = 0; i < daftarResep.JumlahResep; i++ {
			fmt.Printf("   %d. %s\n", i+1, daftarResep.Resep[i].Judul)
		}
		fmt.Println("========================================")
	} else {
		fmt.Println("📭 Belum ada resep yang tersimpan.")
		fmt.Println("========================================")
	}
}

func editResep() {
	var judul string
	var idx, i int
	var temp string

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         ✏️  EDIT RESEP")
	fmt.Println("========================================")

	tampilDaftarJudul()

	if daftarResep.JumlahResep == 0 {
		return // Jika kosong, langsung kembali ke menu utama
	}

	judul = inputString("Masukkan judul resep yang ingin diedit: ")

	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("⚠️ Resep tidak ditemukan.")
		return
	}

	//penambahan fitur, edit nama resep, kategori, Pokonya semuanya dah, kalau mau skip make "." atau "enter" aja
	fmt.Println("\n===== MENGEDIT RESEP =====")
	fmt.Println("Petunjuk: Ketik . atau biarkan kosong (Enter) jika tidak ingin mengubah.")

	temp = inputString(fmt.Sprintf("Judul Baru [%s]: ", daftarResep.Resep[idx].Judul))
	if temp != "." && temp != "" {
		daftarResep.Resep[idx].Judul = temp
	}

	temp = inputString(fmt.Sprintf("Kategori Baru [%s]: ", daftarResep.Resep[idx].Kategori))
	if temp != "." && temp != "" {
		daftarResep.Resep[idx].Kategori = temp
	}

	temp = inputString(fmt.Sprintf("Bahan Utama Baru [%s]: ", daftarResep.Resep[idx].BahanUtama))
	if temp != "." && temp != "" {
		daftarResep.Resep[idx].BahanUtama = temp
	}

	fmt.Println()
	fmt.Println("📝 Edit Daftar Bahan")
	fmt.Println("(Kosongkan input lalu tekan Enter jika sudah selesai)")
	fmt.Println()

	daftarResep.Resep[idx].JumlahBahan = 0

	for i = 0; i < MAX_BAHAN; i++ {

		bahan := inputString(
			fmt.Sprintf("Bahan %d: ", i+1),
		)

		if bahan == "" {
			break
		}

		daftarResep.Resep[idx].Bahan[i] = bahan
		daftarResep.Resep[idx].JumlahBahan++
	}

	fmt.Println("\nLangkah Memasak Baru (Ketik . pada Langkah 1 jika tidak ingin mengubah, atau . untuk selesai):")
	var step int = 1
	var inputLangkah, inputDurasi string
	var d, durasiBaru int
	var langkahStr string

	//nah ini kondisi perulangan buat langkah resep nya
	inputLangkah = inputString(fmt.Sprintf("Langkah %d: ", step))
	for inputLangkah != "." && !(step == 1 && inputLangkah == "") {

		inputDurasi = inputString(fmt.Sprintf("Durasi langkah %d (menit, enter/. jika lewati): ", step))

		var barisLangkah string
		if inputDurasi == "" || inputDurasi == "." {
			barisLangkah = fmt.Sprintf("%d. %s", step, inputLangkah)
		} else {
			d = 0
			fmt.Sscanf(inputDurasi, "%d", &d)
			durasiBaru += d
			barisLangkah = fmt.Sprintf("%d. %s (%d menit)", step, inputLangkah, d)
		}

		if step == 1 {
			langkahStr = barisLangkah
		} else {
			langkahStr += "\n" + barisLangkah
		}

		step++
		inputLangkah = inputString(fmt.Sprintf("Langkah %d: ", step))
	}

	if step > 1 { // kalau ada langkah baru bakal nimpa langkah lama
		daftarResep.Resep[idx].Langkah = langkahStr
		daftarResep.Resep[idx].Durasi = durasiBaru
	}

	fmt.Println("✅ Resep berhasil diubah!")
}

func hapusResep() {
	var judul string
	var idx, i int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         🗑️  HAPUS RESEP")
	fmt.Println("========================================")

	tampilDaftarJudul()

	if daftarResep.JumlahResep == 0 {
		return // kalau kosong kembali ke menu utama
	}

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
	fmt.Println("✅ Resep berhasil dihapus!")
}

// END CRUD RESEP

// MENU SEARCHING START TANPA BIN DAN SEQ
func cariResepMenu() {
	var metode, urutan int
	var judul string
	var idx, j int

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         🔍 CARI RESEP")
	fmt.Println("========================================")

	fmt.Println("Metode Pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Println("0. Kembali")

	fmt.Print("👉 Pilih metode: ")
	fmt.Scan(&metode)

	reader.ReadString('\n')

	if metode == 0 {
		return
	}

	fmt.Println()
	fmt.Println("Urutan Judul:")
	fmt.Println("1. Ascending (A-Z)")
	fmt.Println("2. Descending (Z-A)")
	fmt.Println("0. Kembali")

	fmt.Print("👉 Pilih urutan: ")
	fmt.Scan(&urutan)

	reader.ReadString('\n')

	if urutan == 0 {
		return
	}

	if urutan == 1 {
		selectionSortJudul(true)
	} else if urutan == 2 {
		selectionSortJudul(false)
	} else {
		fmt.Println("❌ Pilihan tidak valid")
		return
	}

	fmt.Println()
	tampilDaftarJudul()

	judul = inputString(
		"Masukkan judul resep yang ingin dicari: ",
	)

	if metode == 1 {

		idx = cariResepSequential(judul)

	} else if metode == 2 {

		idx = cariResepBinary(judul)

	} else {

		fmt.Println("❌ Pilihan tidak valid")
		return
	}

	if idx == -1 {

		fmt.Println("⚠️ Resep tidak ditemukan.")
		return
	}

	daftarResep.Resep[idx].JumlahDicari++

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("         📖 RESEP DITEMUKAN")
	fmt.Println("========================================")

	fmt.Println("🍽️  Judul       :", daftarResep.Resep[idx].Judul)
	fmt.Println("📂  Kategori    :", daftarResep.Resep[idx].Kategori)
	fmt.Println("🥘  Bahan Utama :", daftarResep.Resep[idx].BahanUtama)

	fmt.Println()
	fmt.Println("📝 Daftar Bahan")
	fmt.Println("----------------------------------------")

	for j = 0; j < daftarResep.Resep[idx].JumlahBahan; j++ {
		fmt.Printf(
			"• %s\n",
			daftarResep.Resep[idx].Bahan[j],
		)
	}

	fmt.Println()
	fmt.Println("👨‍🍳 Langkah Memasak")
	fmt.Println("----------------------------------------")
	fmt.Println(daftarResep.Resep[idx].Langkah)

	fmt.Println()
	fmt.Printf(
		"⏱️ Total Durasi : %d menit\n",
		daftarResep.Resep[idx].Durasi,
	)
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