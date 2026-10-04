package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func ConnectDB() {
	// Memuat file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment system standar")
	}

	// Menyusun connection string secara aman dari variabel lingkungan
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Membuka koneksi ke PostgreSQL
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database PostgreSQL: %v\n", err)
	}

	// Tes ping koneksi
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Database PostgreSQL tidak merespons: %v\n", err)
	}

	DB = conn
	fmt.Println("Berhasil terhubung ke database PostgreSQL!")
}
