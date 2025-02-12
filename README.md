# Microservice Golang

## 📌 Overview

Project ini merupakan implementasi microservice backend menggunakan Golang dengan package Gin dan GORM. Backend menggunakan database MySQL dan mendukung fitur CRUD untuk artikel.

## 🛠️ Teknologi yang Digunakan

- Backend: Golang, Gin, GORM, MySQL
- Database: MySQL
- Tools: Postman (untuk testing API)

## 📂 Struktur Project

```
/microservice
 ├── config/       # Konfigurasi database
 ├── controllers/  # Logic untuk setiap endpointny
 ├── models/       # Model database
 ├── main.go       # Entry point backend
```

## 🔧 Instalasi

#### 1. Pastikan sudah menginstal Golang, MySQL, dan Postman

#### 2. Clone repository, jalankan perintah di bawah pada terminal

```
git clone https://github.com/your-repo/microservice-golang-react.git
```

#### 3. Masuk ke project
```
cd microservice
code .
```

#### 4. Setup database MySQL dan sesuaikan config pada file config/database.go
```
	dbUser := flag.String("db_user", "root", "Database user")
	dbPassword := flag.String("db_password", "", "Database password")
	dbHost := flag.String("db_host", "localhost", "Database host")
	dbPort := flag.String("db_port", "3306", "Database port")
	dbName := flag.String("db_name", "article", "Database name")
```
ubah parameter ke 2 dan sesuaikan dengan database MySQL anda

#### 5. Jalankan perintah berikut untuk membuat file go.mod (untuk mengatur dependencies)
```
go mod init microservice
```

#### 6. Jalankan perintah di bawah untuk menginstall dependencies yang diperlukan

```
go get -u gorm.io/gorm
go get gorm.io/driver/mysql
go get github.com/gin-gonic/gin
```

#### 6. Terakhir jalankan project

```
go run main.go
```

> [!NOTE]
> untuk testing apakah microservice di atas sudah sesuai, bisa menggunakan collection "Microservice Article.postman_collection json" import file tersebut pada POSTMAN


