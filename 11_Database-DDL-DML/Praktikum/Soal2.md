### Data Definition Language (DDL)

1. Create database alta_online_shop.
 ![Code Review](</Screenshot/Soal2_1.png> "Code Review")
2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
    - Create table user.
    ![Code Review](</Screenshot/Soal2_2.png> "Code Review")
    - Create table product, product type, operators, product description, payment_method.
    ![Code Review](</Screenshot/Soal2_3.png> "Code Review")
    ![Code Review](</Screenshot/Soal2_4.png> "Code Review")
    ![Code Review](</Screenshot/Soal2_5.png> "Code Review")
    ![Code Review](</Screenshot/Soal2_6.png> "Code Review")
    ![Code Review](</Screenshot/Soal2_7.png> "Code Review")
    - Create table transaction, transaction detail.
    ![Code Review](</Screenshot/Soal2_8.png> "Code Review")
    ![Code Review](</Screenshot/Soal2_9.png> "Code Review")
3. Create tabel kurir dengan field id, name, created_at, updated_at.
![Code Review](</Screenshot/Soal2_10.png> "Code Review")
4. Tambahkan ongkos_dasar column di tabel kurir.
![Code Review](</Screenshot/Soal2_11.png> "Code Review")
5. Rename tabel kurir menjadi shipping.
![Code Review](</Screenshot/Soal2_12.png> "Code Review")
6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
![Code Review](</Screenshot/Soal2_13.png> "Code Review")
7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    - 1-to-1: payment method description.
    ![Code Review](</Screenshot/Soal2_14.png> "Code Review")
    - 1-to-many: user dengan alamat.
    ![Code Review](</Screenshot/Soal2_15.png> "Code Review")
    - many-to-many: user dengan payment method menjadi user_payment_method_detail.
    ![Code Review](</Screenshot/Soal2_16.png> "Code Review")