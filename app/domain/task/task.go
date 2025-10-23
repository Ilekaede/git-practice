package task
import(
	"github.com/git-practice/app/domain/errors"
	"github.com/git-practice/pkg/ulid"
)

type Task struct{
	id string
	userId string
	content Content
	state State
}

func NewTask(
	userId string,
	content string,
	state string,
)(*Task, error){
	validationContent , err := newContent(content)
	if err != nil{
		return nil, err
	}
	validationState, err := newState(state)
	if err != nil{
		return nil, err
	}
	return &Task{
		id: ulid.NewUlid(),
		userId: userId,
		content: validationContent,
		state: validationState
	}, nil
}

func ReconstructTask(
	id string,
	userId string,
	content string,
	state int,
) *Task{
	return &Task(
		id: id,
		userId: userId,
		content: reconstructContent(content),
		state: reconstructState(state),
	)
}

func (t *Task) UpdateState(
	state string
)(*Task, error){
	validationState, err := newState(state)
	if err != nil{
		return nil, err
	}
	return &Task{
		id: t.id,
		userId: t.userId,
		content: t.content,
		state: validationState,
	}, nil
}

// タスクに対してオペレーションをできるユーザの限定
func (t *Task) IsOperableBy(userId string) error {
	if t.userId != userId{
		return errors.ErrForbiddenTaskOperation
	}
	return nil
}

func (t *Task) GetID() string{
	return t.id
}
func (t *Task) GetUserId() string{
	return t.userId
}
func (t *Task) GetContent() Content{
	return t.content
}
func (t *Task) GetState() State{
	return t.state
}