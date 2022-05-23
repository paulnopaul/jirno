package project

import (
	"fmt"
	"github.com/google/uuid"
	"jirno/internal/pkg/domain/project"
	"jirno/internal/smart_parser"
)

type projectUsecase struct {
	repo        project.IProjectRepo
	smartParser smart_parser.ISmartProjectParser
}

func NewProjectUsecase(projectRepo project.IProjectRepo, smartFilterParser smart_parser.ISmartProjectParser) project.IProjectUsecase {
	return &projectUsecase{
		repo:        projectRepo,
		smartParser: smartFilterParser,
	}
}

func (u projectUsecase) GetByID(id uuid.UUID) (*project.Project, error) {
	return u.repo.GetByID(id)
}

func (u projectUsecase) Update(project project.ProjectUpdate) error {
	return u.repo.Update(project)
}

func (u projectUsecase) Delete(id uuid.UUID) error {
	return u.repo.Delete(id)
}

func (u projectUsecase) Create(project project.Project) (uuid.UUID, error) {
	project.ID = uuid.New()
	err := u.repo.Create(project)
	if err != nil {
		return uuid.UUID{}, err
	}
	return project.ID, nil
}

func (u projectUsecase) Complete(id uuid.UUID) error {
	update := project.ProjectUpdate{
		ID: id,
	}
	update.IsCompleted = new(bool)
	*update.IsCompleted = true
	return u.repo.Update(update)
}

func (u projectUsecase) GetByFilter(filter project.SmartProjectFilter) ([]project.Project, error) {
	f, err := filter.ToDomain()
	if err != nil {
		return nil, fmt.Errorf("project get by filter (filter casting): %v", err)
	}
	u.smartParser.Parse(filter.Smart, f)
	return u.repo.GetByFilter(*f)
}
