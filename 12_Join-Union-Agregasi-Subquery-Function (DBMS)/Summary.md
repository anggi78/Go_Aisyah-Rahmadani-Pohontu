# Database - DDL - DML
Database adalah sekumpulan data yang terorganisir.

Database Relationship:
1. **One to One**, 1 user hanya memiliki 1 foto profil.
2. **One to Many**, 1 user bisa memiliki banyak tweets.
3. **Many to Many**, 1 user bisa memiliki banyak follower user, 1 user bisa di follow banyak user.

Software yang menggunakan Relational Database Model sebagai dasarnya, contoh: MySQL.

Jenis Perintah SQL:
- **DDL (Data Definition Language)**, Digunakan untuk membuat, mengubah, dan menghapus objek database seperti tabel, indeks, dan constraint. **Statement**: Create, Use, Drop, Rename, Alter.
- **DML (Data Manipulation Language)**, Perintah yang digunakan untuk memanipulasi data dalam tabel dari suatu database. **Statement Operation**: Insert, Select, Update, Delete. DML Statement: Like/Between, And/Or, Order By, Limit.
- **DCL (Data Control Language)**

Tipe Data MySQL:
- Num
- Huruf
- Date

## Join
Sebuah klausa untuk mengkombinasikan record dari dua atau lebih tabel.

Join Standar SQL ANSI:
1. **Inner**, akan mengembalikan baris-baris dari dua tabel atau lebih yang memnuhi syarat.
2. **Left**, akan mengembalikan seluruh baris dari tabel disebelah kiri yang dikenal kondisi ON dan hanya baris dari tabel disebelah kanan yang memenuhi kondisi join.
3. **Right**, akan mengembalikan semua baris dari tabel sebelah kanan yang dikenai kondisi ON dengan data dari tabel sebelah kiri yangn memenuhi kondisi join. Teknik ini merupakan kebalikan dari left join.

Union, ada hal yang harus diperhatikan dari union adalah jumlah field yang dikeluarkan/dipanggil harus sama.

## Aggregate 
Fungsi agregasi adalah fungsi dimana nilai beberaap baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.

Fungsi Agregasi SQL:
1. **Min**, fungsi dimana nilai beberapa baris dikelompokkan bersama untuk membentuk nilai ringkasan tunggal.
2. **Max**, Digunakan untuk mendapatkan nilai maximum atau nilai terbesar dari sebuah data record ditabel.
3. **Sum**, Digunakan untuk mnedapatkan jumlah total nilai dari sebuah data atau record ditabel.
4. **Avg**, digunakan untuk mencari nilai rata-rata dari sebuah data atau record ditabel.
5. **Count**, digunakan untuk mencari jumlah dari sebuah data atau record ditabel.
6. **Having**, digunakan untuk menyeleksi data berdasarkan kriteria tertentu, dimana kriteria berupa fungsi aggregat.

## SubQuery
Subquery atau Inner query atau Nested query adalah query di dalam query SQL lain.

Sebuah subquery  digunakan untuk mengembalikan data yang akan digunakan dalam query utama sebagai syarat untuk lebih membatasi data yang akan diambil.

## Function
Sebuah kumpulan statetment yang akan mengembalikan sebuah nilai balik pada pemanggilnya.