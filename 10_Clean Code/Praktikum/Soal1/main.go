package main

// Struktur user memiliki beberapa kekurangan:
// 1. Variabel "username" dan "password" seharusnya bertipe string, bukan int.
// 2. Nama variabel yang kurang deskriptif.
type user struct {
	id       int
	username int // Alasan: Tipe data seharusnya string.
	password int // Alasan: Tipe data seharusnya string.
}

// Struktur userservice juga memiliki kekurangan:
// 1. Variabel "t" seharusnya memiliki nama yang lebih deskriptif.
type userservice struct {
	t []user // Alasan: Nama variabel kurang deskriptif.
}

// Metode "getallusers" dan "getuserbyid" menggunakan receiver nilai, bukan pointer receiver.
// Ini berarti metode ini tidak dapat memodifikasi nilai "userservice" yang sesungguhnya.
// Alasan: Untuk memungkinkan modifikasi pada "userservice", gunakan pointer receiver (*userservice).
func (u userservice) getallusers() []user {
	return u.t
}

// Metode "getuserbyid" mengembalikan "user{}" ketika pengguna tidak ditemukan.
// Sebaiknya, menggunakan "nil" atau tipe data yang sesuai seperti "*user" untuk menunjukkan ketidakberhasilan dalam pencarian.
// Alasan: Ini akan membedakan antara pengguna yang ditemukan dengan ID yang cocok dan ketidakberhasilan dalam pencarian.
func (u userservice) getuserbyid(id int) user {
	for _, r := range u.t {
		if id == r.id {
			return r
		}
	}

	return user{}
}
