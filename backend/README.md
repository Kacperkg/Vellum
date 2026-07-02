# Vellum Backend

Backend API for **Vellum**, a personal budgeting and finance management application built with **Go**, **Gin**, **GORM**, **PostgreSQL**, and **Docker**.

---

# Tech Stack

- Go
- Gin
- GORM
- PostgreSQL
- Docker & Docker Compose

---

# MVP

## Dashboard (Overview)

The dashboard provides a high-level summary of the user's financial health.

### Dashboard Information

- Current Balance
- Total Income
- Total Expenses
- Remaining Budget Available
- Recent Activity
- Top Spending Categories (Current Month)
- Monthly Recurring Spend
- Next Recurring Payment
- Largest Recurring Transaction

---

## User

### Authentication

- Register
- Login
- Logout
- JWT Authentication
- Secure Password Hashing

### Profile

- Display Name
- Email Address

### Preferences

- Preferred Currency
- Date Format
- Number Format
- First Day of Week

---

## Accounts

Users can organise their transactions by account.

### Features

- Create Account
- Update Account
- Delete Account

### Account Fields

- Name

Examples:

- Barclays
- Monzo
- Cash
- Revolut

---

## Transactions

The core feature of the application.

### Features

- Create Transaction
- Update Transaction
- Delete Transaction
- View Transaction History
- Filter Transactions
- Export Transactions as CSV

### Transaction Fields

- Name
- Amount
- Type
  - Income
  - Expense
- Category
- Account
- Date
- Frequency
  - One-off
  - Daily
  - Weekly
  - Monthly
  - Yearly

---

## Categories

Predefined categories used across the application.

### Features

- View Categories

### Category Fields

- Name

---

## Budgets

Manage spending limits by category.

### Features

- Create Budget
- Update Budget
- Delete Budget

### Stored Fields

- Category
- Allocated Amount

### Calculated Values

- Amount Spent
- Remaining Budget
- Budget Utilisation

---

## Analytics

Financial insights generated from user transactions.

### Analytics

- Monthly Income vs Expenses
- Cash Flow
- Savings Trajectory
- Net Monthly Surplus
- Top 6 Spending Categories
- Recurring Spending Analysis

---

## Goals

Track personal savings goals.

### Features

- Create Goal
- Update Goal
- Delete Goal

### Goal Fields

- Goal Name
- Target Amount
- Saved Amount
- Monthly Contribution

### Calculated Values

- Remaining Amount

---

# Core Backend Responsibilities

- Authentication
- Authorization
- User Preferences
- CRUD Operations
- Database Persistence
- Request Validation
- Error Handling
- CSV Export
- Business Logic

---

# Database Models

- User
- UserSettings
- Account
- Category
- Transaction
- Budget
- Goal

---

# Project Structure

```text
cmd/
└── api/

internal/
├── config/
├── database/
├── handlers/
├── middleware/
├── models/
├── repositories/
├── routes/
├── server/
└── services/
```

## Folder Responsibilities

| Folder | Responsibility |
|----------|----------------|
| config | Environment configuration |
| database | Database connection & migrations |
| handlers | HTTP handlers |
| middleware | Authentication & middleware |
| models | Database models |
| repositories | Database access layer |
| routes | API route registration |
| server | Application startup |
| services | Business logic |

---

# Backend Roadmap

## Foundation

- [x] Docker
- [x] Docker Compose
- [x] Configuration
- [x] Gin
- [x] PostgreSQL
- [x] GORM
- [x] Auto Migrations

## Architecture

- [ ] Repository Pattern
- [ ] Service Layer
- [ ] Dependency Injection
- [ ] DTOs
- [ ] Validation

## Authentication

- [ ] Register
- [ ] Login
- [ ] JWT
- [ ] Authorization Middleware
- [ ] Protected Routes

## Features

- [ ] Accounts
- [ ] Categories
- [ ] Transactions
- [ ] Budgets
- [ ] Goals
- [ ] Analytics
- [ ] Dashboard

## Quality

- [ ] Unit Testing
- [ ] Integration Testing
- [ ] Logging
- [ ] Configuration Improvements
- [ ] Production Migrations

---

# Development

## Start Application

```bash
docker compose up --build
```

---

## Stop Application

```bash
docker compose down
```

---

## Stop & Remove Database

```bash
docker compose down -v
```

Removes the PostgreSQL volume and starts with a fresh database.

---

## Rebuild Containers

```bash
docker compose up --build --force-recreate
```

---

# PostgreSQL

## Open PostgreSQL

```bash
docker exec -it vellum-postgres psql -U vellum -d vellum
```

---

## List Tables

```sql
\dt
```

---

## Describe Table

```sql
\d users
```

Example

```sql
\d transactions
```

---

## View Table Data

```sql
SELECT * FROM users;
```

Example

```sql
SELECT * FROM transactions;
```

---

## Count Rows

```sql
SELECT COUNT(*) FROM users;
```

---

## List Databases

```sql
\l
```

---

## Current Database

```sql
SELECT current_database();
```

---

## Exit PostgreSQL

```sql
\q
```

---

# Go Commands

## Download Dependencies

```bash
go mod tidy
```

---

## Format Code

```bash
go fmt ./...
```

---

## Run Tests

```bash
go test ./...
```

---

## Download Modules

```bash
go mod download
```

---

## Verify Modules

```bash
go mod verify
```

---

# Docker Commands

## Running Containers

```bash
docker ps
```

---

## Running Images

```bash
docker images
```

---

## Backend Logs

```bash
docker compose logs backend
```

---

## Frontend Logs

```bash
docker compose logs frontend
```

---

## PostgreSQL Logs

```bash
docker compose logs postgres
```

---

## Follow Backend Logs

```bash
docker compose logs -f backend
```

---

## Enter Backend Container

```bash
docker exec -it vellum-backend-1 sh
```

---

## Enter Frontend Container

```bash
docker exec -it vellum-frontend-1 sh
```

---

## Enter PostgreSQL Container

```bash
docker exec -it vellum-postgres sh
```

---

# Future Ideas

These features are intentionally outside the MVP.

- Email Verification
- Password Reset
- Open Banking Integration
- Investment Tracking
- Receipt OCR
- Import Bank Statements
- Notifications
- Shared Budgets
- AI Insights
- Mobile Push Notifications
- Admin Dashboard
- Family Accounts
- Scheduled Reports
- Financial Forecasting

---

# Success Criteria

The MVP is complete when a user can:

- Register an account
- Log in securely
- Manage accounts
- View predefined categories
- Create, edit and delete transactions
- Create and manage budgets
- Create and manage goals
- View dashboard overview
- View analytics
- Export transactions as CSV
- Persist all data in PostgreSQL