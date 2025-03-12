package model

type User struct {
	ID int64 `db:id`
	Name string `db:name`
	ProfilePicture string `db:profilePicture`
	Email string `db:email`
	PasswordHash string `db:passwordHash`
}