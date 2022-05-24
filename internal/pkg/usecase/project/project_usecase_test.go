package project

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"jirno/internal/pkg/domain/project"
	"jirno/internal/pkg/repository/project/mock"
	"jirno/internal/pkg/utils"
	"jirno/internal/smart_parser"
	"testing"
	"time"
)

type getByFilterTest struct {
	expectedError  error
	expectedRes    []project.Project
	expectedFilter project.ProjectFilter
	filter         project.SmartProjectFilter
}

var (
	now                  = time.Now()
	todayStart, todayEnd = utils.GetDayRange(time.Now())
	testUserID           = int64(3)
	resArr               = []project.Project{
		{
			Title:       "Project1",
			Description: "Project1 description",
			DateTo:      &now,
			Users:       []int64{1, 2, 3},
		},
		{
			Title:       "Project2",
			Description: "Project2 description",
			DateTo:      &now,
			Users:       []int64{1, 2, 3},
		},
	}

	getByFilterData = []getByFilterTest{
		{
			expectedError: nil,
			expectedRes:   resArr,
			filter: project.SmartProjectFilter{
				Smart: "today",
			},
			expectedFilter: project.ProjectFilter{
				StartDate: &todayStart,
				EndDate:   &todayEnd,
			},
		},
		{
			expectedError: nil,
			expectedRes:   resArr,
			filter: project.SmartProjectFilter{
				StartDate: &todayStart,
				EndDate:   &todayEnd,
			},
			expectedFilter: project.ProjectFilter{
				StartDate: &todayStart,
				EndDate:   &todayEnd,
			},
		},
		{
			expectedError: nil,
			expectedRes:   resArr,
			filter: project.SmartProjectFilter{
				User: &testUserID,
			},
			expectedFilter: project.ProjectFilter{
				User: &testUserID,
			},
		},
	}
)

func TestProjectUsecase_GetByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	projectRepo := mock.NewMockIProjectRepo(ctrl)
	usecase := NewProjectUsecase(projectRepo, smart_parser.NewDefaultProjectParser())
	for index, testData := range getByFilterData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			projectRepo.EXPECT().GetByFilter(testData.expectedFilter).Return(testData.expectedRes, nil).Times(1)
			res, err := usecase.GetByFilter(testData.filter)
			assert.Nil(t, err)
			assert.Equal(t, testData.expectedRes, res)
		})
	}
}

type completeTest struct {
	expectedError  error
	projectID      uuid.UUID
	expectedUpdate project.ProjectUpdate
}

var (
	completeProjectID   = uuid.New()
	completeIsCompleted = true
	completeData        = []completeTest{
		{
			expectedError: nil,
			projectID:     completeProjectID,
			expectedUpdate: project.ProjectUpdate{
				ID:          completeProjectID,
				IsCompleted: &completeIsCompleted,
			}},
	}
)

func TestProjectUsecase_Complete(t *testing.T) {
	ctrl := gomock.NewController(t)
	projectRepo := mock.NewMockIProjectRepo(ctrl)
	usecase := NewProjectUsecase(projectRepo, smart_parser.NewDefaultProjectParser())
	for index, testData := range completeData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			projectRepo.EXPECT().Update(testData.expectedUpdate).Return(testData.expectedError).Times(1)
			err := usecase.Complete(testData.projectID)
			assert.Equal(t, testData.expectedError, err)
		})
	}
}

type deleteTest struct {
	expectedError error
	projectID     uuid.UUID
}

var (
	deleteProjectID = uuid.New()
	deleteData      = []deleteTest{
		{
			projectID:     deleteProjectID,
			expectedError: nil,
		},
	}
)

func TestProjectUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	projectRepo := mock.NewMockIProjectRepo(ctrl)
	usecase := NewProjectUsecase(projectRepo, smart_parser.NewDefaultProjectParser())
	for index, testData := range deleteData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			projectRepo.EXPECT().Delete(testData.projectID).Return(testData.expectedError).Times(1)
			err := usecase.Delete(testData.projectID)
			assert.Equal(t, testData.expectedError, err)
		})
	}
}

type createTest struct {
	expectedError error
	project       project.Project
}

var (
	createProject = project.Project{
		Title:       "Project",
		Description: "Project description",
	}
	createData = []createTest{
		{
			expectedError: nil,
			project:       createProject,
		},
	}
)

func TestProjectUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	projectRepo := mock.NewMockIProjectRepo(ctrl)
	usecase := NewProjectUsecase(projectRepo, smart_parser.NewDefaultProjectParser())
	for index, testData := range createData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			projectRepo.EXPECT().Create(gomock.Any()).Return(testData.expectedError).Times(1)
			_, err := usecase.Create(testData.project)
			assert.Equal(t, err, testData.expectedError)
		})
	}
}
