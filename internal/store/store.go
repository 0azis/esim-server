package store

import (
	"esim/config"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	User() userRepository
}

type store struct {
	db *sqlx.DB
}

func (s store) User() userRepository {
	return user{s.db}
}

func New(cfg config.Config) (Store, error) {
	fmt.Println(cfg.Database.Uri())
	db, err := sqlx.Connect("mysql", cfg.Database.Uri())
	if err != nil {
		return store{}, err
	}

	db.SetMaxIdleConns(150)
	db.SetMaxOpenConns(150)
	db.SetConnMaxLifetime(5 * time.Minute)

	return store{db}, nil
}
