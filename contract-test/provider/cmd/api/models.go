package main

import (
	"time"
)

type Class struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Discipline string         `json:"discipline"`
	Day        string         `json:"day"`
	Hour       string         `json:"hour"`
	Students   []StudentClass `json:"students"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

type StudentClass struct {
	ID        uint
	ClassID   uint
	StudentId uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
