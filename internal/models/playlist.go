package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Playlist là cấu trúc lưu vào MongoDB
type Playlist struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	UserID  int                  `bson:"user_id" json:"user_id"` // ID của user lấy từ SQL Server
	Name    string               `bson:"name" json:"name"`
	SongIDs []primitive.ObjectID `bson:"song_ids" json:"song_ids"` // Danh sách các ID bài hát
}

// CreatePlaylistRequest là cấu trúc để hứng dữ liệu từ Front-end gửi lên
type CreatePlaylistRequest struct {
	Name    string   `json:"name" binding:"required"`
	SongIDs []string `json:"song_ids"` // Front-end sẽ gửi mảng các chuỗi ID
}