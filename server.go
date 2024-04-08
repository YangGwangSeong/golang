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
	e.PATCH("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))

}

func saveUser(c echo.Context) error {
	return c.String(http.StatusOK, "save User")
}

func getUser(c echo.Context) error {
	return c.String(http.StatusOK, "get User")
}

func updateUser(c echo.Context) error {
	return c.String(http.StatusOK, "update User")
}

func deleteUser(c echo.Context) error {
	return c.String(http.StatusOK, "delete User")
}

