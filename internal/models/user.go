package models

// RegisterRequest là cấu trúc để nhận dữ liệu JSON từ Front-end gửi lên
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`       // Bắt buộc phải có
	Email    string `json:"email" binding:"required,email"`    // Bắt buộc, và phải đúng định dạng email
	Password string `json:"password" binding:"required,min=6"` // Bắt buộc, độ dài tối thiểu 6 ký tự
}

// LoginRequest là cấu trúc để nhận dữ liệu Đăng nhập
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
