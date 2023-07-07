package controllers

import (
	"errors"
	"fmt"
)

func ControllerHelloWorld(nombre string) (string, error) {
	if nombre == "" || nombre == "ERROR" {
		return nombre, errors.New("error en el nombre")
	}
	rtn := fmt.Sprintf("Hola %s!", nombre)
	return rtn, nil

}
