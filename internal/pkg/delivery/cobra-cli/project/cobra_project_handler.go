package project

import (
	"github.com/spf13/cobra"
	domain "jirno/internal/pkg/domain/project"
	"jirno/internal/pkg/localstorage"
)

type projectHandler struct {
	pUsecase domain.IProjectUsecase
	storage  localstorage.LocalStorage
}

func NewProjectHandler(root *cobra.Command, projectUsecase domain.IProjectUsecase, lStorage localstorage.LocalStorage) {
	handler := projectHandler{pUsecase: projectUsecase, storage: lStorage}
	projectRoot := &cobra.Command{Use: "project"}

	addProjectCreateHandler(projectRoot, &handler)

	addProjectUpdateHandler(projectRoot, &handler)

	addProjectDeleteHandler(projectRoot, &handler)

	addProjectFilterHandler(projectRoot, &handler)

	root.AddCommand(projectRoot)
}
