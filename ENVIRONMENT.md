## JC_FIREBASE_SERVICE_ACCOUNT

[Docs](https://firebase.google.com/docs/admin/setup?authuser=0#go)

Tạo ứng dụng firebase

Thêm provider `Email/Password` và `Google`
![provider](https://i.imgur.com/lxCAzKX.png)
Cài đặt chức năng `Storage` ở chế độ `test`
![Imgur](https://i.imgur.com/y69th6l.png)
![Imgur](https://i.imgur.com/qIXHB9w.png)
Lấy private key của ứng dụng
![Imgur](https://i.imgur.com/LzG14RX.png)
Xóa kí tự xuống dòng rồi copy nội dung vào file `.env`
![Imgur](https://i.imgur.com/wBbL3V5.png)

## DATABASE_URL

[Cài đặt Docker](https://docs.docker.com/get-started/)

Chạy docker container được cấu hình trong `docker-compose.yml`

```bash
docker compose up -d
```

Cập nhật giá trị của biến môi trường

```bash
DATABASE_URL = postgres://postgres:postgres@localhost:5432/postgres
```
