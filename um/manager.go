package um

import (
	"github.com/xlzd/gotp"
)

type UserManager struct {
	signature string
	data      map[string]Detail
}

type Detail struct {
	secret  string
	counter int
	lineID  string
}

func NewUserManager() *UserManager {
	return &UserManager{
		gotp.RandomSecret(32),
		make(map[string]Detail),
	}
}

func (um *UserManager) Register(username string) {
	um.data[username] = Detail{
		gotp.RandomSecret(32),
		0,
		"",
	}
}

func (um *UserManager) Login(username string) string {
	return um.generateToken(username)
}

func (um *UserManager) Verify(token string) (bool, string) {
	return um.verifyToken(token)
}
