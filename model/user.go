package model

import "github.com/jmoiron/sqlx"

type User struct {
	ID             int64  `db:"id"`
	Name           string `db:"name"`
	ProfilePicture string `db:"profilePicture"`
	Email          string `db:"email"`
	PasswordHash   string `db:"passwordHash"`
}

func UserSchema() string {
	return `
	CREATE TABLE Users (
		id 				INTEGER PRIMARY KEY AUTOINCREMENT,
		name 			TEXT,
		profilePicture 	TEXT,
		email 			TEXT NOT NULL UNIQUE,
		passwordHash 	TEXT NOT NULL
	);
	`
}

func (u User) Chats(db *sqlx.DB) ([]DirectChat, error) {
	var chats []DirectChat
	err := db.Select(&chats, `SELECT * FROM DirectChats WHERE userA=$1 OR userB=$1`, u.ID)
	return chats, err
}