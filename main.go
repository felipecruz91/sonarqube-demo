package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// just produce duplicated code to make the SonarQube quality gate fail
	fmt.Println("hi")
	fmt.Println("hi")
	fmt.Println("hi")
	fmt.Println("hi")
	
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
