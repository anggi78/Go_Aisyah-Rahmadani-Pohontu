# Concurrent Programing
Yaitu menjalankan tugas secara bersamaan dan secara independent (berdiri sendiri) 

Feature Go :
1. Concurrent execution (Goroutines)
2. Synchronization and messaging (Channels)
3. Multi-way concurrent control (Select)
## Goroutines
Sebuah proses concurrent yang dijalankan di dalam golang. Goroutines adalah sebuah function yang ditentukan akan berjalan secara concurrently dibandingkan dengan function yang lain.

Gomaxprocs : Ini digunakan untuk mengontrol jumlah thread sistem operasi yang tersedia untuk eksekusi goroutine dalam program Go tertentu.

## Channel & Select
Sebuah channel adalah objek komunikasi yang digunakan oleh goroutine untuk berkomunikasi satu sama lain.

Channel dalam Go dapat dibagi menjadi dua jenis:
1. **Buffered channel**,  terdapat kapasitas tertentu untuk menyimpan pesan sebelum penerimaannya oleh goroutine penerima. Ketika buffer penuh, pengirim pesan tidak perlu menunggu penerima. Mereka dapat melanjutkan eksekusi mereka. Hanya ketika buffer penuh, pengiriman pesan akan memblokir pengirim sampai ada ruang di buffer untuk pesan baru. **Contoh:** "ch := make(chan int, 10)" akan membuat buffered channel dengan kapasitas buffer 10.
2. **Unbuffered channel**,  setiap pengiriman pesan dari satu goroutine ke goroutine lainnya akan mengharuskan pengirim menunggu hingga penerima siap untuk menerima pesan tersebut. Ini memastikan bahwa pesan yang dikirim akan segera diterima oleh penerima sebelum goroutine pengirim melanjutkan eksekusi selanjutnya. **Contoh:** "ch := make(chan int)".

**Select** adalah pernyataan yang digunakan dalam Go untuk mengendalikan komunikasi data melalui satu atau banyak channel.

## Race Condition
Race condition (kondisi balapan) terjadi ketika dua thread mengakses memori pada waktu yang bersamaan, di mana salah satu dari mereka sedang menulis. Race condition terjadi karena akses yang tidak disinkronkan ke memori bersama.

Fixing Data Race :
1. WaitGroups : Cara yang paling langsung untuk mengatasi data race adalah dengan memblokir akses baca hingga operasi penulisan telah selesai.
2. Channels Blocking : Ini memungkinkan goroutine-goroutine kita untuk melakukan sinkronisasi satu sama lain untuk sesaat.
3. Mutex : Jika kita tidak peduli dengan urutan baca dan tulis, yang penting adalah bahwa mereka tidak terjadi secara bersamaan, maka kita sebaiknya mempertimbangkan penggunaan mutex.