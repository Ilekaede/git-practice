package task
import (
	"context"
	"testing"
	"github.com/google/go-cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/git-practice/app/domain/errors"
	"github.com/git-practice/app/domain/task"
	"go.uber.org/mock/gomock"
)

func TestTask_CreateTaskUsecase_Run(t *testing.T){
	t.Parallel()
	tests := []struct{
		name string
		mockFn func(mr *task.MockTaskRepository)
		input CreateTaskUsecaseInputDTO
		errType error
		want *CreateTaskUsecaseOutputDTO
		wantErr bool
	}{
		{
			name:"正常系",
			mockFn: func(mr *task.MockTaskRepository){
				mr.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil)
			},
			input: CreateTaskUsecaseInputDTO{
				UserId: "user_id",
				Content: "this is content",
				State: "todo",
			},
			want: &CreateTaskUsecaseOutputDTO{
				UserId: "user_id",
				Content: "this is content",
				State: "todo",
			},
			wantErr: false,
		},
		{
			name: "準正常系: タスク内容が空っぽ",
			mockFn:func(mr *task.MockTaskRepository){
				//saveメソッドは呼ばれない
			},
			input: CreateTaskUsecaseInputDTO{
				userId:"user_id",
				Content: "",
				State: "todo",
			},
			errType:errors.ErrContentEmpty,
			wantErr: true,
		},
	}
	for _, tt := range tests{
		tt := tt
		t.Run(tt.name, func(t *testing.T){
			t.Parallel()
			ctrl := gomock.NewController(t)
			MockTaskRepository := task.NewMockTaskRepository(ctrl)
			tt.mockFn(mockTaskRepository)

			sut := NewCreateTaskUsecase(mockTaskRepository)
			ctx := context.Background()
			got, err := sut.Run(ctx, tt.input)

			if tt.errType != nil && !errors.Is(err, tt.errType){
				t.Errorf("createTaskUsecase.Run = error%v,want errType:%v", err, tt.errType)
				return
			}
			if (err != nil) != tt.wantErr{
				t.Errorf("createTaskUsecase.Run = error:%v, wantErr:%v", err, tt.wantErr)
				return
			}
		})
	}
}