package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Melidee/goth-chat/handler"
	"github.com/Melidee/goth-chat/model"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	app := echo.New()
	db, err := initDB("file::memory:?cache=shared")
	if err != nil {
		app.Logger.Fatal(err)
	}

	fillDB(context.Background(), db)

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userHandler := handler.UsersHandler{DB: db}
	app.GET("/users", userHandler.HandleUsersShow)
	app.Logger.Fatal(app.Start(":8080"))
}

// dbFile `file::memory:?cache=shared` for in memory database
func initDB(dbFile string) (*bun.DB, error) {
	sqldb, err := sql.Open(sqliteshim.ShimName, dbFile)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true), bundebug.FromEnv("BUNDEBUG")))
	return db, nil
}

func fillDB(ctx context.Context, db *bun.DB) {
	db.NewCreateTable().Model((*model.User)(nil)).Exec(ctx)
	db.NewInsert().Model(&[]model.User{{Name: "Amelia", ProfilePicture: "/assets/default.webp", Email: "amelia@example.com", Username: "meli", PasswordHash: "$2a$10$FlaqHRKfzsprw79tqIJNuOyXIljFZNF.NivRy7WNZpwpMINoKNBzm"}}).Exec(ctx)
}
