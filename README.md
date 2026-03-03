# 🎧 Mini-Spotify Backend API

Dự án Backend cho ứng dụng nghe nhạc cá nhân, được xây dựng để thể hiện khả năng thiết kế hệ thống với kiến trúc sạch (Clean Architecture) và xử lý đa cơ sở dữ liệu.

## 🚀 Các tính năng chính
- **Xác thực & Bảo mật:**
  - Đăng ký và Đăng nhập người dùng.
  - Băm mật khẩu bảo mật bằng thuật toán `bcrypt`.
  - Phân quyền và bảo vệ API bằng **JWT (JSON Web Token)** qua Middleware.
- **Quản lý âm nhạc:**
  - Thêm mới bài hát và lấy danh sách bài hát từ hệ thống.
  - Tạo Playlist cá nhân và lưu trữ danh sách ID bài hát.
- **Thư viện cá nhân:** Lấy danh sách Playlist dựa trên thông tin người dùng được trích xuất từ JWT.

## 🛠 Công nghệ sử dụng (Tech Stack)
- **Ngôn ngữ:** GoLang (Go)
- **Framework:** Gin Gonic (Web Framework hiệu năng cao)
- **Cơ sở dữ liệu (Polyglot Persistence):**
  - **SQL Server:** Quản lý dữ liệu người dùng (Đảm bảo tính nhất quán).
  - **MongoDB:** Lưu trữ Metadata bài hát và Playlist (Đảm bảo tính linh hoạt).
- **Công cụ test:** Thunder Client (VS Code).

## 📁 Cấu trúc thư mục (Clean Architecture)
```text
├── cmd/             # Điểm khởi đầu (Entry point) của ứng dụng
├── internal/
│   ├── handlers/    # Xử lý logic nghiệp vụ cho các API
│   ├── models/      # Định nghĩa cấu trúc dữ liệu (BSON/JSON/SQL)
│   ├── middlewares/ # Bộ gác cổng bảo mật (JWT Auth)
├── pkg/
│   ├── database/    # Cấu hình kết nối SQL Server và MongoDB
├── .gitignore       # Khai báo các file không đưa lên GitHub
└── database.sql     # Script tạo bảng cho SQL Server


## Hướng dẫn cài đặt và khởi chạy
Clone dự án:

Bash
git clone [https://github.com/quanghieu1605/mini-spotify-backend.git](https://github.com/quanghieu1605/mini-spotify-backend.git)
Thiết lập Database:

Chạy các lệnh trong file database.sql trên SQL Server.

Đảm bảo MongoDB đang chạy ở cổng mặc định.

Chạy ứng dụng:

Bash
go run cmd/main.go