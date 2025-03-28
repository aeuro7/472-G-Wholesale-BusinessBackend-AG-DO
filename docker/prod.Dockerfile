FROM golang:1.22 AS builder
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum
COPY go.mod go.sum ./
RUN go mod download

# คัดลอกโค้ดแอปพลิเคชัน
COPY . .

# ขั้นตอนการสร้างภาพสุดท้าย
FROM golang:1.22-alpine
WORKDIR /root/

# คัดลอกโค้ดแอปพลิเคชันจาก builder
COPY --from=builder /app .

# รันแอปพลิเคชันด้วย go run main.go
CMD ["go", "run", "./cmd/main.go"]