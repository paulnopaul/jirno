package task

import (
	"fmt"
	"github.com/google/uuid"
	"jirno/internal/pkg/domain/task"
	"jirno/internal/smart_parser"
)

type taskUsecase struct {
	repo        task.ITaskRepo
	smartParser smart_parser.ISmartTaskParser
}

func NewTaskUsecase(taskRepo task.ITaskRepo, smartFilterParser smart_parser.ISmartTaskParser) task.ITaskUsecase {
	return &taskUsecase{
		repo:        taskRepo,
		smartParser: smartFilterParser,
	}
}

func (t taskUsecase) Create(task task.Task) (uuid.UUID, error) {
	task.ID = uuid.New()
	return task.ID, t.repo.Create(task)
}

func (t taskUsecase) GetByID(id uuid.UUID) (*task.Task, error) {
	return t.GetByID(id)
}

func (t taskUsecase) GetByFilter(filter task.DeliveryTaskFilter) ([]task.Task, error) {
	domainFilter, err := filter.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("task get by filter (filter casting): %v", err)
	}
	t.smartParser.Parse(filter.Smart, domainFilter)
	return t.repo.GetByFilter(*domainFilter)
}

func (t taskUsecase) Complete(id uuid.UUID) error {
	update := task.TaskUpdate{
		ID: id,
	}
	update.IsCompleted = new(bool)
	*update.IsCompleted = true
	return t.repo.Update(update)
}

func (t taskUsecase) Update(update task.TaskUpdate) error {
	return t.repo.Update(update)
}

func (t taskUsecase) Delete(id uuid.UUID) error {
	return t.repo.Delete(id)
}
