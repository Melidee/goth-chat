package handler

import (
	"github.com/Melidee/goth-chat/model"
	"github.com/Melidee/goth-chat/view/user"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type UserHandler struct {
	DB *bun.DB
}

func (h UserHandler) HandleUsersShow(c echo.Context) error {
	u := model.User{Email: "amelia@example.com"}
	return render(c, user.Show(u))
}