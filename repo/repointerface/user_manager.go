package repointerface

import "github.com/etalage/etalage-server/model"

type UserManager interface {
	AddUser(user model.User) int
}
