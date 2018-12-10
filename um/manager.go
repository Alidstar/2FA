package um

import (
	"time"

	"github.com/xlzd/gotp"
)

type UserManager struct {
	data map[string]Detail
}

type Detail struct {
	secret  string
	counter int
	lineID  string
}

func NewUserManager() *UserManager {
	return &UserManager{
		make(map[string]Detail),
	}
}

func (um *UserManager) Register(username string) {
	um.data[username] = Detail{
		gotp.RandomSecret(64),
		0,
		"",
	}
}

func (um *UserManager) GetHOTP(username string) string {
	detail := um.data[username]
	hotp := gotp.NewDefaultHOTP(detail.secret)
	counter := detail.counter
	detail.counter += 1
	um.data[username] = detail
	return hotp.At(counter)
}

func (um *UserManager) VerifyHOTP(username, otp string) bool {
	hotp := gotp.NewDefaultHOTP(um.data[username].secret)
	return hotp.Verify(otp, um.data[username].counter-1)
}

func (um *UserManager) GetTOTP(username string) string {
	return gotp.NewDefaultTOTP(um.data[username].secret).Now()
}

func (um *UserManager) VerifyTOTP(username, otp string) bool {
	totp := gotp.NewDefaultTOTP(um.data[username].secret)
	return totp.Verify(otp, int(time.Now().Unix()))
}
