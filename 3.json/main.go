package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func addUser(c echo.Context) error {
	user := User{}

	b, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()

	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	log.Printf("we got a user: %#v", user)
	return c.String(http.StatusOK, fmt.Sprintf("This is your user"))
}

func main() {
	e := echo.New()

	e.POST("/adduser", addUser)

	e.Start(":8000")
}
