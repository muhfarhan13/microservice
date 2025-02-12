package main

import (
	"microservice/config"      // Import package untuk konfigurasi database
	"microservice/controllers" // Import controller yang menangani request
	"github.com/gin-gonic/gin" // Import Gin sebagai framework web
	"github.com/gin-contrib/cors" // Tambahkan ini
	"time"
)

func main() {
	config.ConnectDB() // Koneksi ke database sebelum server berjalan

	r := gin.Default() // Inisialisasi router Gin

	// Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Hanya izinkan permintaan dari frontend yang berjalan di localhost:5173
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Mengizinkan metode HTTP yang boleh digunakan
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Header yang diizinkan dalam permintaan
		ExposeHeaders:    []string{"Content-Length"}, // Header yang boleh diakses oleh frontend setelah permintaan berhasil
		AllowCredentials: true, // Mengizinkan penggunaan cookie dan header Authorization untu request lintas origin
		MaxAge:           12 * time.Hour, // Menyimpan izin akses selama 12 jam agar browser tidak perlu mengecek ulang setiap kali mengirim request
	}))

	r.POST("/article", controllers.CreateArticle) // Rute untuk membuat artikel baru (POST)
	r.GET("/article", controllers.GetArticles) // Rute untuk mengambil daftar artikel (GET)
	r.GET("/article/:id", controllers.GetArticle) // Rute untuk mengambil detail artikel berdasarkan ID (GET)
	r.PUT("/article/:id", controllers.UpdateArticle) // Rute untuk update artikel berdasarkan ID (PUT)
	r.DELETE("/article/:id", controllers.DeleteArticle) // Rute untuk hapus artikel berdasarkan ID (DELETE)

	r.Run(":8080") // Jalankan server di port 8080
}