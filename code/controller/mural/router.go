package mural

import (
	"mural/controller/mural/routes"
	"mural/model"

	"github.com/labstack/echo/v4"
)

type MuralRouter struct {

}

func NewMuralRouter() MuralRouter{
	return MuralRouter{}
}

func (mc MuralController) GetRoutes() map[string]func(c echo.Context) error {
	router := map[string]func(c echo.Context) error{}
	router["flip-tile"] = routes.FlipTile
	router["mural"] = routes.GetMuaralPage
	router["select-tile"] = routes.SelectTile
	router["select-answer"] = routes.SelectAnswer
	router["submit-answer"] = routes.SubmitAnswer
	router["open-stats-dialog"] = routes.OpenStatsDialog
	router["close-dialog"] = routes.CloseDialog
	return router
}

func (r MuralRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	e.GET("/", c.GetRoutes()["mural"])
	e.PUT("/mural/flip-tile", c.GetRoutes()["flip-tile"])
	e.PUT("/mural/select-tile", c.GetRoutes()["select-tile"])
	e.PUT("/mural/select-answer", c.GetRoutes()["select-answer"])
	e.PUT("/mural/submit-answer", c.GetRoutes()["submit-answer"])
	e.GET("/mural/share-answer", c.GetRoutes()["share-answer"])
	e.GET("/mural/open-stats-dialog", c.GetRoutes()["open-stats-dialog"])
	e.GET("/mural/close-dialog", c.GetRoutes()["close-dialog"])
}