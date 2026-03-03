package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Song đại diện cho cấu trúc một bài hát
type Song struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title    string             `bson:"title" json:"title" binding:"required"`
	Artist   string             `bson:"artist" json:"artist" binding:"required"`
	AudioURL string             `bson:"audio_url" json:"audio_url" binding:"required"`
}