package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func user(c echo.Context) error {
	Name := c.QueryParam("name")
	Type := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("user name is %s\n type is %s", Name, Type))
	}
	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": Name,
			"type": Type,
		})
	}
	return c.String(http.StatusBadRequest, "")
}

func main() {
	e := echo.New()

	e.GET("/user/:data", user)

	e.Start(":8000")
}
