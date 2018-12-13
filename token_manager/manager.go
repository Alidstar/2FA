package tokenmanager

import (
	"log"

	"github.com/xlzd/gotp"
)

type TokenManager struct {
	signature string
	users     map[string]*Detail
}

type Detail struct {
	token       string
	allowedApps map[string]bool
}

func NewTokenManager() *TokenManager {
	return &TokenManager{
		gotp.RandomSecret(32),
		make(map[string]*Detail),
	}
}

func (manager *TokenManager) GenerateToken(username string) string {
	if _, ok := manager.users[username]; !ok {
		manager.users[username] = &Detail{
			"",
			map[string]bool{
				"todo":     true,
				"calendar": true,
				"contact":  true,
			},
		}
	}

	token := generateToken(username, manager.signature)
	log.Println("token", token)

	manager.users[username].token = token

	return token
}

func (manager *TokenManager) Verify(app, token string) (bool, string) {
	var username string
	if result, username := verifyToken(token, manager.signature); !result {
		log.Println("Invalid token", token)
		return result, username
	}

	if manager.users[username].token != token {
		log.Println("Revoked token", token)
		return false, ""
	}

	if !manager.users[username].allowedApps[app] {
		log.Println("App not allowed", app)
		return false, ""
	}

	return true, username
}

func (manager *TokenManager) RevokeToken(username string) {
	manager.users[username].token = ""
}
