package model

type User struct {
	ID int64 `db:"id"`
	Name string `db:"name"`
	ProfilePicture string `db:"profilePicture"`
	Email string `db:"email"`
	PasswordHash string `db:"passwordHash"`
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