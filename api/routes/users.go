package routes

import (
	"api/handlers"

	"github.com/labstack/echo/v4"
)

func Users_Route(e *echo.Echo) {
	// Register Route
	e.POST("/api/register", handlers.Register)
	// Login Route
	e.POST("/api/login", handlers.Login)
	// // Get User Route
	// e.GET("/api/users/:id", handlers.Fetch_User)
	// // Update User Route
	// e.PUT("/api/users/:id", handlers.Update_User)
	// // Delete User Route
	// e.DELETE("/api/users/:id", handlers.Delete_User)
}
