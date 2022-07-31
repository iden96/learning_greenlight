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
	Movies      MovieModel
	Permissions PermissionModel
	Tokens      TokenModel
	Users       UserModel
}

func NewModels(db *sql.DB) Models {
	movieModel := MovieModel{
		DB: db,
	}

	permissionModel := PermissionModel{
		DB: db,
	}

	tokenModel := TokenModel{
		DB: db,
	}

	userModel := UserModel{
		DB: db,
	}

	return Models{
		Movies:      movieModel,
		Permissions: permissionModel,
		Tokens:      tokenModel,
		Users:       userModel,
	}
}
