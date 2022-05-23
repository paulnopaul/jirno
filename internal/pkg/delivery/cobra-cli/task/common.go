package task

import (
	"fmt"
	domain "jirno/internal/pkg/domain/task"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type taskType int8

const (
	typeTask       taskType = 0
	typeTaskUpdate taskType = 1
)

func parseTask(cmd *cobra.Command, args []string, t taskType) (*domain.DeliveryTask, error) {
	res := &domain.DeliveryTask{}

	if t == typeTaskUpdate {
		parsedId, err := cmd.Flags().GetString("id")
		if err != nil {
			return nil, fmt.Errorf("id flag error: %v", err)
		}
		res.ID = parsedId
	}

	parsedUid, err := cmd.Flags().GetInt64("uid")
	if err != nil {
		return nil, fmt.Errorf("user id flag error: %v", err)
	}
	if parsedUid != 0 {
		res.User = new(int64)
		*res.User = parsedUid
	}

	parsedPid, err := cmd.Flags().GetString("pid")
	if err != nil {
		return nil, fmt.Errorf("project id flag error: %v", err)
	}
	res.Project = parsedPid

	parsedDescription, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("description flag error: %v", err)
	}
	res.Description = parsedDescription

	if len(args) != 0 {
		res.Title = args[0]
	}

	parsedCompleted, err := cmd.Flags().GetBoolSlice("completed")
	if err != nil {
		return nil, fmt.Errorf("completed flag error: %v", err)
	}
	if parsedCompleted != nil && len(parsedCompleted) != 0 {
		res.IsCompleted = new(bool)
		*res.IsCompleted = parsedCompleted[0]
	}

	parsedDateCompleted, err := cmd.Flags().GetString("datecompleted")
	if err != nil {
		return nil, fmt.Errorf("datecompleted flag error: %v", err)
	}
	if parsedDateCompleted != "" {
		res.CompletedDate = new(time.Time)
		*res.CompletedDate, err = time.Parse(time.RFC3339, parsedDateCompleted)
	}

	parsedDateTo, err := cmd.Flags().GetString("dateto")
	if err != nil {
		return nil, fmt.Errorf("dateto flag error: %v", err)
	}
	if parsedDateTo != "" {
		res.DateTo = new(time.Time)
		*res.DateTo, err = time.Parse(time.RFC3339, parsedDateTo)
	}

	return res, nil
}

func parseFilter(cmd *cobra.Command, args []string) (*domain.DeliveryTaskFilter, error) {
	res := &domain.DeliveryTaskFilter{}

	if len(args) > 0 {
		res.Smart = args[0]
	}

	parsedStartDate, err := cmd.Flags().GetString("datestart")
	if err != nil {
		return nil, fmt.Errorf("datecompleted flag error: %v", err)
	}
	if parsedStartDate != "" {
		res.StartDate = new(time.Time)
		*res.StartDate, err = time.Parse(time.RFC3339, parsedStartDate)
	}

	parsedEndDate, err := cmd.Flags().GetString("dateend")
	if err != nil {
		return nil, fmt.Errorf("datecompleted flag error: %v", err)
	}
	if parsedEndDate != "" {
		res.EndDate = new(time.Time)
		*res.EndDate, err = time.Parse(time.RFC3339, parsedEndDate)
	}

	parsedUid, err := cmd.Flags().GetInt64("uid")
	if err != nil {
		return nil, fmt.Errorf("user id flag error: %v", err)
	}
	if parsedUid != 0 {
		res.User = new(int64)
		*res.User = parsedUid
	}

	pid, err := cmd.Flags().GetString("pid")
	if err != nil {
		return nil, fmt.Errorf("project id flag error: %v", err)
	}

	res.Project = pid

	return res, nil
}

func (h taskHandler) getTaskID(cmd *cobra.Command, args []string) (uuid.UUID, error) {
	var id uuid.UUID
	if len(args) == 0 {
		idFlag, err := cmd.Flags().GetString("uuid")
		if idFlag == "" {
			return uuid.UUID{}, fmt.Errorf("id not provided")
		}
		id, err = uuid.Parse(idFlag)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("uuid parsing failed: %v\n", err)
		}
	} else {
		idStr := args[0]
		idNum, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("id number parsing failed: %v\n", err)
		}
		id, err = h.storage.GetTaskID(int(idNum))
		if err != nil {
			return uuid.UUID{}, fmt.Errorf("id localstorage request failed: %v\n", err)
		}
	}
	return id, nil
}
