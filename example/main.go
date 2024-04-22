package main

import (
	"fmt"
	"github.com/oxxzz/egg"
)

func main() {
	s := egg.New().Debug()

	s.GET("/", func(c *egg.Context) error {
		return c.JSON(200, egg.M{"message": "Hello, World!"})
	})

	s.GET("/hello/:name/fast/:id", func(c *egg.Context) error {
		return c.JSON(200, egg.M{"message": "Hello, " + c.PathValue("id") + " " + c.PathValue("name")})
	})

	s.GET("/assets/*filepath", func(c *egg.Context) error {
		return c.JSON(200, egg.M{"message": "Hello, " + c.PathValue("filepath")})
	})

	fmt.Println("application running on http://localhost:8001")
	fmt.Println(s.Run(":8001"))
}
