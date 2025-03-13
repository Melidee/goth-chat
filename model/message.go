package model

type Message struct {
	ID     int64  `db:"id"`
	Author int64  `db:"author"`
	Body   string `db:"body"`
	Media  string `db:"media"`
}

func MessageSchema() string {
	return `
	CREATE TABLE Messages (
		id 			INTEGER PRIMARY KEY AUTOINCREMENT,
		author 		INTEGER NOT NULL,
		timestamp 	DATETIME DEFAULT CURRENT_TIMESTAMP,
		body 		TEXT,
		media   	TEXT,
		FOREIGN KEY (author) REFERENCES Users(id)
	);
	`
}
