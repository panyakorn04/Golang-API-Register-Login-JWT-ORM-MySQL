package main

import (
	"log"

	AuthController "example.com/greetings/controller/auth"
	UserController "example.com/greetings/controller/user"
	"example.com/greetings/middleware"
	"example.com/greetings/orm"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	orm.ConnectDatabase()

	r := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// CORS middleware
	r.Use(cors.Default())

	// Routes
	r.POST("/register", AuthController.Register)
	r.POST("/login", AuthController.Login)
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/users", middleware.AuthorizeJWT())
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.GET("/", UserController.GetUsers)
	authorized.GET("/:id", UserController.GetUser)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
