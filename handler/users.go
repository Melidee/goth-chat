package handler

import (
	"context"

	"github.com/Melidee/goth-chat/model"
	"github.com/Melidee/goth-chat/view/users"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type UsersHandler struct {
	DB *bun.DB
}

func (h UsersHandler) HandleUsersShow(c echo.Context) error {
	var u []model.User
	h.DB.NewSelect().Model(&u).Scan(context.Background())
	return render(c, users.Show(u))
}
