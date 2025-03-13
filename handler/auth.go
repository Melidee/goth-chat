package handler

import (
	"fmt"
	"net/http"

	"github.com/Melidee/goth-chat/model"
	"github.com/Melidee/goth-chat/view/auth"
	"github.com/alexedwards/argon2id"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	DB *sqlx.DB
}

func (h AuthHandler) LoginShow(c echo.Context) error {
	return render(c, auth.LoginShow(false))
}

func (h AuthHandler) LoginPost(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user := new(model.User)
	err := h.DB.Get(user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		fmt.Println("no user")
		return render(c, auth.LoginShow(true))
	}
	hash := user.PasswordHash
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if !match || err != nil { // TODO: research what to do if failing to check hash
		fmt.Println("no match")
		return render(c, auth.LoginShow(true))
	}
	fmt.Printf("creds: %s %s %t\n", email, hash, match)
	return c.Redirect(http.StatusSeeOther, "/")
}

func (h AuthHandler) RegisterShow(c echo.Context) error {
	return nil
}

func (h AuthHandler) RegisterPost(c echo.Context) error {
	return nil
}
