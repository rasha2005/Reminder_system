# Reminder System API

A backend reminder management system built using Go, Gin, GORM, and PostgreSQL.

This system supports:

- Creating, updating, deleting, and listing reminder rules
- Activating / deactivating reminder rules
- Background scheduler to trigger reminders
- Audit trail logging for:
  - Reminder rule changes
  - Reminder executions

---

## Tech Stack

- Go
- Gin (HTTP framework)
- GORM (ORM)
- PostgreSQL
- Environment variables (.env)
- Background scheduler using time.Ticker

---

## Project Structure

internal/
    handlers/
    routes/
    models/
    database/
cmd/
    main.go

---

## Setup Instructions

### 1. Clone the Repository

cd reminder-system

---

### 2. Create .env File

Create a `.env` file in the root directory:

DB_HOST=localhost  
DB_USER=postgres  
DB_PASSWORD=yourpassword  
DB_NAME=reminder_db  
DB_PORT=5432  

---

### 3. Install Dependencies

go mod tidy

---

### 4. Run the Application

go run ./cmd

---

## API Endpoints

### Reminder Rules

#### Create Rule
POST /rules

{
  "name": "10 minutes before",
  "offset": 10,
  "unit": "minutes"
}

---

#### Get All Rules
GET /rules

---

#### Update Rule
PUT /rules/:id

{
  "name": "15 minutes before",
  "offset": 15,
  "unit": "minutes",
  "isActive": true
}

---

#### Activate / Deactivate Rule
PATCH /rules/:id/activate

{
  "isActive": false
}

---

#### Delete Rule
DELETE /rules/:id

---

### Audit Logs

#### Get All Logs
GET /audits

#### Filter Logs

GET /audits?type=rule  
GET /audits?type=execution  

---

## Scheduler Behavior

- Background scheduler runs at fixed intervals.
- Only active reminder rules are evaluated.
- When a reminder condition is met, it is triggered.
- Each execution is recorded in the audit_logs table.
- Duplicate reminders are prevented.

---

## Database Tables

### Tasks
- ID
- Title
- Description
- DueDate
- CreatedAt

### ReminderRules
- ID
- Name
- Offset
- Unit (minutes / hours / days)
- IsActive
- CreatedAt

### AuditLogs
- ID
- RuleID
- Event
- Message
- TriggeredAt

---

## Features Implemented

- Reminder rule CRUD operations
- Rule activation/deactivation
- Background reminder scheduler
- Audit trail logging
- Clean modular architecture
- Environment-based configuration
