# Vellum - High Level Architecture Specification

## Overview

Vellum is a full-stack personal finance and budgeting platform.

The platform enables users to manage their finances through budgeting, transaction tracking, financial goals, subscriptions and analytics.

The application follows a modern client-server architecture.

---

# Technology Stack

## Frontend

- React
- TypeScript
- Vite
- Tailwind CSS
- React Router

Responsibilities

- User Interface
- Authentication
- Dashboard
- Transactions
- Budgets
- Goals
- Accounts
- Subscriptions
- Settings
- Analytics
- REST API Communication

---

## Backend

- Go
- Gin
- GORM
- JWT Authentication
- Repository Pattern
- Service Layer
- Dependency Injection

Responsibilities

- REST API
- Authentication
- Authorization
- Business Logic
- Validation
- Financial Calculations
- Data Persistence

---

## Database

- PostgreSQL

Stores

- Users
- User Settings
- Accounts
- Transactions
- Categories
- Budgets
- Goals
- Subscriptions

---

## Infrastructure

- Docker
- Docker Compose

Containers

- Frontend
- Backend
- PostgreSQL

---

# High Level System

The system consists of three primary components.

Frontend

↓

REST API

↓

Backend

↓

PostgreSQL

---

# Frontend Responsibilities

The frontend is responsible for

- User Authentication
- Dashboard
- Viewing Transactions
- Managing Accounts
- Creating Budgets
- Managing Goals
- Managing Subscriptions
- Updating User Settings
- Displaying Analytics

The frontend communicates only with the backend using HTTPS REST APIs.

It never communicates directly with PostgreSQL.

---

# Backend Responsibilities

The backend exposes REST endpoints.

Internally it consists of

- Routes
- Middleware
- Handlers
- Services
- Repositories

Authentication is handled using JWT.

Business logic exists inside Services.

Repositories are responsible for all database communication.

---

# Internal Backend Architecture

HTTP Request

↓

Routes

↓

JWT Middleware

↓

Handlers

↓

Services

↓

Repositories

↓

PostgreSQL

---

# Authentication

Authentication uses JWT.

Flow

User Login

↓

JWT Generated

↓

Frontend Stores Token

↓

Frontend sends

Authorization: Bearer <token>

↓

Backend validates JWT

↓

Protected endpoint executes

---

# Core Business Modules

Authentication

Users

Accounts

Transactions

Categories

Budgets

Goals

Subscriptions

Dashboard

Analytics

Settings

---

# Database Entities

User

UserSettings

Account

Transaction

Category

Budget

Goal

Subscription

Relationships

User owns

- Accounts
- Transactions
- Budgets
- Goals
- Subscriptions
- User Settings

Transactions belong to

- Account
- Category

Budgets belong to

- Category

---

# External Integrations (Future)

Open Banking APIs

Exchange Rate APIs

Email Service

Push Notifications

AI Financial Insights

---

# Architectural Principles

- Layered Architecture
- Repository Pattern
- Dependency Injection
- Service Layer
- Stateless REST API
- JWT Authentication
- Separation of Concerns
- Containerised Deployment

---

# Deployment

Docker Compose orchestrates

Frontend Container

Backend Container

PostgreSQL Container

All communication between frontend and backend occurs over HTTPS.

The backend is the only component with access to PostgreSQL.

---

# Target Architecture

```
                 User
                  │
                  ▼
      +----------------------+
      | React + Vite Client  |
      +----------------------+
                  │
          HTTPS / REST API
                  │
                  ▼
      +----------------------+
      |   Go Backend (Gin)   |
      +----------------------+
                  │
      +----------------------+
      | Authentication       |
      | Business Logic       |
      | Validation           |
      | Services             |
      | Repositories         |
      +----------------------+
                  │
                  ▼
      +----------------------+
      |    PostgreSQL DB     |
      +----------------------+

Future Integrations

Open Banking
Exchange Rates
Notifications
AI Services
```

The backend is the central component responsible for authentication, business logic, validation and persistence.

The frontend is responsible only for presentation and user interaction.

The database is accessed exclusively through the backend.