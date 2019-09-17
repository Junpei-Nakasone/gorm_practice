package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	initMigration()
	handleRequest()
}

func initMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	checkConnectError(err)
	defer db.Close()
	db.AutoMigrate(&User{})
}

func checkConnectError(err error) {
	if err != nil {
		panic("failed to connect to DBtest.")
	}
}

func handleRequest() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", allUsers)
	e.DELETE("/user/:name", deleteUser)
	e.PUT("/user/:name/:email", updateUser)
	e.POST("/user/:name/:email", newUser)

	fmt.Println("server listening")
	log.Fatal(e.Start(":8000"))
}

func allUsers(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "test.db")
	checkConnectError(err)
	defer db.Close()

	var users []User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func newUser(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "test.db")
	checkConnectError(err)
	defer db.Close()

	id := c.Param("id")
	email := c.Param("email")

	db.Create(&User{Name: id, Email: email})
	return c.NoContent(http.StatusOK)
}

func updateUser(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "test.db")
	checkConnectError(err)
	defer db.Close()

	user := new(User)
	name := c.Param("name")
	email := c.Param("email")

	db.Where("name=?", name).Find(&user)
	user.Email = email
	db.Save(&user)

	return c.NoContent(http.StatusOK)
}

func deleteUser(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "test.db")
	checkConnectError(err)
	defer db.Close()
	var user User

	name := c.Param("name")
	db.Where("name=?", name).Find(&user)
	db.Delete(&user)

	return c.JSON(http.StatusOK, "ok")
}
