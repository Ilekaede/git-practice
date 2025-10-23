package user
import "context"

type UserRepository interface{
	FindByEmail(ctx context.Content, email Email) (*User, error)

}