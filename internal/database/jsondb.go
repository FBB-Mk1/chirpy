package database

import (
	"fmt"
	"sync"
)

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
	fullpath := path + "/database.json"
	db := &DB{path: fullpath, mux: &sync.RWMutex{}}
	err := db.ensureDB()
	return db, err
}

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {
	chirps, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}
	db.writeDB(chirps)
	return Chirp{}, nil
}

// GetChirps reads from disk and returns to reader
func (db *DB) GetChirps() ([]Chirp, error) {
	chirps, err := db.loadDB()
	chirpSlice := []Chirp{}
	if err != nil {
		return chirpSlice, err
	}
	for _, v := range chirps.Chirps {
		chirpSlice = append(chirpSlice, Chirp{v.Id, v.Body})
	}
	return chirpSlice, nil
}

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error {
	//
	return nil
}

// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error) {

	return DBStructure{}, nil
}

// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {
	fmt.Println(dbStructure.Chirps)
	return nil
}
