package smart_parser

import (
	"jirno/internal/pkg/domain/project"
	"jirno/internal/pkg/domain/task"
)

type ISmartProjectParser interface {
	Parse(string, *project.ProjectFilter)
}

type ISmartTaskParser interface {
	Parse(string, *task.TaskFilter)
}
