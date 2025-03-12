package handler

import (
	"context"

	"github.com/Melidee/goth-chat/model"
	"github.com/Melidee/goth-chat/view/user"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type UsersHandler struct {
	DB *bun.DB
}

func (h UsersHandler) HandleUsersShow(c echo.Context) error {
	u := new(model.User)
	h.DB.NewSelect().Model(u).Limit(1).Scan(context.Background())
	return render(c, user.Show(*u))
}