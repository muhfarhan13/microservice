package controllers

import (
	"microservice/config"      	// Import konfigurasi database
	"microservice/models"      	// Import model untuk tabel Posts
	"net/http"                	// Import HTTP untuk menangani request & response
	"strconv"                 	// Import strconv untuk konversi string ke int
	"time"                    	// Import time untuk timestamp
	"github.com/gin-gonic/gin" 	// Import Gin
)

// Validasi artikel sebelum disimpan
func validateArticle(article models.Posts) string {
	if len(article.Title) < 20 { // Judul minimal 20 karakter
		return "Title minimal 20 karakter"
	}
	if len(article.Content) < 200 { // Konten minimal 200 karakter
		return "Content minimal 200 karakter"
	}
	if len(article.Category) < 3 { // Kategori minimal 3 karakter
		return "Category minimal 3 karakter"
	}
	// Status hanya boleh "publish", "draft", atau "trash"
	if article.Status != "publish" && article.Status != "draft" && article.Status != "trash" {
		return "Status harus 'publish', 'draft', atau 'trash'"
	}
	return "" // Jika validasi lolos, return string kosong
}

// Membuat artikel baru
func CreateArticle(c *gin.Context) {
	var article models.Posts

	// Set timestamp untuk created_date & updated_date
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	article.CreatedDate = currentTime
	article.UpdatedDate = currentTime
	
	// Binding JSON dari request ke struct article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input artikel
	if msg := validateArticle(article); msg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	// Simpan artikel ke database
	config.DB.Create(&article)
	c.JSON(http.StatusCreated, gin.H{"message": "Article created"})
}

// Mengambil daftar artikel dengan pagination
func GetArticles(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "5")  // Ambil limit dari query param, default 5
	offsetStr := c.DefaultQuery("offset", "0") // Ambil offset dari query param, default 0

	// Konversi limit dan offset ke integer
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	var articles []models.Posts
	// Ambil data artikel dengan batasan limit & offset
	config.DB.Limit(limit).Offset(offset).Find(&articles)

	// Kirim response dalam bentuk JSON
	c.JSON(http.StatusOK, articles)
}

// Mengambil detail artikel berdasarkan ID
func GetArticle(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	var article models.Posts

	// Cari artikel berdasarkan ID
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Kirim response jika artikel ditemukan
	c.JSON(http.StatusOK, article)
}

// Mengupdate artikel berdasarkan ID
func UpdateArticle(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL
	var article models.Posts
	
	// Set timestamp untuk updated_date
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	article.UpdatedDate = currentTime
	
	// Cari artikel berdasarkan ID
	if err := config.DB.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Binding JSON dari request ke struct article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi input artikel
	if msg := validateArticle(article); msg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": msg})
		return
	}

	// Simpan perubahan ke database
	config.DB.Save(&article)

	// Kirim response jika artikel berhasil diubah
	c.JSON(http.StatusOK, gin.H{"message": "Article updated"})
}

// Menghapus artikel berdasarkan ID
func DeleteArticle(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	// Hapus artikel berdasarkan ID
	if err := config.DB.Delete(&models.Posts{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Kirim response jika artikel berhasil dihapus
	c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
}