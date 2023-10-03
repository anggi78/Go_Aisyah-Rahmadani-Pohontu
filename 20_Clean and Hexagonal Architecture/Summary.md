# Clean Architecture & Hexagonal Architecture
## Clean Architecture
Arsitektur yang baik adalah pemisahan perhatian menggunakan lapisan untuk membangun aplikasi yang modular, terukur, dapat dipelihara dan dapat diuji.
#### Benefit clean architecture
1. memiliki standar struktur
2. lebih cepat dalam development
3. mocking depedensi
4. easy switching

#### CA Layer
1. entities layer -> bisnis object
2. use case layer / domain layer/services -> bisnis logic
3. controller/presenter/handler -> user interaction, keberadaan framework
4. drivers-data layer - menerima data, mengambil, menyimpan data.

myapp/
├── entities/
│   └── user.go
├── usecases/
│   └── user_usecase.go
├── interfaces/
│   ├── controllers/
│   │   └── controller.go
│   └── repositories/
│       └── repository.go
├── main.go

## Contex Golang
Context merupakan sebuah data yang membawa value, sinyal cancel, sinyal timeout dan sinyal deadline. Context biasanya dibuat per request (misal setiap ada request masuk ke server web melalui http request). 

Context digunakan untuk mempermudah kita meneruskan value, dan sinyal antar proses
Membuat context kosong.
ctx :=  context.Background()

Parent dan child context
Context menganut konsep parent dan child Artinya, saat kita membuat context, kita bisa membuat child context dari context yang sudah ada
Parent context bisa memiliki banyak child, namun child hanya bisa memiliki satu parent context. fitur yang dimiliki oleh parent maka dimiliki juga oleh childnya
Context merupakan object yang Immutable, artinya setelah Context dibuat, dia tidak bisa diubah lagi.

Contex with value(menambahkan isi contex)
kita bisa menambahkan value ke sebuah root context. Saat menambah value ke context, secara otomatis akan tercipta child context baru, artinya original context nya tidak akan berubah sama sekali Untuk membuat menambahkan value ke context, kita bisa menggunakan function context.WithValue(parent, key, value)

 ctx :=  context.Background()
 ctxA = context.WithValue(ctx, key, value)
 ctxB = context.WithValue(ctxA, key, value)