package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type Chirp struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
	Users  map[int]User  `json:"users"`
}

// NewDB creates a new database connection
// and creates the database file if it doesn't exist
func NewDB(path string) (*DB, error) {
	fullpath := path + "/database.json"
	db := &DB{path: fullpath, mux: &sync.RWMutex{}}
	err := db.ensureDB()
	return db, err
}

// ensureDB creates a new database file if it doesn't exist
func (db *DB) ensureDB() error {
	_, err := os.OpenFile(db.path, os.O_RDWR|os.O_CREATE, 0644)
	return err
}

// loadDB reads the database file into memory
func (db *DB) loadDB() (DBStructure, error) {
	data, err := os.ReadFile(db.path)
	dbStruct := DBStructure{}
	if err != nil {
		return dbStruct, err
	}
	if len(data) > 0 {
		err = json.Unmarshal(data, &dbStruct)
		if err != nil {
			return dbStruct, err
		}
	}
	return dbStruct, nil
}

func (db *DB) GetChirpbyID(id int) (Chirp, error) {
	data, err := db.loadDB()
	chirp := Chirp{}
	if err != nil {
		return chirp, err
	}
	chirp, ok := data.Chirps[id]
	if !ok {
		return chirp, errors.New("Chirp not found")
	}
	return chirp, nil
}

// writeDB writes the database file to disk
func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()
	data, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, data, 0000)
	if err != nil {
		return err
	}
	return nil
}
