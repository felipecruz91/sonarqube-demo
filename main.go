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

func doNothing() { // Noncompliant
}

func run() {
	fmt.Println("This should be a constant") // Noncompliant; 'This should ...' is duplicated 3 times
	fmt.Println("This should be a constant")
	fmt.Println("This should be a constant")
}

func setCoordinates(x1 int, y1 int, z1 int, x2 int, y2 int, z2 int) { // Noncompliant
	// ...
}

func nonExistentOperators() {
	var target, num = -5, 3

	target = -num // Noncompliant: target = -3. Is that the expected behavior?
	target = +num // Noncompliant: target = 3
	fmt.Println(target)
}
