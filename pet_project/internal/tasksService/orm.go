package tasksService

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
