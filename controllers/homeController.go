package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type HomeController struct {
	Echo *echo.Echo
}

func (h *HomeController) index(c echo.Context) error {
	return c.Render(
		http.StatusOK, "index.html", nil, )
}

func (h *HomeController) room(c echo.Context) error {
	nickname := c.Request().FormValue("nickname")
	http.SetCookie(c.Response().Writer, &http.Cookie{
		Name: "nickname", Value: nickname,
		Expires: time.Now().Add(time.Hour * 120000),
	})
	return c.Render(
		http.StatusOK, "room.html", nil, )
}

func (h *HomeController) Init() {
	router := h.Echo
	router.GET("/", h.index)
	router.POST("/room", h.room)
	router.Any("/room", h.room)
}
