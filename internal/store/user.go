package store

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

type authType string

const EmailType authType = "email"
const TelegramType authType = "telegram"

type User struct {
	ID         int      `json:"-" db:"id"`
	Email      string   `json:"email" db:"email"`
	TelegramID int      `json:"telegramId" db:"telegram_id"`
	Type       authType `json:"type" db:"type"`
}

func NewEmailUser(email string) User {
	return User{
		Email: email,
		Type:  EmailType,
	}
}

func NewTelegramUser(telegramID int) User {
	return User{
		TelegramID: telegramID,
		Type:       TelegramType,
	}
}

type userRepository interface {
	Create(user User) (int, error)
	ExistsByEmail(email string) (bool, error)
}

type user struct {
	db *sqlx.DB
}

func (u user) Create(user User) (int, error) {
	sqlResult, err := u.db.Exec("insert into users (email, telegram_id, type) values (?, ?, ?) on duplicate key update id=id", user.Email, user.TelegramID, user.Type)
	if err != nil {
		return 0, err
	}

	id, err := sqlResult.LastInsertId()
	return int(id), err
}

func (u user) ExistsByEmail(email string) (bool, error) {
	var id int
	err := u.db.Get(&id, `select id from users where email = ?`, email)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
