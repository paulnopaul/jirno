package project

import (
	"github.com/spf13/cobra"
	domain "jirno/internal/pkg/domain/project"
)

type projectHandler struct {
	pUsecase domain.IProjectUsecase
}

func NewProjectHandler(root *cobra.Command, projectUsecase domain.IProjectUsecase) {
	handler := projectHandler{pUsecase: projectUsecase}
	projectRoot := &cobra.Command{Use: "project"}

	addProjectCreateHandler(projectRoot, &handler)

	addProjectUpdateHandler(projectRoot, &handler)

	addProjectDeleteHandler(projectRoot, &handler)

	addProjectFilterHandler(projectRoot, &handler)

	root.AddCommand(projectRoot)
}
