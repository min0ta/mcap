package auth

import (
	"encoding/json"
	"log"
	"mcap/internal/utils"
)

type JsonDB struct {
	records []usersRecord
}

func (db *JsonDB) contains(predicate func(usersRecord) bool) Role {
	for _, v := range db.records {
		if predicate(v) {
			return v.Role
		}
	}
	return RoleGuest
}
func (db *JsonDB) Connect(path string) {
	if len(path) < 5 {
		log.Fatal("Invalid json db path", path)
	}
	jsonFile := utils.RequireFile(path)

	if err := json.Unmarshal(jsonFile, &db.records); err != nil {
		log.Fatal("Cannot parse json with error", err)
	}
}

func newJsonDb() *JsonDB {
	return &JsonDB{
		records: []usersRecord{},
	}
}
