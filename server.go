package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	Msg string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/", Welcome)
	e.GET("/usopathparams/:value", UsoParametros)
	e.GET("/usoqueryparams", UsoQuery)
	e.Logger.Fatal(e.Start(":3000"))
}

func Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, Message{
		Msg: "Bienvenido",
	})
}

func UsoParametros(c echo.Context) error {
	params := c.Param("value")
	return c.JSON(http.StatusOK, Message{
		Msg: "El valor entregado por Path Params a 'value' es: " + params,
	})
}

func UsoQuery(c echo.Context) error {
	query := c.QueryParam("value")
	return c.JSON(http.StatusOK, Message{
		Msg: "El valor entregado por Query Params a 'value' es: " + query,
	})
}
