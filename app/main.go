package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Melidee/goth-chat/handler"
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

	_, err = db.ExecContext(context.Background(), "SELECT 1")

	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	userHandler := handler.UserHandler{DB: db}
	app.GET("/user", userHandler.HandleUsersShow)
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