# Time Complexity and Space Complexity

**Time Complexity**
Penggunaan kompleksitas waktu memudahkan estimasi waktu eksekusi program dengan melihat jumlah maksimum operasi primitif yang mungkin dilakukan. Operasi primitif termasuk penambahan, perkalian, penugasan, dan lainnya. Operasi yang paling sering dilakukan disebut operasi dominan, sementara beberapa operasi lainnya bisa diabaikan.

**Comparison of Different Time Complexities**
1. Constant time - O(1), jumlah operasi yang tetap.
2. Linear time - O(n), Jika nilai pertama dari larik A adalah 0, program akan segera berakhir. Namun, dalam menganalisis kompleksitas waktu, fokus pada kasus terburuk diperlukan. Program akan memakan waktu paling lama untuk dieksekusi ketika larik A tidak mengandung angka 0.
3. Linear time - O(n + m), Waktu linear O(n+m) mengindikasikan kompleksitas waktu yang tumbuh sejalan dengan total elemen dalam dua input yang berbeda, yaitu n dan m.
4. Logarithmic time - O(log n), Program ini memiliki loop di mana nilai n dibagi dua setiap iterasinya. Jika n awalnya adalah 2^x, maka jumlah iterasi yang diperlukan adalah x (di mana log n = x). Waktu eksekusi program tergantung pada nilai n yang diberikan sebagai masukan.
5. Quadratic time - O(n^2), Ketika menghitung kompleksitas, fokus pada istilah yang tumbuh paling cepat, jadi abaikan konstanta dan istilah lainnya. Hasilnya adalah kompleksitas waktu kuadrat. Terkadang, kompleksitas juga bergantung pada lebih banyak variabel.

 **Exponential and Factorial Time**
 Ada jenis kompleksitas waktu lain, seperti kompleksitas faktorial O(n!) dan kompleksitas eksponensial O(2^n). Dimanamhanya cocok untuk masalah dengan nilai n yang sangat kecil karena terlalu lambat jika digunakan untuk nilai n yang besar.

**Space Complexity**
Batasan memori memberikan info tentang kompleksitas ruang yang diharapkan. Perkirakan jumlah variabel yang boleh dideklarasikan. Singkatnya, jumlah variabel tetap berarti kompleksitas ruang tetap.










