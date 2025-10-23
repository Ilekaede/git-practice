package user

import (
	"github/git-practice/pkg/ulid"
)

type User struct {
	id             string
	email          Email
	name           string
	hashedPassword HashedPassword
}

func NewUser(
	email string,
	name string,
	password string,
) (*User, error) {
	validatedEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}
	HashedPassword, err := newHashedPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		id:             ulid.NewUlid(),
		email:          validatedEmail,
		name:           name,
		hashedPassword: HashedPassword,
	}, err
}

// インスタンスの再構成
func ReconstructUser(
	id string,
	email string,
	name string,
	hashedPassword string,
) *User {
	return &User{
		id:             id,
		email:          reconstructEmail(email),
		name:           name,
		hashedPassword: reconstructHashedPassword(hashedPassword),
	}
}

// ユーザーオブジェクト更新
func (u *User) UpdateUser(
	email string,
	name string,
) (*User, error) {
	validatedEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		id:             u.id,
		email:          validatedEmail,
		name:           name,
		hashedPassword: u.hashedPassword,
	}, nil

}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) ComparePassword(plainPassword string) error {

	// UserがHashedPasswordを操作している
	return u.hashedPassword.compare(plainPassword)
}
