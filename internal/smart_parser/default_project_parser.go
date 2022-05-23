package smart_parser

import (
	"jirno/internal/pkg/domain/project"
	"jirno/internal/pkg/utils"
	"time"
)

type defProjectParser struct {
}

func NewDefaultProjectParser() ISmartProjectParser {
	return &defProjectParser{}
}

func (d defProjectParser) Parse(s string, filter *project.ProjectFilter) {
	if s == "" || filter == nil{
		return
	}

	filter.StartDate = new(time.Time)
	filter.EndDate = new(time.Time)

	now := time.Now()
	switch s {
	case "today":
		*filter.StartDate, *filter.EndDate = utils.GetDayRange(now)
	case "tomorrow":
		*filter.StartDate, *filter.EndDate = utils.GetDayRange(now.AddDate(0, 0, -1))
	}
}
