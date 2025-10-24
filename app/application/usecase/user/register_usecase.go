package user
import(
	"context"
	"github.com/git-practice/app/domain/errors"
	"github.com/git-practice/app/domain/user"
)

type RegisterUsecase struct{
	userRepository user.UserRepository
	userDomainService user.UserDomainService
}

func NewRegisterUsecase(
	userRepository user.UserRepository,
	userDomainService user.UserDomainService,
) *RegisterUsecase{
	return &RegisterUsecase{
		userRepository: userRepository,
		userDomainService: userDomainService,
	}
}

func (ru *RegisterUsecase) Run(ctx context.Context, input RegisterUsecaseInputDTO) (*RegisterUsecaseOutputDTO, error){
	// userインスタンスを生成
	u, err := user.NewUser(
		input.Email,
		input.Name,
		input.Password,
	)
	if err != nil{
		return nil, err
	}
	// userが既に登録済みか？
	ok, err := ru.userDomainService.IsExists(ctx, u.GetEmail())
	if err != nil{
		return nil, err
	}
	if ok{
		return nil, errors.ErrAlreadyRegistered
	}

	// userが保存されたときにエラーが出るか？のチェック
	if err := ru.userRepository.Save(ctx, u); err != nil{
		return nil, err
	}
	return &RegisterUsecaseOutputDTO{
		ID: u.GetID(),
		Name: u.GetName(),
		Email: u.GetEmail().Value(),
	},nil
}