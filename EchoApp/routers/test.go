package routers

import (
	"TestGO/EchoApp/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Message struct {
	Msg string `json:"message"`
}

func InitTestRoutes(e *echo.Echo) {
	e.GET("/hello_world", HelloWorld)
}

// HelloWorld godoc
// @Summary hello world sencillo.
// @Description Endpoint que muestra el HelloWorld de la variable ingresada.
// @Param nombre query string true "Valor que se mostrará en la respuesta"
// @Tags HelloWorld
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /hello_world [get]
func HelloWorld(c echo.Context) error {
	query := c.QueryParam("nombre")

	if val, err := controllers.ControllerHelloWorld(query); err != nil {
		return c.JSON(http.StatusBadRequest, Message{
			Msg: err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, Message{
			Msg: val,
		})
	}

}
