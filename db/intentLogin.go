package db

import (
	"github.com/jorgealbertogomezgomez77/social-app-like-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntentLogin check login in database */
func IntentLogin(email string, password string) (models.User, bool) {
	usu, founded, _ := CheckUserExist(email)
	if founded == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}