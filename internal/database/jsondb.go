package database

import (
	"cmp"
	"encoding/json"
	"errors"
	"os"
	"slices"
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
	chirps, err := db.GetChirps()
	slices.SortFunc(chirps, func(i, j Chirp) int {
		return cmp.Compare(i.Id, j.Id)
	})
	if err != nil {
		return Chirp{}, err
	}
	nextId := 1
	if len(chirps) > 0 {
		nextId = chirps[len(chirps)-1].Id + 1
	}
	newChirp := Chirp{Id: nextId, Body: body}
	chirps = append(chirps, newChirp)
	err = db.writeDB(chirpSliceToStruct(chirps))
	if err != nil {
		return Chirp{}, err
	}
	return newChirp, nil
}

func chirpSliceToStruct(chirps []Chirp) DBStructure {
	dbStruct := make(map[int]Chirp)
	for _, val := range chirps {
		dbStruct[val.Id] = val
	}
	return DBStructure{dbStruct}
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
