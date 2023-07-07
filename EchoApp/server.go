package main

import (
	"TestGO/EchoApp/routers"
	_ "TestGO/docs/EchoApp"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title API ECHO de prueba
// @version 1.0
// @description Este es un ejemplo sencillo utilizando Echo para crear una api.
// @host localhost:3333
// @BasePath /
// @schemes http

type Message struct {
	Msg string `json:"message"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", Welcome)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.GET("/usopathparams/:value", UsoParametros)
	e.GET("/usoqueryparams", UsoQuery)
	// Esta estructura debería reemplazar lo anterior
	routers.InitTestRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":3333"))
}

// Welcome godoc
// @Summary Muestra mensaje de bienvenida.
// @Description Endpoint básico que muestra mensaje de bienvenida en el servidor.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func Welcome(c echo.Context) error {
	return c.JSON(http.StatusOK, Message{
		Msg: "Bienvenido",
	})
}

// UsoParametros godoc
// @Summary Muestra valor de variable especificada.
// @Description Endpoint que muestra el valor de la variable especificada por path.
// @Param value path string true "Valor que se mostrará en la respuesta"
// @Tags UsoParametros
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /usopathparams/{value} [get]
func UsoParametros(c echo.Context) error {
	params := c.Param("value")
	return c.JSON(http.StatusOK, Message{
		Msg: "El valor entregado por Path Params a 'value' es: " + params,
	})
}

// UsoQuery godoc
// @Summary Muestra valor de la variable value especificada.
// @Description Endpoint que muestra el valor de la variable 'value' especificada por query.
// @Param value query string true "Valor que se mostrará en la respuesta"
// @Tags UsoQuery
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /usoqueryparams [get]
func UsoQuery(c echo.Context) error {
	query := c.QueryParam("value")
	return c.JSON(http.StatusOK, Message{
		Msg: "El valor entregado por Query Params a 'value' es: " + query,
	})
}
