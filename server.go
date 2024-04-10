package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// users
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.PATCH("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))

}

type person struct {
	ID string `json:"id"`
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

