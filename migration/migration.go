package migration

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"main.core/go-cors-management/app/Model/penduduk"
	"main.core/go-cors-management/config/database"
)

var DB = database.ConnectDatabase()

func Migration() {
	dsn := os.Getenv("MARIADB")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	database.Table("penduduk").AutoMigrate(&penduduk.Penduduk{})

	//database.Table("penduduk").Create(&Penduduk{Nik: "9999", Nama: "xxx"})
}
