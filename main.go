package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.core/go-cors-management/app/Http/Controllers/penduduk"
	"main.core/go-cors-management/config/database"
	"main.core/go-cors-management/migration"
)

func init() {
	database.ConnectDatabase()
	migration.Migration()
}

func main() {

	r := gin.Default()

	// Menambahkan middleware CORS
	r.Use(cors.Default())

	// Jika CORS diberikan secara spesifik maka gunakan kode dibawah ini
	c := cors.DefaultConfig()
	c.AllowOrigins = []string{"https://www.contoh.com", "http://localhost:3200", "https://www.ini-contoh.co.id"}
	r.Use(cors.New(c))

	r.GET("/api/penduduk", penduduk.Data)
	r.GET("/api/penduduk/:nik", penduduk.Detail)
	r.POST("/api/penduduk/search", penduduk.Search)
	r.POST("/api/penduduk", penduduk.Create)
	r.PUT("/api/penduduk/:nik", penduduk.Update)
	r.DELETE("/api/penduduk", penduduk.Delete)

	r.Run()

}
