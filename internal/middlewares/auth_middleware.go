package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("chuoi_bi_mat_sieu_cap_cua_mini_spotify")

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bạn chưa đăng nhập! (Thiếu Token)"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		// 👉 ĐOẠN NÂNG CẤP: Đọc thông tin từ vé và nhét vào túi (Context) cho các hàm sau dùng
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Ép kiểu float64 (chuẩn của JWT) về int và lưu lại với tên "user_id"
			c.Set("user_id", int(claims["user_id"].(float64)))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Vé thông hành không hợp lệ hoặc đã hết hạn!", "details": err})
			return
		}
	}
}