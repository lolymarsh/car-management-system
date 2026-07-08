# Car Management System

ระบบบันทึกข้อมูลรถยนต์ของบริษัท

---

## Tech Stack

- หลังบ้าน: Go + Echo + Bun + PostgreSQL
- หน้าบ้าน: React (Vite) + Ant Design

## Structure คร่าวๆ

```
backend/          # Go API
  main.go
  internal/
    car/          # CRUD รถ
    model/       
    server/       # ตั้งค่า Echo + routes
  migrations/
  Dockerfile

frontend/         # React
  src/
    pages/
      CarList.jsx
      CarForm.jsx
    components/
      AppLayout.jsx
    api.js        # axios เรียก backend
    App.jsx
  vite.config.js
```

---

## รันยังไง

เปิด 2 terminal

### Terminal 1 — Backend

```bash
cd backend
cp _env_example .env
docker compose up -d --build
```

หรือ ถ้าไม่มี docker ก็ต้องมี postgres local แล้ว

```bash
cd backend
go run main.go
```

อย่าลืมรัน migrate ก่อนครั้งแรก หรือถ้ามี migration ใหม่

```bash
cd backend
make migrate-up
```

ถ้าอยากย้อนกลับ

```bash
make migrate-down
```

API พร้อมที่ `http://localhost:8080`

### Terminal 2 — Frontend

```bash
cd frontend
npm install
npm run dev
```

เปิด `http://localhost:5173`

---

## API มีแค่นี้ (ไม่ต้อง login)
---
| Method | Path | เอาไว้ |
|---|---|---|
| POST | `/api/cars` | เพิ่มรถ |
| POST | `/api/cars/filter` | ค้นหา + filter + paginate |
| GET | `/api/cars/:id` | ดูรายละเอียด |
| PUT | `/api/cars/:id` | แก้ไข |
| DELETE | `/api/cars/:id` | ลบ (soft delete) |
---