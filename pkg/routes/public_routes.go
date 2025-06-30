package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/physicist2018/lidar-web-app/internal/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/user/:id", controllers.GetUser) // get list of all books
	//route.Get("/book/:id", controllers.GetBook)            // get one book by ID
	//route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
