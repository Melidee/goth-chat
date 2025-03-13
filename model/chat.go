package model

import "github.com/jmoiron/sqlx"

type Chat struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}

func ChatSchema() string {
	return `
	CREATE TABLE Chats (
		id 			INTEGER PRIMARY KEY AUTOINCREMENT,
		name 		TEXT NOT NULL
	);
	`
}

func (c Chat) Members(db sqlx.DB) ([]User, error) {
	var members []User
	err := db.Select(members, `SELECT * FROM Users JOIN ChatUserJunctions USING (user)`)
	return members, err
}

func (c Chat) Messages(db sqlx.DB) ([]User, error) {
	var members []User
	err := db.Select(members, `SELECT * FROM Messages JOIN ChatMessageJunctions USING (message)`)
	return members, err
}

type ChatUser struct {
	User int64 `db:"user"`
	Chat int64 `db:"chat"`
}

func ChatUserSchema() string {
	return `
	CREATE TABLE ChatUserJunctions (
		user	INTEGER NOT NULL,
		chat	INTEGER NOT NULL,
		FOREIGN KEY (user) REFERENCES Users(id),
		FOREIGN KEY (chat) REFERENCES Chats(id)
	);
	`
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
