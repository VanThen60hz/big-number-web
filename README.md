# Big Number Web Application

Ứng dụng web cho phép thực hiện phép cộng các số lớn và hiển thị chi tiết các bước tính toán.

## Tính năng

- Cộng hai số nguyên lớn không giới hạn độ dài
- Hiển thị chi tiết từng bước trong quá trình tính toán
- Giao diện web thân thiện, dễ sử dụng

## Yêu cầu hệ thống

- Go 1.22.3 hoặc cao hơn
- Các dependencies được quản lý bởi Go modules

## Cài đặt

1. Clone repository:

   ```bash
   git clone https://github.com/VanThen60hz/big-number-web.git
   cd big-number-web
   ```

2. Cài đặt dependencies:

   ```bash
   go mod tidy
   ```

3. Chạy ứng dụng:

   ```bash
   go run main.go
   ```

Ứng dụng sẽ chạy tại địa chỉ: http://localhost:8080

## Cách sử dụng

1. Truy cập vào http://localhost:8080
2. Nhập số thứ nhất vào ô "Number 1"
3. Nhập số thứ hai vào ô "Number 2"
4. Nhấn nút "Add Numbers"
5. Kết quả sẽ hiển thị cùng với các bước tính toán chi tiết

## Cấu trúc project

```
big-number-web/
├── main.go              # Entry point của ứng dụng
├── templates/           # Chứa các template HTML
│   └── index.html      # Trang chủ
├── static/             # Chứa các file tĩnh (CSS, JS)
│   └── style.css      # File CSS
├── go.mod             # File quản lý dependencies
└── README.md          # File hướng dẫn
```

## Dependencies chính

- [big-number-addition](https://github.com/VanThen60hz/big-number-addition) - Package xử lý phép cộng số lớn
- [gin-gonic/gin](https://github.com/gin-gonic/gin) - Web framework
- [sirupsen/logrus](https://github.com/sirupsen/logrus) - Logging framework
