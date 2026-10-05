package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// Memuat file .env (opsional jika root sudah memuatnya)
	_ = godotenv.Load()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Membuat connection pool untuk concurrency yang aman
	var err error
	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Gagal membuat connection pool PostgreSQL: %v\n", err)
	}

	// Tes ping koneksi
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Database PostgreSQL tidak merespons: %v\n", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL (Pool Mode)!")
}
