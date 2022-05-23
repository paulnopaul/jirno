package project

import (
	"github.com/google/uuid"
	"time"
)

type ProjectUpdate struct {
	ID            uuid.UUID
	Title         string
	Description   string
	ParentProject *uuid.UUID
	Users         []int64
	IsCompleted   *bool
	CompletedDate *time.Time
	DateTo        *time.Time
}
