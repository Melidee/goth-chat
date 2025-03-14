package handler

import (
	"github.com/Melidee/goth-chat/model"
	"github.com/Melidee/goth-chat/view/chat"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	DB *sqlx.DB
}

func (h ChatHandler) ChatShow(c echo.Context) error {
	var users []model.User
	h.DB.Select(&users, "SELECT * FROM Users")
	return render(c, chat.Show(users[0], users))
}