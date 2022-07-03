package hub

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartHub() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/blockholders", func(c echo.Context) error {
		services, err := GetServiceDescription("blockholder")
		if err != nil {
			return c.JSON(http.StatusBadRequest, apiResponse{
				Message: err.Error(),
			})
		}
		if len(services) == 0 {
			return c.JSON(http.StatusNotFound, services)
		}
		return c.JSON(http.StatusOK, services)
	})
	e.GET("/miners", func(c echo.Context) error {
		services, err := GetServiceDescription("miner")
		if err != nil {
			return c.JSON(http.StatusBadRequest, apiResponse{
				Message: err.Error(),
			})
		}
		if len(services) == 0 {
			return c.JSON(http.StatusNotFound, services)
		}
		return c.JSON(http.StatusOK, services)
	})
	e.GET("/hubs", func(c echo.Context) error {
		services, err := GetServiceDescription("hub")
		if err != nil {
			return c.JSON(http.StatusBadRequest, apiResponse{
				Message: err.Error(),
			})
		}
		if len(services) == 0 {
			return c.JSON(http.StatusNotFound, services)
		}
		return c.JSON(http.StatusOK, services)
	})
	e.POST("/blockholder", func(c echo.Context) error {
		serviceDesc := new(serviceDescription)
		c.Bind(serviceDesc)
		serviceDesc.Type = "blockholder"
		if err := SaveServiceDescription(*serviceDesc); err != nil {
			return c.JSON(http.StatusBadRequest, apiResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, apiResponse{
			Message: "blockholder created",
		})
	})
	e.POST("/miner", func(c echo.Context) error {
		serviceDesc := new(serviceDescription)
		c.Bind(serviceDesc)
		serviceDesc.Type = "miner"
		SaveServiceDescription(*serviceDesc)
		return c.JSON(http.StatusOK, apiResponse{
			Message: "miner created",
		})
	})
	e.POST("/dispatcher", func(c echo.Context) error {
		serviceDesc := new(serviceDescription)
		c.Bind(serviceDesc)
		serviceDesc.Type = "dispatcher*"
		SaveServiceDescription(*serviceDesc)
		return c.JSON(http.StatusOK, apiResponse{
			Message: "dispatcher created",
		})
	})
	e.POST("/hubs", func(c echo.Context) error {
		serviceDesc := new(serviceDescription)
		c.Bind(serviceDesc)
		serviceDesc.Type = "hub"
		SaveServiceDescription(*serviceDesc)
		return c.JSON(http.StatusOK, apiResponse{
			Message: "hub created",
		})
	})
	e.DELETE("/blockholder", func(c echo.Context) error {
		serviceDesc := new(serviceDescription)
		c.Bind(serviceDesc)
		if err := DeleteServiceDescription("blockholder", serviceDesc.DNS); err != nil {
			return c.JSON(http.StatusBadRequest, apiResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, apiResponse{
			Message: "blockholder deleted",
		})
	})
	e.DELETE("/miner", func(c echo.Context) error {
		serviceDesc := new(serviceDescription)
		c.Bind(serviceDesc)
		if err := DeleteServiceDescription("miner", serviceDesc.DNS); err != nil {
			return c.JSON(http.StatusBadRequest, apiResponse{
				Message: err.Error(),
			})
		}
		return c.JSON(http.StatusOK, apiResponse{
			Message: "miner deleted",
		})
	})
	e.Logger.Fatal(e.Start(":8000"))
}
