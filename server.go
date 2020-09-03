package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	r.Run(":" + getEnv("PORT", Port))
}

func setupRouter() *gin.Engine {
	r := gin.New()

	environment := getEnv("ENV", Environment)

	log.Println("ENV = " + environment)

	if environment != "local" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/health"),
		gin.Recovery(),
	)

	type Db *sql.DB
	db := connectDb()

	env := &Env{db: db}

	api := r.Group("/api")
	{
		api.GET("/health", checkHealth)
		api.POST("/auth", env.authen)
		api.POST("/get-txnType", TokenAuthMiddleware(), getHardCodeDropdown)
		api.POST("/get-hFrom", TokenAuthMiddleware(), getHardCodeDropdown)
		api.POST("/get-mFrom", TokenAuthMiddleware(), getHardCodeDropdown)
		api.POST("/get-sFrom", TokenAuthMiddleware(), getHardCodeDropdown)
	}

	return r
}
