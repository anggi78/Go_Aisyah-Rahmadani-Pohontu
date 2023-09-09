# Clean Code
Clean code adalah istilah untuk kode yang mudah dibaca, dipahami dan diubah  oleh programmer. 

Karakteristik:
1. Mudah dipahami
2. Mudah dieja dan dicari
3. Singkat namun mendeskripsikan konteks
4. Konsisten
5. Hindari penambahan konteks yang tidak perlu
6. Komentar
7. Good function
8. Gunakan konvensi
9. Formatting

Principle:
1. KISS (Keep It So Simple), Hindari membuat fungsi yang dibuat untuk melakukan A, sekalian memodifikasi B, mengecek fungsi C, dst.
2. DRY (Don't Repeat Yourself), Duplikasi code terjadii karena sering copy paste. Untuk menghindari duplikasi code buatlah fungsi yang dapat digunakan secara berulang-ulang.

Refactoring, adalah proses restrukturisasi kode yang dibuat, dengan cara mengubah struktur internal tanpa mengubah perilaku eksternal.Prinsip KISS dan DRY bisa dicapai dengan cara refactor.

Teknik refactoring:

- Membuat sebuah abstraksi
- Memcah kode dengan fungsi/class
- Perbaiki penamaan dan lokasi kode
- Deteksi kode yang memiliki duplikasi