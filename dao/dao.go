package dao

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type User struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func New(db *sqlx.DB) *Dao {
	return &Dao{db: db}
}

type Dao struct {
	db *sqlx.DB
}

func (d *Dao) CreateUser(name string) error {
	if _, err := d.db.Exec(`insert into users (name) values ($1)`, name); err != nil {
		return fmt.Errorf("can't insert user: %s", err)
	}
	return nil
}

func (d *Dao) GetAllUsers() ([]*User, error) {
	users := make([]*User, 0)
	err := d.db.Select(&users, `select id, name, created_at from users`)
	if err != nil {
		return nil, err
	}
	return users, nil
}
