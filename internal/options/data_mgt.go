package options

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

// CreateImpliedVolatilityTable creates the ImpliedVolatility table in the database.
func CreateImpliedVolatilityTable(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS ImpliedVolatility (
	    id SERIAL PRIMARY KEY,
        option_symbol VARCHAR(255) NOT NULL,
        implied_volatility FLOAT NOT NULL,
        calculation_date TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating ImpliedVolatility table: %v", err)
	} else {
		log.Println("ImpliedVolatility table created successfully.")
	}
}

// CreateHistoricalDataTable creates the HistoricalData table in the database.
func CreateHistoricalDataTable(db *sql.DB) {
	createTableQuery := `
	CREATE TABLE HistoricalData (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    volume BIGINT,
    date TIMESTAMP WITHOUT TIME ZONE NOT NULL
	);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating HistoricalData table: %v", err)
	} else {
		log.Println("HistoricalData table created successfully.")
	}
}

// SaveIVData saves the calculated IV data to the database.
func SaveIVData(optionSymbol string, iv float64) error {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return err
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO ImpliedVolatility (option_symbol, implied_volatility) VALUES ($1, $2)", optionSymbol, iv)
	if err != nil {
		log.Printf("Error inserting data: %v\n", err)
		return err
	}

	return nil
}

// SaveHistoricalData saves historical price data to the database.
func SaveHistoricalData(db *sql.DB, symbol string, price float64, volume int64, date time.Time) error {
	_, err := db.Exec("INSERT INTO HistoricalData (symbol, price, volume, date) VALUES ($1, $2, $3, $4)", symbol, price, volume, date)
	if err != nil {
		log.Printf("Error inserting historical data: %v\n", err)
		return err
	}

	return nil
}

// GetHistoricalData retrieves historical data for a given symbol and time range.
func GetHistoricalData(db *sql.DB, symbol string, startDate, endDate time.Time) ([]HistoricalData, error) {
	rows, err := db.Query("SELECT symbol, price, volume, date FROM HistoricalData WHERE symbol = $1 AND date BETWEEN $2 AND $3", symbol, startDate, endDate)
	if err != nil {
		log.Printf("Error retrieving historical data: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var data []HistoricalData
	for rows.Next() {
		var hd HistoricalData
		if err := rows.Scan(&hd.Symbol, &hd.Price, &hd.Volume, &hd.Date); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		data = append(data, hd)
	}

	return data, nil
}

// GetDBConnection establishes a connection to the PostgreSQL database.
func GetDBConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %v\n", err)
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}

	return db
}
