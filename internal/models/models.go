package models

import "time"

type Task struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	DueDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ReminderRule struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Offset    int
	Unit      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuditLog struct {
	ID          uint `gorm:"primaryKey"`
	TaskID      *uint
	RuleID      *uint
	Event       string
	TriggeredAt time.Time
	Message     string
	CreatedAt   time.Time
}
