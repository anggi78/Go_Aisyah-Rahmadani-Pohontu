# System Design

## Diagram
Yaitu representasi simbolis dari informasi yang menggunakan teknik visualisasi Tools design : smartdraw, Lucidchart, draw.io, Whimsical, Visio. 
Desain perangkat lunak yang umum digunakan: 
1. Flowchart, teorinya mewakili satu proses, terdapat process, decision, terminator.
2. Use Case, merangkum rincian penggunaan system (dikenal sebagai aktor) dan interaksi mereka dengan sistem.
3. ERD, jenis diagram alur yang menggambarkan bagaimana "entitas" seperti orang, objek, atau konsep berhubungan satu sama lain dengan sistem.
4. HLA, desain sistem secara keseluruhan, memeriksa audiens target. 

Karakteristik utama sistem terdistribusi
- Scalability (kemampuan)
- Reliability (keandalan)
- Availability (ketersediaan)
- Efficiency (efisiensi)
- Serviceability or Manageability (kemudahan servis atau pengelolaan)

#### Job/Work Queue
Dalam perangkat lunak sistem, job queue (kadang disebut antrian batch) adalah struktur data yang dikelola oleh perangkat lunak penjadwalan pekerjaan yang berisi pekerjaan yang akan dijalankan.

#### Load Balancing
Load Balancer (LB) adalah komponen kritis lainnya dari setiap sistem terdistribusi. Ini membantu menyebar lalu lintas di antara sekelompok server untuk meningkatkan responsivitas dan ketersediaan aplikasi, situs web, atau basis data.

Untuk memanfaatkan skalabilitas penuh dan redundansi, kita dapat mencoba menyeimbangkan beban di setiap lapisan sistem. Kita dapat menambahkan Load Balancer (LB) pada tiga tempat berikut:

1. Antara pengguna dan server web.
2. Antara server web dan lapisan platform internal, seperti server aplikasi atau server cache.
3. Antara lapisan platform internal dan basis data.

#### Monolithic and Microservices
Aplikasi monolithic memiliki satu basis kode dengan beberapa modul. Modul-modul tersebut dibagi menjadi modul untuk fitur bisnis atau fitur teknis. Aplikasi ini memiliki satu sistem pembangunan yang membangun seluruh aplikasi dan/atau dependensinya. Ini juga memiliki satu binary yang dapat dieksekusi atau didistribusikan.

Microservices adalah layanan yang dapat dideploy secara independen yang didasarkan pada domain bisnis tertentu. Mereka berkomunikasi satu sama lain melalui jaringan, dan sebagai pilihan arsitektur, mereka menawarkan banyak opsi untuk menyelesaikan masalah yang mungkin Anda hadapi. Oleh karena itu, arsitektur microservices didasarkan pada beberapa microservices yang bekerja sama.

#### SQL dan NoSQL

SQL dan NoSQL (atau basis data relasional dan basis data non-relasional).

- Basis data relasional adalah terstruktur dan memiliki skema yang telah ditentukan sebelumnya, seperti buku telepon yang menyimpan nomor telepon dan alamat.
- Basis data non-relasional adalah tidak terstruktur dan memiliki skema dinamis, seperti folder berkas yang dapat berisi segala hal mulai dari alamat dan nomor telepon seseorang hingga 'like' Facebook dan preferensi berbelanja online mereka.

#### Caching
Sebuah cache mirip dengan ingatan jangka pendek: ia memiliki jumlah ruang yang terbatas, tetapi biasanya lebih cepat daripada sumber data asli dan berisi item yang paling baru diakses. Cache dapat ada di semua tingkat dalam arsitektur, tetapi seringkali ditemukan di tingkat yang paling dekat dengan bagian depan di mana mereka diimplementasikan untuk mengembalikan data dengan cepat tanpa membebani tingkat lebih bawah.

#### Database Replication
##### Redundancy dan Replication
Redundansi adalah duplikasi komponen atau fungsi kritis dari suatu sistem dengan tujuan meningkatkan keandalan sistem, biasanya dalam bentuk cadangan atau mekanisme keamanan, atau untuk meningkatkan kinerja sistem sebenarnya.

#### Database Indexing
Pengindeksan adalah cara untuk mengoptimalkan kinerja basis data dengan meminimalkan jumlah akses ke disk yang diperlukan saat sebuah query dijalankan. Ini adalah teknik struktur data yang digunakan untuk dengan cepat menemukan dan mengakses data dalam sebuah basis data.