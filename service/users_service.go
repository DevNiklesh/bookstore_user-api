package service

import (
	"github.com/DevNiklesh/bookstore_user-api/domain/users"
	"github.com/DevNiklesh/bookstore_user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}
	if err := result.GetById(); err != nil {
		return nil, err
	}

	return result, nil
}
