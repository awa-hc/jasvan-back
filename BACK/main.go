package main

import (
	"jasvan/config/initializers"
	"jasvan/internal/delivery/handlers"
	"jasvan/internal/domain/services"
	"jasvan/internal/repository/auth"
	"jasvan/internal/repository/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnvVariables()
	db, err := initializers.InitDB()
	if err != nil {
		panic("Error connecting to database")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Status(200)
	})

	UserRepository := user.NewUserrepository(db)
	userService := services.NewUserService(UserRepository)
	usersHandler := handlers.NewUserHandler(*userService)

	router.POST("/users", usersHandler.CreateUser)
	router.GET("/users/:email", usersHandler.GetUserByEmail)
	router.GET("/users/:username", usersHandler.GetUserByUsername)
	router.GET("/users", usersHandler.GetAllUser)

	AuthRepository := auth.NewAuthRepository(db)
	authService := services.NewAuthService(AuthRepository)
	authHandler := handlers.NewAuthHandler(*authService)

	Auth := router.Group("/auth")
	{
		Auth.POST("/login", authHandler.LoginWithEmail)
	}

	if err := router.Run(":8080"); err != nil {
		panic("Error running server")
	}

}
