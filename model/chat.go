package model

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DirectChat struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
	UserA int64 `db:"userA"`
	UserB int64 `db:"userB"`
}

func ChatSchema() string {
	return `
	CREATE TABLE DirectChats (
		id 		INTEGER PRIMARY KEY AUTOINCREMENT,
		name 	TEXT NOT NULL,
		userA   INTEGER,
		userB	INTEGER
	);
	`
}

func (c DirectChat) Friend(db sqlx.DB, u User) (*User, error) {
	var userA, userB User
	err := db.Get(userA, `SELECT * FROM Users WHERE id=$1`, c.UserA)
	if err != nil {
		return nil, err
	}
	err = db.Get(userB, `SELECT * FROM Users WHERE id=$1`, c.UserB)
	if err != nil {
		return nil, err
	}
	if u.Email == userA.Email {
		return &userB, nil
	}
	if u.Email == userB.Email {
		return &userA, nil
	}
	return nil, fmt.Errorf("user %s is not a member of this chat", u.Email)
}

func (c DirectChat) Messages(db *sqlx.DB) ([]Message, error) {
	var messages []Message
	err := db.Select(&messages, `SELECT * FROM Messages WHERE chat=$1`, c.ID)
	return messages, err
}

type ChatMessage struct {
	Message int64 `db:"message"`
	Chat    int64 `db:"chat"`
}

func ChatMessageSchema() string {
	return `
	CREATE TABLE ChatMessageJunctions (
		message	INTEGER NOT NULL,
		chat	INTEGER NOT NULL,
		FOREIGN KEY (message) REFERENCES Messages(id),
		FOREIGN KEY (chat) REFERENCES Chats(id)
	);
	`
}
