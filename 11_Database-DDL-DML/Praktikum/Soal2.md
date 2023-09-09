### Data Definition Language (DDL)

1. Create database alta_online_shop.
![alt text](/Screenshot/Soal2_1.png)
2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
    - Create table user.
    ![alt text](/Screenshot/Soal2_2.png)
    - Create table product, product type, operators, product description, payment_method.
    ![alt text](/Screenshot/Soal2_3.png)
    ![alt text](/Screenshot/Soal2_4.png)
    ![alt text](/Screenshot/Soal2_5.png)
    ![alt text](/Screenshot/Soal2_6.png)
    ![alt text](/Screenshot/Soal2_7.png)
    - Create table transaction, transaction detail.
    ![alt text](/Screenshot/Soal2_8.png)
    ![alt text](/Screenshot/Soal2_9.png)
3. Create tabel kurir dengan field id, name, created_at, updated_at.
![alt text](/Screenshot/Soal2_10.png)
4. Tambahkan ongkos_dasar column di tabel kurir.
![alt text](/Screenshot/Soal2_11.png)
5. Rename tabel kurir menjadi shipping.
![alt text](/Screenshot/Soal2_12.png)
6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
![alt text](/Screenshot/Soal2_13.png)
7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    - 1-to-1: payment method description.
    ![alt text](/Screenshot/Soal2_14.png)
    - 1-to-many: user dengan alamat.
    ![alt text](/Screenshot/Soal2_15.png)
    - many-to-many: user dengan payment method menjadi user_payment_method_detail.
    ![alt text](/Screenshot/Soal2_16.png)