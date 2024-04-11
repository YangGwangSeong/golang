package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// users
	e.POST("/users", saveUser)
	e.POST("/save", save)
	e.POST("/multi/save", multiSave)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.PATCH("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))

}

type person struct {
	ID string `json:"id"`
}
func multiSave(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}
 
 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()
 
 	// Destination
 	dst, err := os.Create(avatar.Filename)
 	if err != nil {
 		return err
 	}
 	defer dst.Close()
 
 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}

// Form application/x-www-form-urlencoded
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:" + name + ", email: " + email)
}

func saveUser(c echo.Context) error {
	return c.String(http.StatusOK, "save User")
}

// Query Parameters
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member);
}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	id := c.Param("id")
	p := person{
		ID: id,
	}
	return c.JSON(http.StatusOK,p)
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "update User")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "delete User")
}

