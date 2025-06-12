
# Sirnawa Backend

**Sirnawa** adalah aplikasi manajemen warga untuk lingkungan RT, dibangun untuk mempermudah pengelolaan data rumah, penduduk, pembayaran iuran, ronda malam, pengumuman, komunitas, dan marketplace warga.  
Ini adalah bagian backend dari aplikasi, dibangun menggunakan **Golang**, **Fiber**, **GORM**, dan **PostgreSQL**.

---

## 🚀 Fitur Utama

- Manajemen rumah & penduduk
- Pembayaran iuran IPL bulanan
- Iuran ronda malam dan pelaporan
- Komunitas posting & komentar
- Marketplace (warung warga)
- Pengumuman dari pengurus RT
- Autentikasi & role-based access
- D.L.L

---

## 🛠️ Teknologi

- **Golang** (1.21+)
- **Fiber** - Web framework
- **GORM** - ORM untuk PostgreSQL
- **JWT** - Autentikasi
- **PostgreSQL** - Database relasional
- **Docker** (optional - untuk dev & deploy)

---

## 📦 Struktur Proyek

```bash
sirnawa-backend/
├── 📁 cmd/                   # Entry point (main.go)
├── 📁 internal/
│   ├── 📁 config/            # Load .env & konfigurasi
│   ├── 📁 database/          # Inisialisasi DB & migrasi
│   └── 📁 handlers/          # Fiber handler
│   └── 📁 middleware/        # JWT, Auth middleware
│   └── 📁 models/            # GORM model 
│   └── 📁 repository/        # Akses database
│   └── 📁 routes/            # Inisialisasi semua route 
│   └── 📁 services/          # Bisnis logic
├── 📁 pkg/
│   └── 📁 utils/             # Helper functions
│   └── 📁 validators/        # Validator custom
├── Makefile                  # Build semua services
├── .env                      # Variabel lingkungan# Build semua services
└── README.md                 # File ini

```
---

## 🧪 Cara Menjalankan (Development)

```bash
# clone repo
git clone https://github.com/IKHINtech/sirnawa-backend.git
cd sirnawa-backend

# install dependency
go mod tidy

# jalankan aplikasi
go run cmd/main.go
