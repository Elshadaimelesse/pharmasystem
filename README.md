# PharmRx - Multi-Pharmacy Management & Prescription Verification System

A complete polyglot healthcare platform enabling patients to verify medications from home, upload prescriptions for clinical pharmacist review, compare drug prices across local pharmacies, and receive comprehensive administration guides ("How to Take It").

## Project Architecture
- **Frontend**: React 19, TypeScript, Tailwind CSS, Lucide Icons, Motion
- **Backend (Go)**: `backend-go/` High-performance microservice for pharmacy orders and prescription queuing
- **Backend (Rust)**: `backend-rust/` High-throughput Axum/Actix-web service for drug inventory and rapid price indexing
- **Backend (Node/JS)**: `backend-node/` Express.js REST API with AI Prescription OCR & interaction verification
- **Deployment**: Docker Compose included

## 1. Running the React Frontend
```bash
cd frontend-react
npm install
npm run dev
# Running at http://localhost:3000
```

## 2. Running the Go Backend
```bash
cd backend-go
go mod tidy
go run main.go
# Server listening on :8080
```

## 3. Running the Rust Backend
```bash
cd backend-rust
cargo run
# Server listening on :8081
```

## 4. Running the Node.js / JavaScript Backend
```bash
cd backend-node
npm install
node server.js
# Server listening on :5000
```

## Features
- **Client Home Prescription Checker**: Upload prescription photos/PDFs with OCR parsing
- **Pharmacist Review Dashboard**: Multi-pharmacy druggist verification portal
- **Live Price Comparison**: Compare across registered pharmacies
- **Comprehensive Drug Guide**: Full clinical dosage, side effects, dietary warnings, storage
- **Multi-Role Authentication**: Dedicated portals for Patients & Pharmacies
