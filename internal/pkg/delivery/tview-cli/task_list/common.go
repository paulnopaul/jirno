package task_list

import (
	"fmt"
	"jirno/internal/pkg/domain"
	"jirno/internal/pkg/utils"
	"time"
)

func taskToString(t domain.Task) string {
	completed := ' '
	if t.IsCompleted == true {
		completed = 'x'
	}
	dateStr := ""
	ts, te := utils.GetDayRange(time.Now())
	if t.DateTo != nil {
		if t.DateTo.Unix() > ts.Unix() && t.DateTo.Unix() < te.Unix() {
			dateStr = "today"
		} else {
			dateStr = t.DateTo.Format("2006-01-03")
		}
	}
	dateStr = fmt.Sprintf("(%v)", dateStr)
	return fmt.Sprintf("[ %c ] %v ", completed, t.Title) + dateStr
}
