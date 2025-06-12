package main

import (
	"fmt"
	"net/http"

	"github.com/Chayakorn2002/thb-amount-to-text-go/dto"
	"github.com/Chayakorn2002/thb-amount-to-text-go/services"
	"github.com/cnc-csku/task-nexus-go-lib/jsonvalidator"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Validator = jsonvalidator.NewValidator()

	service := services.NewConverterService()

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &echo.Map{
			"message": "pong",
		})
	})

	e.GET("/convert", func(c echo.Context) error {
		req := new(dto.ConvertDecimalToBahtTextRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, &echo.Map{
				"message": "Invalid request format",
			})
		}

		if err := c.Validate(req); err != nil {
			return c.JSON(http.StatusBadRequest, &echo.Map{
				"message": "Validation failed",
				"error":   err.Error(),
			})
		}

		resp, err := service.ConvertDecimalToBahtText(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &echo.Map{
				"message": "Failed to convert amount",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, resp)
	})

	for _, route := range e.Routes() {
		fmt.Printf("Method: %s, Path: %s, Name: %s\n", route.Method, route.Path, route.Name)
	}

	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal("Failed to start server:", err)
	}
}
