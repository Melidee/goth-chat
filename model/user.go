package model

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID int64 `bun:",pk,autoincrement"`
	Name string
	ProfilePicture string
	Email string
	Username string
	PasswordHash string
}