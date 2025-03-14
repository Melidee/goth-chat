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
	err := h.DB.Select(&users, "SELECT * FROM Users")
	if err != nil {
		panic(err)
	}
	chats, err := users[0].Chats(h.DB)
	if err != nil {
		panic(err)
	}
	messages, err := chats[0].Messages(h.DB)
	if err != nil {
		panic(err)
	}
	return render(c, chat.Show(users[0], users, users[1], messages))
}