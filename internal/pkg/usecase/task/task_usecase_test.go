package task

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"jirno/internal/pkg/domain/task"
	"jirno/internal/pkg/repository/task/mock"
	"jirno/internal/pkg/utils"
	"jirno/internal/smart_parser"
	"testing"
	"time"
)

type getByFilterTest struct {
	expectedError  error
	expectedRes    []task.Task
	expectedFilter task.TaskFilter
	filter         task.DeliveryTaskFilter
}

var (
	now                  = time.Now()
	todayStart, todayEnd = utils.GetDayRange(time.Now())
	testUserID           = int64(3)

	resArr = []task.Task{
		{
			Title:       "Task1",
			Description: "Task1 description",
			DateTo:      &now,
			User:        testUserID,
		},
		{
			Title:       "Task2",
			Description: "Task2 description",
			DateTo:      &now,
			User:        testUserID,
		},
	}

	getByFilterData = []getByFilterTest{

		{
			expectedError: nil,
			expectedRes:   resArr,
			filter: task.DeliveryTaskFilter{
				Smart: "today",
			},
			expectedFilter: task.TaskFilter{
				StartDate: &todayStart,
				EndDate:   &todayEnd,
			},
		},

		{
			expectedError: nil,
			expectedRes:   resArr,
			filter: task.DeliveryTaskFilter{
				StartDate: &todayStart,
				EndDate:   &todayEnd,
			},
			expectedFilter: task.TaskFilter{
				StartDate: &todayStart,
				EndDate:   &todayEnd,
			},
		},

		{
			expectedError: nil,
			expectedRes:   resArr,
			filter: task.DeliveryTaskFilter{
				User: &testUserID,
			},
			expectedFilter: task.TaskFilter{
				User: &testUserID,
			},
		},
	}
)

func TestTaskUsecase_GetByFilter(t *testing.T) {
	ctrl := gomock.NewController(t)
	taskRepo := mock.NewMockITaskRepo(ctrl)
	usecase := NewTaskUsecase(taskRepo, smart_parser.NewDefaultTaskParser())

	for index, testData := range getByFilterData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().GetByFilter(testData.expectedFilter).Return(testData.expectedRes, nil).Times(1)
			res, err := usecase.GetByFilter(testData.filter)
			assert.Nil(t, err)
			assert.Equal(t, testData.expectedRes, res)
		})
	}
}

type completeTest struct {
	expectedError  error
	projectID      uuid.UUID
	expectedUpdate task.TaskUpdate
}

var (
	completeTaskID      = uuid.New()
	completeIsCompleted = true
	completeData        = []completeTest{
		{
			expectedError: nil,
			projectID:     completeTaskID,
			expectedUpdate: task.TaskUpdate{
				ID:          completeTaskID,
				IsCompleted: &completeIsCompleted,
			}},
	}
)

func TestTaskUsecase_Complete(t *testing.T) {
	ctrl := gomock.NewController(t)
	taskRepo := mock.NewMockITaskRepo(ctrl)
	usecase := NewTaskUsecase(taskRepo, smart_parser.NewDefaultTaskParser())
	for index, testData := range completeData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().Update(gomock.Any()).Return(testData.expectedError).Times(1)
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
	deleteTaskID = uuid.New()
	deleteData   = []deleteTest{
		{
			projectID:     deleteTaskID,
			expectedError: nil,
		},
	}
)

func TestTaskUsecase_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	taskRepo := mock.NewMockITaskRepo(ctrl)
	usecase := NewTaskUsecase(taskRepo, smart_parser.NewDefaultTaskParser())
	for index, testData := range deleteData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().Delete(testData.projectID).Return(testData.expectedError).Times(1)
			err := usecase.Delete(testData.projectID)
			assert.Equal(t, testData.expectedError, err)
		})
	}
}

type createTest struct {
	expectedError error
	project       task.Task
}

var (
	createTask = task.Task{
		Title:       "Task",
		Description: "Task description",
	}
	createData = []createTest{
		{
			expectedError: nil,
			project:       createTask,
		},
	}
)

func TestTaskUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	taskRepo := mock.NewMockITaskRepo(ctrl)
	usecase := NewTaskUsecase(taskRepo, smart_parser.NewDefaultTaskParser())
	for index, testData := range createData {
		t.Run(fmt.Sprintf("#%v", index), func(t *testing.T) {
			taskRepo.EXPECT().Create(gomock.Any()).Return(testData.expectedError).Times(1)
			_, err := usecase.Create(testData.project)
			assert.Equal(t, err, testData.expectedError)
		})
	}
}
