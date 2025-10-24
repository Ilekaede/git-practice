package user
import (
	"context"
	"github.com/git-practice/app/domain/user"
)

type UpdateProfileUsecase struct{
	userRepository user.UserRepository
}

func NewUpdateProfileUsecase(
	userRepository user.UserRepository,
) *UpdateProfileUsecase{
	return &UpdateProfileUsecase{
		userRepository: userRepository,
	}
}

func (epu *UpdateProfileUsecase) Run(ctx context.Context, input UpdateProfileUsecaseInputDTO)(
	*UpdateProfileUsecaseOutputDTO, error,
){
	// 存在してるユーザーしか編集できない
	u, err := epu.userRepository.FindById(ctx, input.ID)
	if err != nil || u == nil{
		return nil, err
	}
	if input.Name == ""{
		input.Name = u.GetName()
	}
	if input.Email == ""{
		input.Email = u.GetEmail().Value()
	}
	// input情報をもとに更新情報を反映したインスタンスを作成
	updateUser, err := u.UpdateUser(
		input.Email,
		input.Name,
	)
	if err != nil{
		return nil, err
	}
	// 更新したオブジェクトをDTOに詰め替える
	return &UpdateProfileUsecaseOutputDTO{
		ID: updateUser.GetUD(),
		Email: updateUser.GetEmail().value(),
		Name: updateUser.GetName(),
	}, nil

}
