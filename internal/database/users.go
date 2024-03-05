package database

import (
	"cmp"
	"errors"
	"slices"
)

func (db *DB) CreateUser(email string) (User, error) {
	database, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	userMap := database.Users
	ok := checkUniqueUser(userMap, email)
	if !ok {
		return User{}, errors.New("email already in use")
	}
	users := userMapToSlice(userMap)

	nextId := 1
	if len(users) > 0 {
		nextId = users[len(users)-1].Id + 1
	}
	newUser := User{Id: nextId, Email: email}
	users = append(users, newUser)
	userMap = userSliceToMap(users)
	database.Users = userMap
	err = db.writeDB(database)
	if err != nil {
		return User{}, err
	}
	return newUser, nil
}

func checkUniqueUser(userMap map[int]User, email string) bool {
	for _, u := range userMap {
		if u.Email == email {
			return false
		}
	}
	return true
}

func userMapToSlice(userMap map[int]User) []User {
	users := []User{}
	for _, val := range userMap {
		users = append(users, val)
	}
	slices.SortFunc(users, func(i, j User) int {
		return cmp.Compare(i.Id, j.Id)
	})
	return users
}

func userSliceToMap(users []User) map[int]User {
	userMap := make(map[int]User)
	for _, val := range users {
		userMap[val.Id] = val
	}
	return userMap
}
