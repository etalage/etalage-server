package repo

import "github.com/etalage/etalage-server/model"

type UserRepo struct {
}

func (repo UserRepo) AddUser(user model.User) int {
	return 1
}
