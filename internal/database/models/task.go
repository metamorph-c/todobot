package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID     string `gorm:"type:char(36);primary_key"`
	ChatId int64  `gorm:"column:chat_id"`
	Task   string `gorm:"column:task"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}
