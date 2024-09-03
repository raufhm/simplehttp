package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type NumbersRequest struct {
	Numbers []int32 `json:"numbers"`
}

func sumNumbers(numbers []int32) int32 {
	var sum int32 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the endpoint
	e.POST("/sum", func(c echo.Context) error {
		req := new(NumbersRequest)
		if err := c.Bind(req); err != nil {
			return c.String(http.StatusBadRequest, "Invalid request")
		}
		sum := sumNumbers(req.Numbers)
		return c.JSON(http.StatusOK, map[string]int32{"result": sum})
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
