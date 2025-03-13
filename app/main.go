package main

import (
	"net/http"

	"github.com/Melidee/goth-chat/handler"
	"github.com/Melidee/goth-chat/model"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	app := echo.New()
	db, err := sqlx.Connect("sqlite3", "sqlite3.db")
	if err != nil {
		app.Logger.Fatal(err)
	}
	fillDB(db)

	app.Static("/assets", "public/assets")

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userHandler := handler.UsersHandler{DB: db}
	app.GET("/users", userHandler.HandleUsersShow)

	authHandler := handler.AuthHandler{DB: db}
	app.GET("/login", authHandler.LoginShow)
	app.POST("/login", authHandler.LoginPost)
	app.GET("/register", authHandler.RegisterShow)
	app.POST("/register", authHandler.RegisterPost)

	app.Logger.Fatal(app.Start(":8080"))
}

func fillDB(db *sqlx.DB) {
	db.MustExec(model.UserSchema())
	db.MustExec(model.MessageSchema())
	db.MustExec(model.ChatSchema())
	db.MustExec(model.ChatUserSchema())
	db.MustExec(model.ChatMessageSchema())
	users := []model.User{
		{
			ID: 1,
			Name:           "Amelia",
			ProfilePicture: "/assets/default.webp",
			Email:          "amelia@example.com",
			PasswordHash:   "$argon2id$v=19$m=65536,t=1,p=12$oKcFeeeCLbJ+MJkLE21lAg$BREuFcbA/AVS2KFwxlqXcE90sJ8fuDbsxRaq96UGQXI",
		},
		{
			ID: 2,
			Name:           "Drew",
			ProfilePicture: "/assets/default.webp",
			Email:          "drew@example.com",
			PasswordHash:   "$argon2id$v=19$m=65536,t=1,p=12$oKcFeeeCLbJ+MJkLE21lAg$BREuFcbA/AVS2KFwxlqXcE90sJ8fuDbsxRaq96UGQXI",
		},
	}
	chats := []model.Chat{
		{
			ID: 1,
			Name: "this chat",
		},
	}
	messages := []model.Message{
		{
			ID: 1,
			Author: 1,
			Body: "Hello, world!",
			Media: "",
		},
		{
			ID: 2,
			Author: 1,
			Body: "Hello, again!",
			Media: "",
		},
		{
			ID: 3,
			Author: 2,
			Body: "Hello, friend!",
			Media: "",
		},
	}
	chatUserJunctions := []model.ChatUser{
		{
			User: 1,
			Chat: 1,
		},
		{
			User: 2,
			Chat: 1,
		},
	}
	chatMessageJunctions := []model.ChatMessage{
		{
			Message: 1,
			Chat: 1,
		},
		{
			Message: 2,
			Chat: 1,
		},
		{
			Message: 3,
			Chat: 1,
		},
	}
	tx := db.MustBegin()
	for _, user := range users {
		_, _ = tx.NamedExec(`
			INSERT INTO Users (id, name, profilePicture, email, passwordHash) 
			VALUES (:id, :name, :profilePicture, :email, :passwordHash);
		`, user)
	}
	for _, chat := range chats {
		_, _ = tx.NamedExec(`
			INSERT INTO Chats (id, name) 
			VALUES (:id, :name);
		`, chat)
	}
	for _, message := range messages {
		_, _ = tx.NamedExec(`
			INSERT INTO Messages (id, author, body, media) 
			VALUES (:id, :author, :body, :media);
		`, message)
	}
	for _, junction := range chatUserJunctions {
		_, _ = tx.NamedExec(`
			INSERT INTO ChatUserJunctions (user, chat) 
			VALUES (:user, :chat);
		`, junction)
	}
	for _, junction := range chatMessageJunctions {
		_, _ = tx.NamedExec(`
			INSERT INTO ChatMessageJunctions (message, chat) 
			VALUES (:message, :chat);
		`, junction)
	}
	_ = tx.Commit()
}
