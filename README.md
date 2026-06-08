**Cooking Mama** adalah aplikasi *Command Line Interface* (CLI) interaktif berbasis **Golang** untuk mencatat dan mengelola resep masakan. Aplikasi ini sangat cocok untuk menyimpan resep rahasia Anda dan dilengkapi dengan fitur pencarian, pengurutan, manipulasi data pintar, hingga statistik memasak.
---
## ✨ Fitur Utama
*   **➕ Tambah Resep Dinamis**
    *   Mendukung input langkah memasak yang tak terbatas (menggunakan *while loop*). Cukup ketik `.` untuk berhenti.
    *   Durasi tiap langkah akan diakumulasikan secara otomatis menjadi total waktu memasak.
*   **📖 Tampilkan Resep**
    *   Menampilkan data resep secara detail, rapi, dan terstruktur.
*   **🔍 Cari Resep (Algoritma Cerdas)**
    *   Dapat melihat daftar menu yang tersedia sebelum mulai mencari.
    *   Mendukung algoritma **Sequential Search** dan **Binary Search**.
    *   Pencarian bersifat *Case-Insensitive* (tidak terpengaruh huruf besar/kecil).
*   **✏️ Edit Resep Pintar (Smart Update)**
    *   Anda tidak perlu mengetik ulang semua data! Cukup tekan `Enter` atau ketik `.` pada bagian data (Judul, Bahan, Langkah) yang **tidak ingin diubah**.
*   **🗑️ Hapus Resep**
    *   Menghapus resep secara permanen dari daftar memori.
*   **↕️ Urutkan Resep (Sorting Algorithms)**
    *   **Selection Sort**: Mengurutkan resep berdasarkan Judul A-Z atau Z-A.
    *   **Insertion Sort**: Mengurutkan resep berdasarkan Total Durasi tercepat atau terlama.
*   **📊 Statistik Resep**
    *   Sistem melacak otomatis: Total Resep, Rata-rata Durasi Memasak, Resep Tercepat, Resep Terlama, hingga **Resep yang Paling Sering Dicari/Populer**.
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