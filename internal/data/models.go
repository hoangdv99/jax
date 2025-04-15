package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Users     UserModel
	Tokens    TokenModel
	Stores    StoreModel
	Tags      TagModel
	Products  ProductModel
	Blacklist BlacklistModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: db},
		Tokens:    TokenModel{DB: db},
		Stores:    StoreModel{DB: db},
		Tags:      TagModel{DB: db},
		Products:  ProductModel{DB: db},
		Blacklist: BlacklistModel{DB: db},
	}
}
