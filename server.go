package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	e.GET("/handle", func(c echo.Context) error {
		cc := c.(*CustomContext)
		cc.Foo()
		cc.Bar()
		return cc.String(200, "OK")
	})

	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /middleware/users")
			return next(c)
		}
	}

	// Struct Tag Binding
	// bind/users?id=<userID> 이렇게 설정 했었는데 not found 에러로 발생했다.
	// 라우터 설정을 bind/users 이렇게 해두고 bind/users?id=1234 이런식으로 요청하면 id값이 정상적으로 받을 수 있다!
	// Struct Tag Binding이란걸 잘 활용하면 Struct에 파라미터들을 선언하고 유효성 검사를 할 수 있을것 같다.
	// 근데 혹시 body도 가능한가? 가능한것 같다.
	// 시도 해볼것 bind/users?id=1234/1234
	e.GET("bind/users",getBindUser)

	// Route level middleware
	e.GET("/middleware/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// users
	e.POST("/users", saveUser)
	e.POST("/save", save)
	e.POST("/multi/save", multiSave)
	// Handling Request
	e.POST("/handle/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		return c.JSON(http.StatusCreated,u)
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.PATCH("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	// image,  pdf , fonts ...
	e.Static("/static","static!!");
	e.Logger.Fatal(e.Start(":1323"))

}

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

// query 파라미터는 Struct Tag Binding으로 받는것 같다
type BindUSer struct{
	ID string `query:"id"`
}

type User struct { // handle/users로 form데이터 전송과 raw json 전송은 정상적으로 받아지는데 여기서 query는 쿼리 파라미터가 아닌듯
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
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

func getBindUser(c echo.Context) error {
	var user BindUSer
	err := c.Bind(&user); if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	return c.String(http.StatusOK, user.ID)
}

