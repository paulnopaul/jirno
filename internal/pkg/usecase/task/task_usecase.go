package task

import (
	"fmt"
	"github.com/google/uuid"
	"jirno/internal/pkg/domain"
)

type taskUsecase struct {
	repo domain.ITaskRepo
}

func NewTaskUsecase(taskRepo domain.ITaskRepo) domain.ITaskUsecase {
	return &taskUsecase{
		repo: taskRepo,
	}
}

func (t taskUsecase) Create(task domain.Task) (uuid.UUID, error) {
	task.ID = uuid.New()
	return task.ID, t.repo.Create(task)
}

func (t taskUsecase) GetByID(id uuid.UUID) (*domain.Task, error) {
	panic("implement me")
}

func (t taskUsecase) GetByFilter(filter domain.SmartTaskFilter) ([]domain.Task, error) {
	domainFilter, err := filter.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("task get by filter (filter casting): %v", err)
	}
	return t.repo.GetByFilter(*domainFilter)
}

func (t taskUsecase) Complete(id uuid.UUID) error {
	update := domain.TaskUpdate{
		ID: id,
	}
	update.IsCompleted = new(bool)
	*update.IsCompleted = true
	return t.repo.Update(update)
}

func (t taskUsecase) Update(update domain.TaskUpdate) error {
	return t.repo.Update(update)
}

func (t taskUsecase) Delete(id uuid.UUID) error {
	return t.repo.Delete(id)
}
