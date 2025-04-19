# Gunakan image Go resmi untuk membangun aplikasi
FROM golang:1.24-alpine AS builder

# Set environment variable
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org

# Set direktori kerja di dalam container
WORKDIR /app

# Salin go.mod dan go.sum
COPY go.mod go.sum ./

# Install dependencies yang dibutuhkan
RUN go mod tidy

# Salin seluruh kode sumber ke dalam container
COPY . .

# Build aplikasi Go (ganti dengan nama binary yang sesuai)
RUN go build -o sirnawa_app cmd/main.go

# Gunakan image yang lebih kecil untuk menjalankan aplikasi
FROM alpine:latest

# Install dependencies runtime (misalnya, libc untuk menjalankan Go binary)
RUN apk --no-cache add ca-certificates

# Set direktori kerja
WORKDIR /root/

# Salin binary yang sudah dibuild dari stage sebelumnya
COPY --from=builder /app/sirnawa_app .

# Salin file .env ke dalam container
COPY --from=builder /app/.env /root/.env

# Expose port yang akan digunakan oleh aplikasi (sesuaikan dengan aplikasi Anda)
EXPOSE 5050

# Perintah untuk menjalankan aplikasi saat container berjalan
CMD ["./sirnawa_app"]
