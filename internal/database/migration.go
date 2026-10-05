package database

import (
	"context"
	"fmt"
	"log"
)

// RunMigrations menjalankan migrasi database otomatis saat aplikasi start
func RunMigrations() {
	if DB == nil {
		log.Fatal("Koneksi database belum diinisialisasi!")
	}

	ctx := context.Background()

	// Query untuk membuat tabel users (contoh data klien financial planner)
	queryUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// Query untuk membuat tabel portfolios (contoh data aset/keuangan klien)
	queryPortfolios := `
	CREATE TABLE IF NOT EXISTS portfolios (
		id SERIAL PRIMARY KEY,
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		asset_name VARCHAR(100) NOT NULL,
		total_value NUMERIC(15, 2) NOT NULL,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	// Eksekusi migrasi tabel users
	_, err := DB.Exec(ctx, queryUsers)
	if err != nil {
		log.Fatalf("Gagal membuat tabel users: %v\n", err)
	}

	// Eksekusi migrasi tabel portfolios
	_, err = DB.Exec(ctx, queryPortfolios)
	if err != nil {
		log.Fatalf("Gagal membuat tabel portfolios: %v\n", err)
	}

	fmt.Println("Migrasi database berhasil dijalankan: Tabel 'users' & 'portfolios' siap!")
}
