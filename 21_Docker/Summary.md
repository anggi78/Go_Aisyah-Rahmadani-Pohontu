# Docker
Docker adalah platform perangkat lunak yang digunakan untuk mengembangkan, mengemas, dan menjalankan aplikasi dalam wadah (container) yang ringan. Ini memungkinkan pengembang untuk mengisolasi aplikasi dan semua dependensinya dalam lingkungan yang konsisten, portabel, dan dapat diimplementasikan di berbagai platform, termasuk mesin fisik, mesin virtual, atau cloud.

**Docker basic**
1. Image
2. Container
3. Engine
4. Registry
5. Control Plane

## Container
Kontainer adalah unit standar yang dapat dijalankan yang berisi aplikasi dan semua dependensinya. Kontainer memungkinkan pengembang untuk mengemas aplikasi dan semua komponennya (pustaka, kode, variabel lingkungan, dll.) ke dalam lingkungan terisolasi yang konsisten dan portabel.

#### Perbedaan Container dan Virtual Machines
Container :

1. Abstraksi pada lapisan aplikasi
2. Kontainer memakan lebih sedikit ruang dibandingkan VM
3. Menangani lebih banyak aplikasi dan memerlukan lebih sedikit
4. VM dan Sistem Operasi

Virtual Machines :
1. Abstraksi perangkat keras fisik
2. Setiap VM menyertakan salinan lengkap Sistem operasi
3. Juga lambat untuk boot

#### Syntax Docker
- FROM -> Mendapatkan gambar dari registri buruh pelabuhan
- RUN -> Jalankan perintah bash saat membuat Kontainer
- ENV -> Tetapkan variabel di dalam Kontainer
- ADD -> Salin file dengan beberapa proses lainnya
- COPY -> Salin file
- WORKDIR -> Atur direktori file kerja
- ENTRYPOINT -> Jalankan perintah setelah selesai membangun Kontainer
- CMD -> Jalankan perintah tetapi dapat ditimpa

## Introduction Deployment
Deployment adalah kegiatan yang bertujuan untuk menyebarkan aplikasi/produk yang telah dikerjakan oleh para pengembang seringkali untuk mengubah dari status sementara ke permanen. 

**Strategi Deployment**
1. Big-Bang Deployment Strategy
2. Rollout Deployment Strategy
3. Blue/Green Deployment Strategy
4. Canary Deployment Strategy

### Docker Volume
Docker Volume bisa dianggap sebagai storage atau tempat penyimpanan data di container. Tentunya saat kita membuat container kita tidak ingin ketika container kita mati data yang ada pada container ikut terhapus juga. Untuk itu kita dapat memanfaatkan Volume pada docker.

### Docker Network
Defaultnya container pada docker akan saling terisolasi satu sama lain. Kita tidak dapat melakukan request api (misal) dari container satu ke container lain. Untuk itu kita harus membuat dan mendaftarkan container pada network yang sama.