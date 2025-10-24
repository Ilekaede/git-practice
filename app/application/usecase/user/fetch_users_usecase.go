package user
import(
	"context"
	"github.com/git-practice/app/domain/user"
)

type FetchUsersUsecase struct{
	userRepository user.UserRepository
}

func NewFetchUsersUsecase(
	userRepository user.UserRepository
) *FetchUsersUsecase{
	return & FetchUsersUsecase{
		userRepository: userRepository
	}
}

func (fuu *FetchUsersUsecase) Run(ctx context.Context) ([]*FetchUsersUsecaseOutoutDTO, error){
	us, err := fuu.userRepository.FetchAllUsers(ctx)
	if err != nil{
		return nil, err
	}
	outputs := make([]*FetchUserUsecaseOutoutDTO, 0, len(us))
	for _, u := range us{
		outputs = append(outputs, &FetchUserUsecaseOutoutDTO{ID: u.GetID(), Name: u.GetName()})
	}
	return outputs, nil
}
