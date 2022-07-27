package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel
}

func NewModels(db *sql.DB) Models {
	movieModel := MovieModel{
		DB: db,
	}

	userModel := UserModel{
		DB: db,
	}

	return Models{
		Movies: movieModel,
		Users:  userModel,
	}
}
