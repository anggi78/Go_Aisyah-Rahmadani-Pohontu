# RESTful API
API adalah sekumpulan fungsi dan prosedur yang memungkinkan pembuatan aplikasi yang mengakses fitur atau data sistem operasi, aplikasi atau layanan lainnya.

Request & Response format :
- JSON
- XML
- SOAP

HTTP request Method :
- GET
- POST
- PUT
- DELETE
- HEAD
- OPTION
- PATCH

**Postman** klien HTTP untuk menguji layanan web. postman memudahkan pengujian, pengembangan dan dokumentasi API dengan memungkinkan pengguna dengan cepat menyusun permintaan HTTP yang sederhana dan kompleks.

# Intro restful API
- introduction Api
- introduction to postman
- REST
- JSON
- HTTP respon code
- package net/http

api bekerja dengan client(browser) mengirimkan request ke server dan kemudians server akan memberikan response yang sesuai dengan requestnya.
 
Implementasi dan kegunaan rest API:
- mengintegrasikan aplikasi yang berbasis front end (web. mobile, dll) dan backend. front end ada banyak dan be hanya 1 tetap bisa digunakan
- mengintegrasikan BE ke BE
misalnya  service sepulsa penyedia layanan pembelian pulsa, sepulsa dapat menggunakan fitur pembayaran, seperti yang disediakan oleh midtrans(service yang melayani pembayaran)
 
### open API
open API adalah API yang dapat digunakan dengan bebas.
contoh-contoh open Api: https://github.com/public-apis/public-apis

1. OpenWeatherMap API: [OpenWeatherMap API](https://openweathermap.org/api)

2. Google Maps API: [Google Maps API](https://developers.google.com/maps)

3. Twitter API: [Twitter Developer](https://developer.twitter.com/en)

4. GitHub API: [GitHub Developer](https://developer.github.com/)

5. YouTube API: [YouTube API](https://developers.google.com/youtube/)

### Best Practice RestAPI
1. use Nouns instead of verbs

    | Baik (Nouns)    | Kurang Baik (Verbs) |
    |-----------------|---------------------|
    | /users          | /createUser         |
    | /products       | /getProductDetails  |
    | /orders         | /updateOrderStatus  |
    | /posts          | /createNewPost      |
    | /comments       | /deleteComment      |
    | /categories     | /addCategory        |

2. Naming using plurals Nouns
   
    | Baik (Plurals Nouns) | Kurang Baik (Singular Nouns) |
    |-----------------------|------------------------------|
    | /users                | /user                        |
    | /products             | /product                     |
    | /orders               | /order                       |
    | /posts                | /post                        |
    | /comments             | /comment                     |
    | /categories           | /category                    |

3. using resourse nesting to show relations or hierarchy
   
   /users <- user`s list
   /user/1 <- spesific user
   /user/1/orders <- order list that belongs to spesific user 
   /user/1/orders/12<- spesific order to spesifis user    
  