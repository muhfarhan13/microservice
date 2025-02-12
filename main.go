package main

import (
	"microservice/config"      // Import package untuk konfigurasi database
	"microservice/controllers" // Import controller yang menangani request
	"github.com/gin-gonic/gin" // Import Gin sebagai framework web
)

func main() {
	config.ConnectDB() // Koneksi ke database sebelum server berjalan

	r := gin.Default() // Inisialisasi router Gin

	r.POST("/article", controllers.CreateArticle) // Rute untuk membuat artikel baru (POST)
	r.GET("/article", controllers.GetArticles) // Rute untuk mengambil daftar artikel (GET)
	r.GET("/article/:id", controllers.GetArticle) // Rute untuk mengambil detail artikel berdasarkan ID (GET)
	r.PUT("/article/:id", controllers.UpdateArticle) // Rute untuk update artikel berdasarkan ID (PUT)
	r.DELETE("/article/:id", controllers.DeleteArticle) // Rute untuk hapus artikel berdasarkan ID (DELETE)

	r.Run(":8080") // Jalankan server di port 8080
}