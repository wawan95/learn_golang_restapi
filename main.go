package main

import (
	"log"
	"os"
	"path/filepath"

	"golang-user-api/config"
	"golang-user-api/controllers"
	"golang-user-api/models"
	"golang-user-api/repositories"
	"golang-user-api/routes"
	"golang-user-api/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// 1. Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	// 2. Connect MySQL
	config.ConnectDatabase()

	// 3. Auto Migration
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// 4. Dependency Injection
	userRepo := repositories.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// 5. Initialize Gin
	router := gin.Default()

	// 6. Register API Routes
	routes.SetupRoutes(router, userController)

	// 7. Get Working Directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Cannot get working directory:", err)
	}

	log.Println("Working Directory:", workingDir)

	// 8. Define Swagger Files
	swaggerHTML := filepath.Join(workingDir, "docs", "index.html")
	openAPIFile := filepath.Join(workingDir, "docs", "openapi.yaml")

	// 9. Check Swagger Files
	for _, file := range []string{
		swaggerHTML,
		openAPIFile,
	} {
		info, err := os.Stat(file)

		if err != nil {
			log.Println("FILE NOT FOUND:", file, err)
			continue
		}

		if info.IsDir() {
			log.Println("ERROR: Expected file but found directory:", file)
			continue
		}

		log.Println("FILE FOUND:", file)
	}

	// 10. Register Swagger UI Route
	// Register Swagger UI Route
	router.GET("/swagger/index.html", func(c *gin.Context) {
		c.File(swaggerHTML)
	})

	// Support /swagger/
	router.GET("/swagger/", func(c *gin.Context) {
		c.File(swaggerHTML)
	})

	// Support /swagger
	router.GET("/swagger", func(c *gin.Context) {
		c.File(swaggerHTML)
	})

	// 11. Register OpenAPI YAML Route
	router.GET("/openapi.yaml", func(c *gin.Context) {
		c.File(openAPIFile)
	})

	// 12. Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "API is running",
		})
	})

	// 13. Test Route
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "GIN ROUTE WORKING!")
	})

	// 14. Debug Docs Folder
	files, err := os.ReadDir(filepath.Join(workingDir, "docs"))

	if err != nil {
		log.Println("Cannot read docs folder:", err)
	} else {
		for _, file := range files {
			log.Println("DOCS FILE:", file.Name())
		}
	}

	// 15. Debug Registered Routes
	log.Println("========== REGISTERED ROUTES ==========")

	for _, route := range router.Routes() {
		log.Printf(
			"METHOD: %s | PATH: %s",
			route.Method,
			route.Path,
		)
	}

	// 16. Start Server
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("========================================")
	log.Println("Server running on port:", port)
	log.Println("Health Check: http://localhost:" + port + "/health")
	log.Println("Swagger UI: http://localhost:" + port + "/swagger/index.html")
	log.Println("OpenAPI YAML: http://localhost:" + port + "/openapi.yaml")
	log.Println("========================================")

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Server failed:", err)
	}
}
