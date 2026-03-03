package handlers

import (
	"context"
	"net/http"
	"time"

	"mini-spotify/internal/models"
	"mini-spotify/pkg/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Hàm thêm bài hát mới vào MongoDB
func AddSong(c *gin.Context) {
	var newSong models.Song

	// 1. Kiểm tra dữ liệu gửi lên
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu bài hát không hợp lệ!"})
		return
	}

	// 2. Trỏ tới bảng (collection) "songs" trong MongoDB
	collection := database.MongoDB.Collection("songs")

	// 3. Thực hiện lưu vào DB (Giới hạn thời gian xử lý là 5 giây)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, newSong)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lưu bài hát vào MongoDB"})
		return
	}

	// 4. Báo thành công
	c.JSON(http.StatusOK, gin.H{
		"message": "🎵 Đã thêm bài hát thành công!",
		"song_id": result.InsertedID,
	})
}
func GetSongs(c *gin.Context) {
	collection := database.MongoDB.Collection("songs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find với điều kiện rỗng (bson.M{}) nghĩa là lấy ra TẤT CẢ
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi tìm kiếm bài hát"})
		return
	}
	defer cursor.Close(ctx) // Nhớ đóng kết nối sau khi đọc xong

	var songs []models.Song
	// Lặp qua từng dòng dữ liệu trong MongoDB và đẩy vào mảng 'songs'
	if err = cursor.All(ctx, &songs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi đọc dữ liệu"})
		return
	}

	// Trả mảng bài hát về cho người dùng
	c.JSON(http.StatusOK, gin.H{
		"message": "🎵 Lấy danh sách bài hát thành công!",
		"total":   len(songs), // Báo xem có bao nhiêu bài
		"data":    songs,
	})
}