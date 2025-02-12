package config

import (
	"fmt"
	"log"
	"flag"
	"microservice/models"       // Import model untuk migrasi
	"gorm.io/driver/mysql"      // Import driver MySQL untuk GORM
	"gorm.io/gorm"              // Import GORM sebagai ORM
)

var DB *gorm.DB // Variabel global untuk menyimpan koneksi database

func ConnectDB() {
	// Mendefinisikan flag untuk konfigurasi database
	dbUser := flag.String("db_user", "root", "Database user")
	dbPassword := flag.String("db_password", "", "Database password")
	dbHost := flag.String("db_host", "localhost", "Database host")
	dbPort := flag.String("db_port", "3306", "Database port")
	dbName := flag.String("db_name", "article", "Database name")
	
	// Parse flags
	flag.Parse()
	
	// dsn := "root:@tcp(localhost:3306)/article?charset=utf8mb4&parseTime=True&loc=Local"
	// Data Source Name (DSN) untuk koneksi ke database MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", *dbUser, *dbPassword, *dbHost, *dbPort, *dbName)
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Buka koneksi ke database
	if err != nil {
		log.Fatal("Failed to connect to database:", err) // Jika gagal, tampilkan error dan hentikan program
	}

	fmt.Println("Connected to MySQL database") // Log jika koneksi berhasil
	models.MigrateDB(DB) // Jalankan migrasi tabel berdasarkan model yang sudah dibuat
}