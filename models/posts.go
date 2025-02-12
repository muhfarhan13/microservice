package models

import (
	"gorm.io/gorm"
)

// Struktur model untuk tabel "posts"
type Posts struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"` // ID utama, auto-increment
	Title       string `json:"title" gorm:"not null;size:255"`     // Judul artikel (maksimal 255 karakter)
	Content     string `json:"content" gorm:"not null"`            // Isi artikel (tidak boleh kosong)
	Category    string `json:"category" gorm:"not null;size:50"`   // Kategori artikel (maksimal 50 karakter)
	CreatedDate string `gorm:"type:timestamp"`                     // Tanggal dibuat (format timestamp)
	UpdatedDate string `gorm:"type:timestamp"`                     // Tanggal diperbarui (format timestamp)
	Status      string `json:"status" gorm:"not null;size:10"`     // Status artikel (misalnya: "draft" atau "publish")
}

// Fungsi untuk migrasi database, membuat tabel jika belum ada
func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&Posts{}) // AutoMigrate akan membuat atau memperbarui tabel sesuai struktur Posts
}