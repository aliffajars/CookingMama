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

func initData() {
	daftarResep.Resep[0] = Resep{
		Judul:       "Nasi Goreng",
		Kategori:    "Makanan Berat",
		BahanUtama:  "Nasi",
		JumlahBahan: 3,
		Durasi:      15,
		Langkah: "1. Tumis bawang (5 menit)\n" +
			"2. Masukkan nasi (5 menit)\n" +
			"3. Tambahkan kecap (5 menit)",
	}

	daftarResep.Resep[0].Bahan[0] = "Nasi"
	daftarResep.Resep[0].Bahan[1] = "Kecap"
	daftarResep.Resep[0].Bahan[2] = "Bawang"

	daftarResep.Resep[1] = Resep{
		Judul:       "Mie Goreng",
		Kategori:    "Makanan Berat",
		BahanUtama:  "Mie",
		JumlahBahan: 3,
		Durasi:      10,
		Langkah: "1. Rebus mie (5 menit)\n" +
			"2. Tiriskan mie (2 menit)\n" +
			"3. Campur bumbu (3 menit)",
	}

	daftarResep.Resep[1].Bahan[0] = "Mie"
	daftarResep.Resep[1].Bahan[1] = "Bumbu"
	daftarResep.Resep[1].Bahan[2] = "Minyak"

	daftarResep.JumlahResep = 2

		daftarResep.Resep[2] = Resep{
		Judul:       "Rendang",
		Kategori:    "Makanan Berat",
		BahanUtama:  "Daging Sapi",
		JumlahBahan: 4,
		Durasi:      120,
		Langkah: "1. Tumis bumbu halus (20 menit)\n" +
			"2. Masukkan daging sapi (10 menit)\n" +
			"3. Tambahkan santan (10 menit)\n" +
			"4. Masak hingga kering (80 menit)",
	}

	daftarResep.Resep[2].Bahan[0] = "Daging Sapi"
	daftarResep.Resep[2].Bahan[1] = "Santan"
	daftarResep.Resep[2].Bahan[2] = "Cabai"
	daftarResep.Resep[2].Bahan[3] = "Bawang"

	daftarResep.Resep[3] = Resep{
		Judul:       "Sate Ayam",
		Kategori:    "Makanan Berat",
		BahanUtama:  "Ayam",
		JumlahBahan: 4,
		Durasi:      45,
		Langkah: "1. Potong ayam (10 menit)\n" +
			"2. Tusuk ayam dengan tusukan sate (10 menit)\n" +
			"3. Olesi bumbu kacang (5 menit)\n" +
			"4. Bakar sate hingga matang (20 menit)",
	}

	daftarResep.Resep[3].Bahan[0] = "Ayam"
	daftarResep.Resep[3].Bahan[1] = "Bumbu Kacang"
	daftarResep.Resep[3].Bahan[2] = "Kecap"
	daftarResep.Resep[3].Bahan[3] = "Tusuk Sate"

	daftarResep.JumlahResep = 4
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

	initData()

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
			urutkanResep()
		case 7:
			statistik()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi resep masakan!")
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
	fmt.Println("===== MENU APLIKASI RESEP MASAKAN =====")
	fmt.Println("1. Tambah Resep")
	fmt.Println("2. Tampilkan Resep")
	fmt.Println("3. Cari Resep")
	fmt.Println("4. Edit Resep")
	fmt.Println("5. Hapus Resep")
	fmt.Println("6. Urutkan Resep")
	fmt.Println("0. Keluar")
	fmt.Println("7. Statistik")
	fmt.Print("Pilihan menu: ")
	fmt.Scan(&pilihan,)
	if pilihan < 0 || pilihan > 7 {
		fmt.Println("Pilihan tidak valid, silakan pilih menu yang tersedia.")
		return menu()
	}
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
	fmt.Println("===== TAMBAH RESEP =====")
	fmt.Println()

	r.Judul = inputString("Judul Resep: ")
	r.Kategori = inputString("Kategori: ")
	r.BahanUtama = inputString("Bahan Utama: ")
	fmt.Println()
	fmt.Println("===== Tambahkan Bahan =====")
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
		"%d bahan berhasil ditambahkan\n",
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

	fmt.Println("Resep berhasil ditambahkan!")
}

func tampilResep() {
	var i int

	fmt.Println()
	fmt.Println("===== DAFTAR RESEP =====")

	if daftarResep.JumlahResep == 0 {
		fmt.Println("Belum ada resep yang tersimpan.")
		return
	}

	fmt.Println()
	for i = 0; i < daftarResep.JumlahResep; i++ {
		fmt.Printf("%d. %s\n", i+1, daftarResep.Resep[i].Judul)
	}
}

