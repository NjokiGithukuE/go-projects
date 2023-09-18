package main

import (
	"errors"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret = []byte("your-secret-key")
	users = []User{
		{ID: 1, Username: "user1", Password: "password1", Role: "user"}, 
		{ID: 2, Username: "admin", Password: "adminpass", Role: "admin"}, 
	}
)

func generateJWT(user User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.Username, 
		"iss": "your-app-name", 
		"iat": time.Now().Unix(), 
		"exp": time.Now().Add(time.Hour * 24).Unix(), 
		"role": user.Role, 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func authenticate(username, password string) (User, error) {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return user, nil
		}
	}
	return User{}, errors.New("authentication failed")
}