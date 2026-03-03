package handlers

import (
	"net/http"
	"database/sql"
	"time"

	"mini-spotify/internal/models"  // Nhớ sửa "mini-spotify" thành tên thư mục dự án của em nếu khác
	"mini-spotify/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
var jwtSecretKey = []byte("chuoi_bi_mat_sieu_cap_cua_mini_spotify")

// Hàm xử lý việc đăng ký
func Register(c *gin.Context) {
	var req models.RegisterRequest

	// 1. Nhận và kiểm tra dữ liệu từ người dùng gửi lên
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ, vui lòng kiểm tra lại!", "details": err.Error()})
		return
	}

	// 2. Băm mật khẩu (Hash Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi hệ thống khi mã hóa mật khẩu"})
		return
	}

	// 3. Lưu thông tin vào SQL Server
	// Dùng @p1, @p2, @p3 để chống lỗi bảo mật SQL Injection
	query := `INSERT INTO Users (Username, Email, PasswordHash) VALUES (@p1, @p2, @p3)`
	
	_, err = database.SQL.Exec(query, req.Username, req.Email, string(hashedPassword))
	if err != nil {
		// Thường lỗi ở đây là do Email bị trùng (vì lúc thiết kế DB mình đặt Email là UNIQUE)
		c.JSON(http.StatusConflict, gin.H{"error": "Email này đã được sử dụng hoặc có lỗi xảy ra!"})
		return
	}

	// 4. Báo thành công cho Front-end
	c.JSON(http.StatusOK, gin.H{
		"message": "🎉 Đăng ký tài khoản thành công!",
		"user": req.Username,
	})
}
func Login(c *gin.Context) {
	var req models.LoginRequest

	// 1. Nhận dữ liệu từ Front-end
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ!"})
		return
	}

	// 2. Tìm User trong SQL Server dựa vào Email
	var user struct {
		ID           int
		Username     string
		PasswordHash string
	}
	
	query := `SELECT UserID, Username, PasswordHash FROM Users WHERE Email = @p1`
	err := database.SQL.QueryRow(query, req.Email).Scan(&user.ID, &user.Username, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email không tồn tại trên hệ thống!"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi truy xuất Database"})
		}
		return
	}

	// 3. Đọ sức mật khẩu (So sánh cái người dùng gõ với cái bị băm trong DB)
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Mật khẩu không chính xác!"})
		return
	}

	// 4. Tạo vé JWT (JSON Web Token)
	// Vé này sẽ chứa ID và Tên của user, và có hạn sử dụng là 24 giờ
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
	})

	// Ký tên lên vé bằng chìa khóa bí mật
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo vé thông hành"})
		return
	}

	// 5. Trả vé về cho người dùng cất vào trình duyệt
	c.JSON(http.StatusOK, gin.H{
		"message": "🎉 Đăng nhập thành công!",
		"token":   tokenString,
	})
}