func editResep() {
	var judul string
	var idx, i int
	var temp string

	fmt.Println()
	fmt.Println("===== EDIT RESEP =====")

	tampilDaftarJudul()

	if daftarResep.JumlahResep == 0 {
		return // Jika kosong, langsung kembali ke menu utama
	}

	judul = inputString("Masukkan judul resep yang ingin diedit: ")

	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println(" Resep tidak ditemukan.")
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
	fmt.Println("Edit Daftar Bahan")
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

	fmt.Println("Resep berhasil diubah!")
}

func hapusResep() {
	var judul string
	var idx, i int

	fmt.Println()
	fmt.Println("===== HAPUS RESEP =====")

	tampilDaftarJudul()

	if daftarResep.JumlahResep == 0 {
		return // kalau kosong kembali ke menu utama
	}

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

// END CRUD RESEP

// MENU SEARCHING START TANPA BIN DAN SEQ
func cariResep() {
	var judul string
	var idx, j int

	fmt.Println()
	fmt.Println("===== CARI RESEP =====")

	tampilDaftarJudul()

	judul = inputString("\nMasukkan judul resep: ")

	idx = cariResepSequential(judul)

	if idx == -1 {
		fmt.Println("Resep tidak ditemukan.")
		return
	}

	daftarResep.Resep[idx].JumlahDicari++

	fmt.Println()
	fmt.Println("===== DETAIL RESEP =====")
	fmt.Println("Judul       :", daftarResep.Resep[idx].Judul)
	fmt.Println("Kategori    :", daftarResep.Resep[idx].Kategori)
	fmt.Println("Bahan Utama :", daftarResep.Resep[idx].BahanUtama)

	fmt.Println("\nDaftar Bahan:")
	for j = 0; j < daftarResep.Resep[idx].JumlahBahan; j++ {
		fmt.Printf("- %s\n", daftarResep.Resep[idx].Bahan[j])
	}

	fmt.Println("\nLangkah Memasak:")
	fmt.Println(daftarResep.Resep[idx].Langkah)

	fmt.Printf("\nTotal Durasi: %d menit\n",
		daftarResep.Resep[idx].Durasi)
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
	var i int
	var idxPopuler int
	var kategori [MAX_RESEP]string
	var jumlahKategori [MAX_RESEP]int
	var jumlahDataKategori int
	var ditemukan bool
	var j int

	fmt.Println()
	fmt.Println("===== STATISTIK RESEP =====")

	if daftarResep.JumlahResep == 0 {
		fmt.Println("Belum ada resep yang tersimpan.")
		return
	}

	for i = 0; i < daftarResep.JumlahResep; i++ {

		ditemukan = false

		for j = 0; j < jumlahDataKategori; j++ {
			if kategori[j] == daftarResep.Resep[i].Kategori {
				jumlahKategori[j]++
				ditemukan = true
				break
			}
		}

		if !ditemukan {
			kategori[jumlahDataKategori] = daftarResep.Resep[i].Kategori
			jumlahKategori[jumlahDataKategori] = 1
			jumlahDataKategori++
		}

		if daftarResep.Resep[i].JumlahDicari >
			daftarResep.Resep[idxPopuler].JumlahDicari {
			idxPopuler = i
		}
	}

	fmt.Println()
	fmt.Println("Jumlah Resep per Kategori:")

	for i = 0; i < jumlahDataKategori; i++ {
		fmt.Printf(
			"- %s : %d resep\n",
			kategori[i],
			jumlahKategori[i],
		)
	}

	fmt.Println()

	if daftarResep.Resep[idxPopuler].JumlahDicari > 0 {
		fmt.Printf(
			"Resep Paling Dicari : %s (%d kali)\n",
			daftarResep.Resep[idxPopuler].Judul,
			daftarResep.Resep[idxPopuler].JumlahDicari,
		)
	} else {
		fmt.Println("Resep Paling Dicari : Belum ada pencarian")
	}
}

// MENU STATISTIK END

// TAMPILKAN JUDUL SAJA (Fungsi Bantuan)
func tampilDaftarJudul() {
	var i int
	if daftarResep.JumlahResep > 0 {

		fmt.Println("Daftar Judul Resep Tersedia:")
		for i = 0; i < daftarResep.JumlahResep; i++ {
			fmt.Printf("   %d. %s\n", i+1, daftarResep.Resep[i].Judul)
		}
	} else {
		fmt.Println("Belum ada resep yang tersimpan.")
	}
}