# Facebook Creative API - Go Project

Dự án Go đơn giản để tạo, danh sách, và xóa creative ads trên Facebook.

## Yêu cầu

- Go 1.25.5 trở lên
- Facebook Access Token
- Ad Account ID

## Cài đặt

### 1. Clone/Setup dự án
```bash
cd d:\creative_fb
```

### 2. Tải dependencies
```bash
go mod download
go get github.com/joho/godotenv
```

### 3. Cấu hình Access Token

Tạo file `.env` từ template:
```bash
copy .env.example .env
```

Sau đó chỉnh sửa `.env` và thêm Access Token của bạn:
```
FACEBOOK_ACCESS_TOKEN=your_actual_token_here
```

## Cách sử dụng

### Tạo Single Image Creative

```bash
go run ./cmd/api -cmd create \
  -account your_account_id \
  -page your_page_id \
  -name "My Single Image Creative" \
  -image "https://example.com/image.jpg" \
  -link "https://example.com"
```

Với message:
```bash
go run ./cmd/api -cmd create \
  -account your_account_id \
  -page your_page_id \
  -name "My Creative with Message" \
  -image "https://example.com/image.jpg" \
  -link "https://example.com" \
  -message "Check this out!"
```

### Danh sách Creatives

```bash
go run ./cmd/api -cmd list -account your_account_id
```

### Xóa Creative

```bash
go run ./cmd/api -cmd delete -account your_account_id -name "My Creative"
```

## Cấu trúc dự án

```
creative_fb/
├── cmd/
│   └── api/
│       └── main.go           # Điểm vào chính
├── internal/
│   └── facebook/
│       └── client.go         # Facebook API client
├── .env.example              # Template biến môi trường
├── go.mod                    # Module dependencies
└── README.md                 # Documentation
```

## API Reference

### CreateCreative
Tạo một creative single image mới cho campaign ads.

```go
linkData := facebook.LinkData{
    ImageURL: "https://example.com/image.jpg",
    Link:     "https://example.com",
    Message:  "Check this out!",
}

creative := facebook.CreateCreativeRequest{
    Name: "My Creative",
    ObjectStorySpec: facebook.ObjectStorySpec{
        PageID:   "123456789",
        LinkData: linkData,
    },
}
result, err := client.CreateCreative(accountID, creative)
```

### ListCreatives
Lấy danh sách tất cả creatives trong ad account.

```go
creatives, err := client.ListCreatives(accountID)
```

### DeleteCreative
Xóa creative theo tên.

```go
success, err := client.DeleteCreative(accountID, "My Creative")
```

## Lấy Access Token

1. Truy cập: https://developers.facebook.com
2. Tạo một app hoặc sử dụng existing app
3. Vào Tools → Graph API Explorer
4. Chọn permission: `ads_management`
5. Copy access token và lưu vào `.env`

## Ghi chú

- API sử dụng Facebook Graph API v18.0
- Cần quyền `ads_management` để tạo creatives
- Creative phải có hình ảnh hợp lệ
- Đảm bảo Ad Account ID đúng format

## Lỗi phổ biến

- **"access_token is invalid"** → Kiểm tra access token trong `.env`
- **"Invalid account id"** → Kiểm tra Ad Account ID (thường là số)
- **"Image is not valid"** → URL hình ảnh phải công khai và accessible

## Phát triển tiếp theo

- Thêm support cho video creatives
- Thêm validation input
- Thêm retry logic
- Thêm caching
- Thêm unit tests
