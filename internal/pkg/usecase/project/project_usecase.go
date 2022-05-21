package project

import (
	"fmt"
	"github.com/google/uuid"
	"jirno/internal/pkg/domain"
)

type projectUsecase struct {
	repo domain.IProjectRepo
}

func NewProjectUsecase(projectRepo domain.IProjectRepo) domain.IProjectUsecase {
	return &projectUsecase{
		repo: projectRepo,
	}
}

func (u projectUsecase) GetByID(id uuid.UUID) (*domain.Project, error) {
	return u.repo.GetByID(id)
}

func (u projectUsecase) Update(project domain.ProjectUpdate) error {
	return u.repo.Update(project)
}

func (u projectUsecase) Delete(id uuid.UUID) error {
	return u.repo.Delete(id)
}

func (u projectUsecase) Create(project domain.Project) (uuid.UUID, error) {
	project.ID = uuid.New()
	err := u.repo.Create(project)
	if err != nil {
		return uuid.UUID{}, err
	}
	return project.ID, nil
}

func (u projectUsecase) Complete(id uuid.UUID) error {
	update :=  domain.ProjectUpdate{
		ID:          id,
	}
	update.IsCompleted = new(bool)
	*update.IsCompleted = true
	return u.repo.Update(update)
}

func (u projectUsecase) GetByFilter(filter domain.SmartProjectFilter) ([]domain.Project, error) {
	f, err := filter.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("project get by filter (filter casting): %v", err)
	}
	return u.repo.GetByFilter(*f)
}
