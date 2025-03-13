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
	db, err := sqlx.Connect("sqlite3", ":memory:")
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
	users := []model.User{
		{
			Name:           "Amelia",
			ProfilePicture: "/assets/default.webp",
			Email:          "amelia@example.com",
			PasswordHash:   "$argon2id$v=19$m=65536,t=1,p=12$oKcFeeeCLbJ+MJkLE21lAg$BREuFcbA/AVS2KFwxlqXcE90sJ8fuDbsxRaq96UGQXI",
		},
	}
	tx := db.MustBegin()
	for _, user := range users {
		_, _ = tx.NamedExec(`
			INSERT INTO Users  (name, profilePicture, email, passwordHash) 
			VALUES (:name, :profilePicture, :email, :passwordHash)
		`, user)
	}
	_ = tx.Commit()

	var usersGet []model.User
	err := db.Select(&usersGet, "SELECT * FROM Users")
	if err != nil {
		panic(err)
	}
	println(usersGet)
}
