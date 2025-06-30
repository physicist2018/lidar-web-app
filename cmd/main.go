package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"github.com/physicist2018/lidar-web-app/pkg/configs"
	"github.com/physicist2018/lidar-web-app/pkg/middleware"
	"github.com/physicist2018/lidar-web-app/pkg/routes"
	"github.com/physicist2018/lidar-web-app/pkg/utils"
)

func main() {
	// TODO: Implement your application logic here
	config := configs.FiberConfig()

	app := fiber.New(config)
	middleware.FiberMiddleware(app)
	//routes.SwaggerRoute(app)
	routes.PublicRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}
