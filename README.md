# 🍳 Cooking Mama
**Cooking Mama** adalah aplikasi *Command Line Interface* (CLI) interaktif berbasis **Golang** untuk mencatat dan mengelola resep masakan. Aplikasi ini sangat cocok untuk menyimpan resep rahasia Anda dan dilengkapi dengan fitur pencarian, pengurutan, manipulasi data pintar, hingga statistik memasak.
---
## ✨ Fitur Utama
*   **➕ Tambah Resep Dinamis**
    *   Mendukung input langkah memasak yang tak terbatas (menggunakan *while loop*). Cukup ketik `.` untuk berhenti.
    *   Durasi tiap langkah akan diakumulasikan secara otomatis menjadi total waktu memasak.
*   **📖 Tampilkan Resep**
    *   Menampilkan seluruh daftar judul resep yang tersimpan secara terstruktur.
*   **🔍 Cari Resep (Algoritma Cerdas)**
    *   Daftar judul resep ditampilkan otomatis sebelum pencarian dimulai.
    *   Terdapat dua pilihan metode pencarian:
        *   **Binary Search** → Cari berdasarkan **Judul Resep**. Data di-sort A-Z secara otomatis di balik layar sebelum pencarian dijalankan.
        *   **Sequential Search** → Cari berdasarkan **Bahan Utama**. Penelusuran dilakukan ke seluruh data sehingga dapat menampilkan **semua resep** yang menggunakan bahan utama yang sama (misal: semua resep berbahan utama "Ayam").
    *   Pencarian bersifat *Case-Insensitive* (tidak terpengaruh huruf besar/kecil).
*   **✏️ Edit Resep Pintar (Smart Update)**
    *   Anda tidak perlu mengetik ulang semua data! Cukup tekan `Enter` atau ketik `.` pada bagian data (Judul, Kategori, Bahan Utama, Bahan, Langkah) yang **tidak ingin diubah**.
*   **🗑️ Hapus Resep**
    *   Menghapus resep secara permanen dari daftar memori.
*   **↕️ Urutkan Resep (Sorting Algorithms)**
    *   **Selection Sort**: Mengurutkan resep berdasarkan Judul A-Z atau Z-A.
    *   **Insertion Sort**: Mengurutkan resep berdasarkan Total Durasi tercepat atau terlama.
*   **📊 Statistik Resep**
    *   Menampilkan **jumlah resep per kategori** secara otomatis.
    *   Menampilkan **resep yang paling sering dicari** berdasarkan riwayat pencarian pengguna.
---
## 🚀 Cara Menjalankan Program
Pastikan Anda sudah menginstal [Golang](https://go.dev/dl/) di komputer atau laptop Anda.
1. **Clone repositori ini**:
   ```bash
   git clone https://github.com/aliffajars/CookingMama.git
   cd CookingMama

2. **Menjalankan Program**:
    ```bash
    go run main.go