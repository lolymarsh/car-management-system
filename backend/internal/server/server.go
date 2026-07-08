package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/uptrace/bun"

	"github.com/lolymarsh/car-management-system/internal/car"
)

func NewServer(db *bun.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	carRepo := car.NewRepository(db)
	carSvc := car.NewService(carRepo)
	carHdl := car.NewHandler(carSvc)

	api := e.Group("/api")
	api.POST("/cars", carHdl.CreateCar)
	api.POST("/cars/filter", carHdl.ListCars)
	api.GET("/cars/:id", carHdl.GetCar)
	api.PUT("/cars/:id", carHdl.UpdateCar)
	api.DELETE("/cars/:id", carHdl.DeleteCar)

	return e
}
