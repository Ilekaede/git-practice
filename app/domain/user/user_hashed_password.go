package user

import (
	"unicode/utf8"

	"github.com/git-practice/app/domain/errors"
	"github.com/git-practice/pkg/hash"
)

type HashedPassword struct {
	value string
}

const (
	minPasswordLength = 6
)

func newHashedPassword(value string) (HashedPassword, error){
	if minPasswordLength >= utf8.RuneCountInString(value){
		return HashedPassword{}, errors.ErrPasswordTooShort
	}
	hashed, err := hash.Hash(value)
	if err != nil{
		return HashedPassword{}, err
	}

	return HashedPassword{value: hashed}, nil
}

func (p HashedPassword) compare(target string) error{
	if err := hash.Compare(p.value, target); err != nil{
		return errors.ErrPasswordMismatch
	}
	return nil
}