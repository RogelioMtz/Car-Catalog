package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	_ "modernc.org/sqlite" // SQLite driver
)

// Car represents the structure of a car in the catalog
type Car struct {
	Marca     string  `json:"marca"`
	Modelo    string  `json:"modelo"`
	Color     string  `json:"color"`
	Año       int     `json:"año"`
	Precio    float64 `json:"precio"`
	DateAdded string  `json:"dateAdded"`
}

var db *sql.DB

func main() {
	var err error
	// Open SQLite database (creates the file if it doesn't exist)
	db, err = sql.Open("sqlite", "./car_catalog.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the cars table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cars (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		marca TEXT,
		modelo TEXT,
		color TEXT,
		ano INTEGER,
		precio REAL,
		date_added TEXT
	)`)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/api/cars", GetCarsHandler).Methods("GET")
	r.HandleFunc("/api/cars", AddCarHandler).Methods("POST")

	// Serve the webpage
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../src")))

	// Graceful shutdown
	go handleShutdown()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", r); err != nil {
		log.Fatal(err)
	}
}

// GetCarsHandler handles GET requests to retrieve all cars
func GetCarsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT marca, modelo, color, ano, precio, date_added FROM cars`)
	if err != nil {
		http.Error(w, "Failed to fetch cars", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cars []Car
	for rows.Next() {
		var car Car
		if err := rows.Scan(&car.Marca, &car.Modelo, &car.Color, &car.Año, &car.Precio, &car.DateAdded); err != nil {
			http.Error(w, "Failed to parse cars", http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// AddCarHandler handles POST requests to add a new car
func AddCarHandler(w http.ResponseWriter, r *http.Request) {
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Set the current date and time in "AAAA-MM-DD HH:MM:SS" format
	car.DateAdded = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.Exec(`INSERT INTO cars (marca, modelo, color, ano, precio, date_added)
		VALUES (?, ?, ?, ?, ?, ?)`,
		car.Marca, car.Modelo, car.Color, car.Año, car.Precio, car.DateAdded)
	if err != nil {
		http.Error(w, "Failed to add car", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

// Handle graceful shutdown
func handleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("Shutting down server...")
	os.Exit(0)
}
