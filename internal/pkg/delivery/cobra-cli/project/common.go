package project

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	domain "jirno/internal/pkg/domain/project"
	"time"
)

type projectType int8

const (
	typeProject       projectType = 0
	typeProjectUpdate projectType = 1
)

func parseProject(cmd *cobra.Command, args []string, t projectType) (*domain.DeliveryProject, error) {
	res := &domain.DeliveryProject{}

	if t == typeProjectUpdate {
		parsedId, err := cmd.Flags().GetString("id")
		if err != nil {
			return nil, fmt.Errorf("id flag error: %v", err)
		}
		res.ID = parsedId
	}

	parsedUid, err := cmd.Flags().GetInt64Slice("uids")
	if err != nil {
		return nil, fmt.Errorf("user id flag error: %v", err)
	}
	if len(parsedUid) != 0 {
		res.Users = parsedUid
	}

	parsedParentPid, err := cmd.Flags().GetString("ppid")
	if err != nil {
		return nil, fmt.Errorf("project id flag error: %v", err)
	}
	res.ParentProject = parsedParentPid

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

func parseProjectID(cmd *cobra.Command, args []string) (res uuid.UUID, err error) {
	if len(args) == 0 {
		err = fmt.Errorf("id not provided")
		return
	}
	id := args[0]
	res, err = uuid.Parse(id)
	if err != nil {
		err = fmt.Errorf("uuid parsing failed: %v\n", err)
		return
	}
	return
}

func parseFilter(cmd *cobra.Command, args []string) (*domain.SmartProjectFilter, error) {
	res := &domain.SmartProjectFilter{}

	if len(args) > 0 {
		if args[0] != "" {
			res.Smart = args[0]
		}
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

	parsedPid, err := cmd.Flags().GetString("ppid")
	if err != nil {
		return nil, fmt.Errorf("parnet project id flag error: %v", err)
	}
	res.ParentProject = parsedPid

	return res, nil
}
