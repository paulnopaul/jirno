package smart_parser

import (
	"jirno/internal/pkg/domain/task"
	"jirno/internal/pkg/utils"
	"time"
)

type defTaskParser struct {
}

func NewDefaultTaskParser() ISmartTaskParser {
	return defTaskParser{}
}

func (d defTaskParser) Parse(s string, filter *task.TaskFilter) {
	if s == "" || filter == nil {
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
