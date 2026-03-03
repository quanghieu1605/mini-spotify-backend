package main

import (
	"fmt"
	"log"
	"net/http"

	"mini-spotify/internal/handlers"
	"mini-spotify/internal/middlewares"
	"mini-spotify/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 Đang khởi động Mini-Spotify Backend...")

	// 1. Khởi tạo kết nối Database
	database.ConnectDB()

	// 2. Khởi tạo Gin Router (Bộ định tuyến API)
	router := gin.Default()
    // Mở chốt an toàn CORS, cho phép Front-end gọi API thoải mái
	router.Use(cors.Default())

	// 3. Viết API đầu tiên (Health Check) để test xem Server có mở cửa không
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong! Server Mini-Spotify đang chạy rất mượt 🎵",
			"status":  "success",
		})
	})
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.POST("/songs", middlewares.RequireAuth(), handlers.AddSong)
	router.GET("/songs", handlers.GetSongs)
	router.POST("/playlists", middlewares.RequireAuth(), handlers.CreatePlaylist)
	router.GET("/playlists", middlewares.RequireAuth(), handlers.GetMyPlaylists)
	// 4. Chạy server ở cổng 8080
	fmt.Println("🌐 Server đang mở cửa đón khách tại: http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("❌ Lỗi sập server: ", err)
	}
}
