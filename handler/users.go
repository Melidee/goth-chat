package handler

import (
	"github.com/Melidee/goth-chat/model"
	"github.com/Melidee/goth-chat/view/users"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	DB *sqlx.DB
}

func (h UsersHandler) HandleUsersShow(c echo.Context) error {
	var u []model.User
	h.DB.Get(u, "SELECT * FROM users")
	println(u)
	return render(c, users.Show(u))
}
