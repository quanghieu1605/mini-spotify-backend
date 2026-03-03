package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb" 
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	SQL     *sql.DB
	MongoDB *mongo.Database
)

func ConnectDB() {
	// 1. KẾT NỐI SQL SERVER (Dùng Windows Authentication)
	// 💡 Senior Note: Dùng dấu backtick (`) thay vì ngoặc kép (") để GoLang không bị lỗi dấu gạch chéo (\) trong tên Server.
	// Bật trusted_connection=yes để dùng quyền Windows.
	sqlConnString := `server=LAPTOP-8U8H70M9\SQLEXPRESS;database=MiniSpotify;trusted_connection=yes;encrypt=true;TrustServerCertificate=true`
	
	var err error
	SQL, err = sql.Open("sqlserver", sqlConnString)
	if err != nil {
		log.Fatal(" Lỗi cấu hình SQL Server: ", err)
	}
	if err := SQL.Ping(); err != nil {
		log.Fatal(" Không thể chạm tới SQL Server: ", err)
	}
	fmt.Println("Đã kết nối thành công tới SQL Server (Bảng Users)!")

	// 2. KẾT NỐI MONGODB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := "mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(" Lỗi cấu hình MongoDB: ", err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(" Không thể chạm tới MongoDB: ", err)
	}
	
	MongoDB = client.Database("mini_spotify_db")
	fmt.Println(" Đã kết nối thành công tới MongoDB (Bảng Songs & Playlists)!")
}