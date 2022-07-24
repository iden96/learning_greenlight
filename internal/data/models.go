package data

import (
	"database/sql"
	"errors"
)

var ErrRecordNotFound = errors.New("record not found")

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	movieModel := MovieModel{
		DB: db,
	}

	return Models{
		Movies: movieModel,
	}
}
