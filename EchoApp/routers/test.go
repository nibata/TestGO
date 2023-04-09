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
	e.GET("/helloworld", helloworld)
}

// helloworld godoc
// @Summary hello world sencillo.
// @Description Endpoint que muestra el helloworld de la variable ingresada.
// @Param nombre path string true "Valor que se mostrará en la respuesta"
// @Tags helloworld
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /helloworld [get]
func helloworld(c echo.Context) error {
	query := c.QueryParam("nombre")

	if val, err := controllers.HelloWorld(query); err != nil {
		return c.JSON(http.StatusBadRequest, Message{
			Msg: err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, Message{
			Msg: val,
		})
	}

}
