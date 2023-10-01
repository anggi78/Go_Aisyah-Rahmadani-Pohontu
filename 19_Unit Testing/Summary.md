# Unit Testing 
## Survey
Apa itu software testing? Yaitu suatu proses menganalisis item perangkat lunak untuk mendeteksi perbedaan antara kondisi yang ada dan yang diperlukan (yaitu cacat) dan untuk mengevaluasi fitur perangkat lunak.

Tujuan testing:
- Mencegah regresi
- Tingkat keyakinan dalam refactoring
- Meningkatkan desain kode
- Dokumentasi kode
- Kode penskalaan tim

Tingkat pengujian:
1. UI, (End To End) menguji interaksi antar keseluruhan melalui antarmuka pengguna.
2. Integration, menguji modul atau subsistem tertentu melalui api.
3. Unit, menguji bagian terkecil yang dapat diuji(operasi logis tunggal) dari suatu aplikasi melalui metode.

### Framework
Framework menyediakan alat dan struktur yang diperlukan untuk menjalankan pengujian secara efisien.

### Structure
2 pola biasa:
1. Pusatkan file pengujian di dalam folder pengujian.
2. Simpan file pengujian bersama dengan file produksi.

### Runner
- alat yang menjalankan file uji
- gunakan mode jam tangan (jika ada perubahan dalam basis kode, secara otomatis jalankan tes lagi, buat tdd lebih efisien)
- pilih pelari yang bisa berlari paling cepat

### Mocking
Kasing uji anda harus mandiri, anda tidak harus menguji:
- modul pihak ke-3
- basis data
- api pihak ke-3 atau sistem eksternal lainnya
- kelas objek yang telah anda uji di tempat lain

### Coverage
Pembuat kode perlu memastikan apakah mereka telah membuat cukup pengujian.
Alat cakupan menganalisis kode aplikasi saat pengujian sedang berjalan.