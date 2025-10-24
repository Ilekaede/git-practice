package task
import(
	"context"
	"github.com/git-practice/app/domain/task"
)

type CreateTaskUsecase struct{
	taskRepository task.TaskRepository
}

func NewCreateTaskUsecase(taskRepository task.TaskRepository) *CreateTaskUsecase{
	return &CreateTaskUsecase{
		taskRepository: taskRepository,
	}
}

func (ctu *CreateTaskUsecase) Run(ctx context.Context, input CreateTaskUsecaseInputDTO){
	t, err := task.NewTask(input.UserId, input.Content, input.State)
	if err != nil{
		return nil, err
	}
	if err != nil{
		return nil, err
	}
	if err := ctu.taskRepository.Sace(ctx, t); err != nil{
		return nil, err
	}
	return &CreateTaskUsecaseOutputDTO{
		ID: t.GetID(),
		UserId: t.GetUserId(),
		Content: t.GetContent().Value(),
		State: t.GetState().StrValue(),
	}, nil
}