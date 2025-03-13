package handler

import (
	"net/http"

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
	err := h.DB.Select(&u, "SELECT * FROM Users")
	if err != nil {
		return c.String(http.StatusInternalServerError, "error fetching from database")
	}
	return render(c, users.Show(u))
}
