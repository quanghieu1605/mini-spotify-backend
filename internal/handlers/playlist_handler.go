package handlers

import (
	"context"
	"net/http"
	"time"

	"mini-spotify/internal/models"
	"mini-spotify/pkg/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePlaylist(c *gin.Context) {
	// 1. Lấy ID người dùng từ anh Bảo vệ
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Không nhận diện được người dùng"})
		return
	}

	// 2. Nhận dữ liệu (Tên Playlist và mảng ID bài hát)
	var req models.CreatePlaylistRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// 3. Biến đổi mảng chuỗi ID bài hát thành chuẩn ObjectID của MongoDB
	var objectIDs []primitive.ObjectID
	for _, idStr := range req.SongIDs {
		objID, err := primitive.ObjectIDFromHex(idStr)
		if err == nil {
			objectIDs = append(objectIDs, objID)
		}
	}

	// 4. Đổ dữ liệu vào khuôn
	newPlaylist := models.Playlist{
		UserID:  userID.(int), // Ép kiểu về int
		Name:    req.Name,
		SongIDs: objectIDs,
	}

	// 5. Lưu vào collection "playlists" trong MongoDB
	collection := database.MongoDB.Collection("playlists")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, newPlaylist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi lưu Playlist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "🎧 Tạo Playlist thành công!",
		"playlist_id": result.InsertedID,
	})
}
func GetMyPlaylists(c *gin.Context) {
	// 1. Lấy ID của mình từ vé (Middleware đã nhét sẵn vào Context lúc nãy)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Không nhận diện được người dùng"})
		return
	}

	// 2. Trỏ tới bảng playlists trong MongoDB
	collection := database.MongoDB.Collection("playlists")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 3. TÌM KIẾM CÓ ĐIỀU KIỆN: Chỉ lấy playlist nào có cột "user_id" trùng khớp với ID của mình
	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi tìm kiếm Playlist"})
		return
	}
	defer cursor.Close(ctx) // Nhớ đóng kết nối sau khi đọc xong

	var playlists []models.Playlist
	if err = cursor.All(ctx, &playlists); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi đọc dữ liệu Playlist"})
		return
	}

	// (Kinh nghiệm đi làm thực tế) Nếu chưa có playlist nào, trả về mảng rỗng [] thay vì giá trị null để Front-end không bị lỗi sập web
	if playlists == nil {
		playlists = []models.Playlist{}
	}

	// 4. Trả kết quả về
	c.JSON(http.StatusOK, gin.H{
		"message": "🎧 Lấy thư viện nhạc thành công!",
		"total":   len(playlists),
		"data":    playlists,
	})
}