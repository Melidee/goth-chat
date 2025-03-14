package model

import "time"

type Message struct {
	ID        int64      `db:"id"`
	Author    int64      `db:"author"`
	Chat      int64      `db:"chat"`
	Timestamp *time.Time `db:"timestamp"`
	Body      string     `db:"body"`
	Media     string     `db:"media"`
}

func MessageSchema() string {
	return `
	CREATE TABLE Messages (
		id 			INTEGER PRIMARY KEY AUTOINCREMENT,
		author 		INTEGER NOT NULL,
		chat        INTEGER NOT NULL,
		timestamp 	DATETIME DEFAULT CURRENT_TIMESTAMP,
		body 		TEXT,
		media   	TEXT,
		FOREIGN KEY (author) REFERENCES Users(id)
		FOREIGN KEY (chat) REFERENCES DirectChats(id)
	);
	`
}
