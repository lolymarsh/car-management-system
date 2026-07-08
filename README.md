# Car Management System

เว็บแอปพลิเคชันสำหรับบันทึกและจัดการข้อมูลรถยนต์ของบริษัท ใช้เทคโนโลยี **Go (Echo) + PostgreSQL + Redis** ฝั่ง Backend และ **React + Ant Design** ฝั่ง Frontend

---

## Tech Stack

| Layer | Technology |
|---|---|
| Backend | Go 1.25, Echo v4, Bun ORM |
| Database | PostgreSQL 16 |
| Cache / Token Store | Redis 7 |
| Frontend | React.js, Ant Design 5 *(กำลังพัฒนา)* |
| Auth | JWT (Access Token 15m + Refresh Token 7d) |
| Migration | Goose |
| Container | Docker + Docker Compose |

---

## Project Structure

```
car-management-system/
├── backend/
│   ├── internal/         # Application logic
│   │   ├── auth/         # Authentication (register, login, refresh, sessions)
│   │   ├── auditlog/     # Audit log (history tracking)
│   │   ├── car/          # Car CRUD + filter + cache
│   │   ├── user/         # User profile management
│   │   ├── router/       # Route registration
│   │   ├── server/       # DI container + server setup
│   │   └── testutil/     # Integration test helpers
│   ├── migrations/       # Database migrations (PostgreSQL + MySQL)
│   │   ├── postgres/
│   │   └── mysql/
│   ├── pkg/              # Shared packages
│   │   ├── cache/        # Redis cache utility
│   │   ├── token/        # Refresh token store (Redis/DB)
│   │   └── ...           # common, middleware, response, etc.
│   ├── docs/             # API documentation
│   ├── Dockerfile
│   └── docker-compose.yml
├── frontend/             # React.js frontend *(coming soon)*
├── .gitignore
└── README.md
```

---

## Quick Start (Docker)

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) + [Docker Compose](https://docs.docker.com/compose/install/)

### Run

```bash
cd backend
docker compose up -d --build
```

รอสักครู่แล้วทดสอบ:

```bash
# Login as admin
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@loly.com","password":"newpass1234"}'

# Create car (ADMIN only)
curl -X POST http://localhost:8080/api/cars \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"registration_number":"กข-1234","brand":"Toyota","model":"Corolla"}'

# Search cars (USER + ADMIN)
curl -X POST http://localhost:8080/api/cars/filter \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{"page":1,"page_size":10}'
```

### Seed Users (Local Dev Only)

Migration จะ insert อัตโนมัติเมื่อรันครั้งแรก

| Role | Email | Password |
|---|---|---|
| ADMIN | `admin@loly.com` | `newpass1234` |
| USER | `user@loly.com` | `password123` |

### Role-Based Access

| Feature | USER | ADMIN |
|---|---|---|
| ดูรายการ + ค้นหารถ | ✅ | ✅ |
| เพิ่ม/แก้ไข/ลบรถ | 🚫 | ✅ |
| จัดการผู้ใช้ | 🚫 | ✅ |
| ดู Audit Log ของตัวเอง | ✅ | ✅ |
| ดู Audit Log ทั้งหมด | 🚫 | ✅ |
| ตั้งค่า Profile + เปลี่ยน Password | ✅ | ✅ |
| ดู Sessions / Revoke | ✅ | ✅ |

### Stop

```bash
cd backend
docker compose down
```

ลบ data ด้วย (DB + Redis):

```bash
docker compose down -v
```

---

## Local Development (No Docker)

### Prerequisites

- Go 1.25+
- PostgreSQL 16+
- Redis 7+ (optional แต่แนะนำ)
- [swag](https://github.com/swaggo/swag) (สำหรับ regenerate API docs)

### Setup

```bash
cd backend

# ติดตั้ง dependencies
make install

# ตั้งค่า environment variables
cp _env_example .env
# แก้ไข .env ให้ตรงกับ local ของคุณ (DB_HOST, DB_USER, DB_PASSWORD ฯลฯ)
```

### Run

```bash
# รัน database migration อัตโนมัติเมื่อเริ่ม server

# โหมด dev (hot reload, ต้องติดตั้ง air)
make dev

# หรือรันตรง ๆ
make run
```

API จะอยู่ที่ `http://localhost:8080`

### API Documentation

```bash
# Regenerate Swagger docs
make swagger

# เปิดใน browser
open http://localhost:8080/swagger/index.html
open http://localhost:8080/redoc
```

### Useful Commands

```bash
make build      # Build binary
make fmt        # Format code
make vet        # Static analysis
make lint       # Lint
make check      # Quick check (lint-quick + vet)
```

### Testing

```bash
# เริ่ม test infrastructure (PostgreSQL + MySQL + Redis)
make docker-test-up

# รัน integration tests
make test-integration

# หยุด test infrastructure
make docker-test-down
```

---

## API Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|---|
| Method | Path | Auth | Role | Description |
|---|---|---|---|---|---|
| `POST` | `/api/auth/register` | ✗ | - | Register new user |
| `POST` | `/api/auth/login` | ✗ | - | Login (captures User-Agent + IP) |
| `POST` | `/api/auth/refresh` | ✗ | - | Refresh access token |
| `POST` | `/api/auth/revoke` | ✓ | ANY | Revoke refresh token |
| `PATCH` | `/api/auth/change-password` | ✓ | ANY | Change password |
| `POST` | `/api/auth/revoke-all` | ✓ | ANY | Revoke all sessions |
| `GET` | `/api/auth/sessions` | ✓ | ANY | List active sessions (UA + IP) |
| `POST` | `/api/cars` | ✓ | ADMIN | Create a new car |
| `POST` | `/api/cars/filter` | ✓ | ANY | List + search + filter + paginate |
| `GET` | `/api/cars/:id` | ✓ | ANY | Get car details |
| `PUT` | `/api/cars/:id` | ✓ | ADMIN | Update car (optimistic lock) |
| `DELETE` | `/api/cars/:id` | ✓ | ADMIN | Soft-delete car |
| `GET` | `/api/user/profile` | ✓ | ANY | Get current user profile |
| `PATCH` | `/api/user/profile` | ✓ | ANY | Update profile |
| `POST` | `/api/audit-log/filter` | ✓ | ANY | Filter audit logs |
| `GET` | `/api/audit-log/detail` | ✓ | ANY | Get audit log detail |

### Car Fields

| Field | Type | Required | Description |
|---|---|---|---|
| `registration_number` | string | ✓ | ทะเบียนรถยนต์ |
| `brand` | string | ✓ | ยี่ห้อรถ |
| `model` | string | ✓ | รุ่นรถ |
| `color` | string | ✗ | สี |
| `year` | int | ✗ | ปีที่ผลิต |
| `notes` | string | ✗ | หมายเหตุ |

---

## Roadmap

- [x] Backend CRUD API (Go + Echo)
- [x] PostgreSQL + Redis (Docker Compose)
- [x] Auth: Register, Login, JWT, Refresh Token
- [x] Role-based access (ADMIN / USER)
- [x] Car filter + search + pagination
- [x] Redis cache for car & user
- [x] Session management (list tokens, revoke, device info)
- [x] Change password
- [x] Audit log
- [ ] Frontend (React.js + Ant Design)
- [ ] Git init + push to public repository

---

## License

Internal project — Haupcar Co., Ltd. Technical Assessment
