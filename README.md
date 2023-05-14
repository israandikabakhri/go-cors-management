# go-cors-management
Project Go Untuk Melakukan Manajemen CORS


CORS (Cross-Origin Resource Sharing) adalah mekanisme keamanan browser yang membatasi permintaan HTTP lintas domain. Mekanisme ini mencegah permintaan yang berasal dari sebuah sumber yang berbeda (origins) dari sumber yang diberikan oleh server.

Dalam aplikasi web modern, sering kali terjadi dimana halaman web yang dibuka di browser memuat sumber daya (resource) dari beberapa domain yang berbeda. Misalnya, sebuah halaman HTML yang memuat gambar dari domain yang berbeda atau halaman JavaScript yang memuat data dari API server. Jika tidak dikelola dengan baik, kebijakan CORS yang diimplementasikan oleh browser akan mencegah akses sumber daya lintas domain tersebut, sehingga menyebabkan error atau masalah lain pada aplikasi web.

Oleh karena itu, manajemen CORS digunakan untuk mengelola kebijakan CORS pada aplikasi web, sehingga memungkinkan sumber daya lintas domain yang diizinkan untuk diakses oleh aplikasi web. Dengan cara ini, pengembang dapat mengizinkan sumber daya yang dibutuhkan oleh aplikasi web diambil dari domain yang berbeda.



## Dalam project ini melibatkan 3 file yaitu:

```
go-cors-management/main.go
```

## Perbedaan penerapan management CORS di main.go

### Mengijinkan semua sumber
```
r.Use(cors.Default())
```

### Mengijinkan pada sumber yang dibatasi hanya beberapa saja
```
c := cors.DefaultConfig()
c.AllowOrigins = []string{"https://www.contoh.com", "http://localhost:3200", "https://www.ini-contoh.co.id"}
r.Use(cors.New(c))
```

## Cara Menjalankan Project:

1. Buka **CMD**
2. Masuk ke Folder project melalui **CMD**
3. ketik **go run main.go** 


## Collection Postman

Untuk mempermudah terting data saya telah menyiapkan di folder ini:
```
go-cors-management/postman-collection/Go - CRUD.postman_collection.json
```
Silahkan import kedalam postman anda

Enjoyy..

Jika ada kendala jangan sungkan email ke andikaisra7@gmail.com


## Catatan Penggunaan Tiap Folder:

- app/Function -> Fungsinya sebagai penyimpanan query kompleks atau fungsi query yang memiliki relasi yang kompleks yang nanti sisa dipanggil oleh controller
- app/Helpers -> Fungsinya sebagai penyimpanan file helper seperti mengubah format tanggal, waktu, uang dan lain-lain
- app/Http/Controllers -> Fungsinya sebagai penyimpanan seluruh controller
- app/Http/Requests -> Fungsinya sebagai penyimpanan semua validator requests yang nanti akan dipanggil di controller untuk validasi input data seperti simpan dan update
- app/Model -> Fungsinya sebagai penyimpanan folder-folder Model, Connect Database serta migration

## Catatan Batasan Penggunaan Komponen:

- Model: Komponen yang berinteraksi dengan database dan pengubah format data seperti mengubah format tanggal yang dipanggil dari Helper
- Helper: Komponen fungsi yang dapat dipanggil dari mana saja
- Controller: Komponen yang mengelola pemanggilan data dari model, pemanggilan validasi requests, pemanggilan Functions Query Kompleks dan Pengelolaan Rest API
- Request: Komponen yang mengelola validasi input dari user termasuk keamanan SQL Injection dan XSS Script
- Main.go: Komponen yang mengelola Routing dan Pengelolaan hak akses serta authentikasi menggunakan JWT auth
