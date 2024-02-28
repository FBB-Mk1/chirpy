package database

import "sync"

type Chirp struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

// NewDB creates a new database connection
// and creates the database file if it doesn't exist
func NewDB(path string) (*DB, error) {
	return &DB{}, nil
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {
	return Chirp{}, nil
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) GetChirps() ([]Chirp, error) {
	return []Chirp{}, nil
}

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error {
	_, err := db.loadDB()
	return err
}

// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error) {
	err := db.ensureDB()
	return DBStructure{}, err
}

// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {
	return nil
}
