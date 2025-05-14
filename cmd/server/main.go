package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"side-project-unnamed-3/backend/internal/config"
	"side-project-unnamed-3/backend/internal/handlers"
	"side-project-unnamed-3/backend/internal/models"
)

func main() {
	// 載入 .env 檔案
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	if err := config.ConnectDatabase(); err != nil {
		log.Fatalf("資料庫連線失敗: %v", err)
	}

	// 自動 migrate User
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("自動 migrate 失敗: %v", err)
	}

	r := gin.Default()

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("伺服器啟動於 :%s", port)
	r.Run(":" + port)
}
