package db

import (
	"airline-voucher-seat-be/config"
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitVouchersDB(ctx context.Context) *sql.DB {
	dbDir := "./data"
	dbName := config.GetEnv("DATABASE_VOUCHERS", "vouchers.db")
	checkDirectory(dbDir)
	dbPath := dbDir + "/" + dbName
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	schemaVoucherSetup(ctx, db)

	return db
}

func InitAircraftsDB(ctx context.Context) *sql.DB {
	dbDir := "./data"
	dbName := config.GetEnv("DATABASE_AIRCRAFTS", "aircrafts.db")
	checkDirectory(dbDir)
	dbPath := dbDir + "/" + dbName
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	schemaAircraftSetup(ctx, db)

	return db
}

func checkDirectory(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal("Error creating data directory: ", err)
		}
	}
}

func schemaVoucherSetup(ctx context.Context, db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		crew_name TEXT,
		crew_id TEXT,
		flight_number TEXT,
		flight_date TEXT,
		aircraft_type TEXT,
		seat1 TEXT,
		seat2 TEXT,
		seat3 TEXT,
		created_at TIMESTAMP
	);
	`
	if _, err := db.ExecContext(ctx, schema); err != nil {
		log.Fatalf("Failed to setup database schema: %v", err)
	}
}

func schemaAircraftSetup(ctx context.Context, db *sql.DB) {
	schema := `
	CREATE TABLE IF NOT EXISTS aircrafts (
		aircraft_type TEXT PRIMARY KEY,
		row_start INTEGER NOT NULL,
		row_end INTEGER NOT NULL,
		seats_per_row TEXT NOT NULL
	);

	INSERT OR IGNORE INTO aircrafts (aircraft_type, row_start, row_end, seats_per_row) VALUES
		('ATR', 1, 18, 'A,C,D,F'),
		('Airbus 320', 1, 32, 'A,B,C,D,E,F'),
		('Boeing 737 Max', 1, 32, 'A,B,C,D,E,F');
	`
	if _, err := db.ExecContext(ctx, schema); err != nil {
		log.Fatalf("Failed to setup database schema: %v", err)
	}
}
