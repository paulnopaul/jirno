package task

import (
	"github.com/google/uuid"
	"time"
)

type TaskUpdate struct {
	ID            uuid.UUID
	User          *int64
	Project       *uuid.UUID
	Title         string
	Description   string
	Additional    map[string]string
	IsCompleted   *bool
	CompletedDate *time.Time
	DateTo        *time.Time
}
