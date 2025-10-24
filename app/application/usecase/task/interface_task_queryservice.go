package task
import "context"

type TaskQueryService interface{
	FetchTaskById(ctx context.Context, id string)(*FetchTaskDTO, error)
	FetchUserTasks(crx context.Context, userId string)([]*FetchTaskDTO, error)
	FetchAllTasks(ctx context.Context) ([]*FetchTaskDTO, error)
}

type FetchTaskDTO struct{
	ID string
	UserName string
	UserId string
	Content string
	State string
}