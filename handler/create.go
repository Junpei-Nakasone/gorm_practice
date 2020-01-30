package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type User struct {
	Id   int    `gorm:"AUTO_INCREMENT"`
	Name string `gorm:"type:varchar(50)"`
}

func gormConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(docker.for.mac.localhost:3306)/gorm_practice")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateUser(c echo.Context) (err error) {
	db := gormConnect()
	defer db.Close()
	u := &User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	u.Name = "test"
	db.Table("User").Create(&u)
	return c.JSON(http.StatusCreated, u)
}
