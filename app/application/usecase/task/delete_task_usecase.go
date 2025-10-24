package user
import(
	"context"
	"github.com/git-practice/app/domain/task"
)

type DeleteTaskUsecase struct{
	taskRepository task.TaskRepository
}

func NewDeleteTaskRepository(taskRepository task.TaskRepository,) *DeleteTaskRepository{
	return &deleteTaskRepository{
		taskRepository: taskRepository,
	}
}

func (dtu *DeleteTaskUsecase) Run(ctx context.Context, input DeleteTaskUsecaseInputDTO) error{
	// 存在しているユーザーしか削除できない
	t, err := dtu.taskRepository.FindById(ctx, input.ID)
	if err != nil || t == nil{
		return err
	}
	// ログインしてるユーザIDと一致したら
	if err := t.IsOperableBy(input.UserId); err != nil{
		return err
	}
	if err := dtu.taskRepository.Delete(ctx, t); err != nil{
		return err
	}
	return nil
}

