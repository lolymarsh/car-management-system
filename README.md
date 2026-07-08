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

เปด 2 terminal

### Terminal 1 - Backend

step แรก ขึ้น docker และ migrate database

```bash
cd backend
make setup
```

จากนั้น start backend

```bash
make run
```

API http://localhost:8080

### Terminal 2 - Frontend

```bash
cd frontend
npm install
npm run dev
```

เปิด http://localhost:5173

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