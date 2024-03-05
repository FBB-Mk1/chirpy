package database

func (db *DB) CreateUser(body string) (User, error) {
	database, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	userMap := database.Users
	db.writeDB(DBStructure{database.Chirps, userMap})
	return User{}, nil
}
