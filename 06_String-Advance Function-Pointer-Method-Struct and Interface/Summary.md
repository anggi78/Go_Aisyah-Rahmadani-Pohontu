# String, Advance Function , Pointer , Method, Struct and Interface

## String
**String** adalah tipe data yang digunakan untuk merepresentasikan urutan karakter dalam pemrograman.

**Menggunakan String**
1. *len:* Mengukur panjang (jumlah karakter) dari sebuah string.
2. *compare:* Membandingkan dua string untuk menentukan urutan leksikal mereka.
3. *contains:* Memeriksa apakah sebuah string mengandung substring tertentu.
4. *substrings:* Membuat substring baru dari string yang ada.
5. *replace:* Mengganti kemunculan substring tertentu dengan substring lainnya dalam string.
6. *insert:* Menyisipkan sebuah substring ke dalam posisi tertentu dalam string.

## Function
**Variadic Function:**
- Untuk menghindari membuat slice sementara hanya untuk dilewatkan ke sebuah fungsi.
- Ketika jumlah parameter input tidak diketahui.
- Untuk menyatakan niat Anda agar kode lebih mudah dibaca.

**Closure** adalah jenis khusus dari fungsi tanpa nama yang merujuk pada variabel yang dideklarasikan di luar fungsi itu sendiri. Dalam kasus ini, kita akan menggunakan variabel yang tidak diteruskan ke dalam fungsi sebagai parameter, tetapi sebaliknya variabel tersebut tersedia ketika fungsi dideklarasikan.

**Fungsi defer**, sebuah fungsi hanya dieksekusi setelah fungsi induknya selesai. Penggunaan return ganda juga dapat digunakan, mereka dijalankan sebagai tumpukan, satu per satu.

## Pointer
**Pointer** adalah variabel yang menyimpan alamat memori dari variabel lain. Pointer memiliki kemampuan untuk mengubah data yang mereka tunjuk.

1. Pointer pada Tipe Data Biasa: Merujuk pada alamat memori variabel dari tipe data yang sama. Contohnya, int pointer merujuk pada alamat memori variabel int, string pointer merujuk pada alamat memori variabel string, dan seterusnya.
2. Pointer Struct: Anda dapat menggunakan pointer pada tipe data struct untuk merujuk pada instansi struktur.
3. Pointer ke Pointer: Anda bisa memiliki pointer yang merujuk pada pointer lain. Ini berguna ketika Anda ingin membuat perubahan pada variabel melalui beberapa tingkat pointer.
4. Pointer ke Fungsi: Anda bisa memiliki pointer yang merujuk pada fungsi. Ini memungkinkan Anda untuk meneruskan fungsi sebagai argumen ke fungsi lain.

## Method
**Method** adalah sebuah fungsi yang dilampirkan pada suatu tipe data (dapat berupa struktur atau tipe data lainnya).

Deklarasi Method mirip dengan fungsi, hanya saja deklarasi variabel objek perlu ditambahkan di antara kata kunci func dan nama fungsi.

Mengapa Menggunakan Method Daripada Fungsi?
1. Membantu Anda Menulis Kode Berorientasi Objek di Go.
2. Method Membantu Menghindari Konflik Nama.
3. Pemanggilan Method Lebih Mudah Dibaca dan Dimengerti Daripada Pemanggilan Fungsi.

## Interface
**Interface** adalah kumpulan dari tanda tangan method yang dapat diimplementasikan oleh suatu objek. Oleh karena itu, interface mendefinisikan perilaku dari objek tersebut.

**Package** adalah kumpulan dari fungsi-fungsi dan data.

Ketika runtime Go mendeteksi kesalahan-kesalahan ini, itu memicu **panic**. Untuk menambahkan kemampuan **recorver** dari error panic, Anda dapat menambahkan sebuah fungsi tanpa nama (anonymous function) atau mendefinisikan sebuah fungsi kustom dan memanggilnya dengan kata kunci "defer" dari dalam method, di mana panic mungkin terjadi dari pemanggilan internal lainnya.