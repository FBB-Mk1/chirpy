package database

import (
	"cmp"
	"slices"
)

// CreateChirp creates a new chirp and saves it to disk
func (db *DB) CreateChirp(body string) (Chirp, error) {
	database, err := db.loadDB()
	if err != nil {
		return Chirp{}, err
	}
	chirps := chirpMapToSlice(database.Chirps)
	slices.SortFunc(chirps, func(i, j Chirp) int {
		return cmp.Compare(i.Id, j.Id)
	})

	nextId := 1
	if len(chirps) > 0 {
		nextId = chirps[len(chirps)-1].Id + 1
	}
	newChirp := Chirp{Id: nextId, Body: body}
	chirps = append(chirps, newChirp)
	chirpMap := chirpSliceToMap(chirps)
	database.Chirps = chirpMap
	err = db.writeDB(database)
	if err != nil {
		return Chirp{}, err
	}
	return newChirp, nil
}

// GetChirps reads from disk and returns to reader
func (db *DB) GetChirps() ([]Chirp, error) {
	database, err := db.loadDB()
	if err != nil {
		return []Chirp{}, err
	}
	chirpSlice := []Chirp{}
	for _, v := range database.Chirps {
		chirpSlice = append(chirpSlice, Chirp{v.Id, v.Body})
	}
	return chirpSlice, nil
}

func chirpSliceToMap(chirps []Chirp) map[int]Chirp {
	chirpValues := make(map[int]Chirp)
	for _, val := range chirps {
		chirpValues[val.Id] = val
	}
	return chirpValues
}

func chirpMapToSlice(chirps map[int]Chirp) []Chirp {
	chirpValues := []Chirp{}
	for _, val := range chirps {
		chirpValues = append(chirpValues, val)
	}
	return chirpValues
}
