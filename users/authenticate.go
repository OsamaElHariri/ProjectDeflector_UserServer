package users

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func issueJwt(user User) (string, error) {
	seconds := 60 * 60 * 5
	claims := jwt.StandardClaims{
		Subject:   user.Id,
		ExpiresAt: time.Now().UTC().Unix() + int64(seconds),
		Issuer:    "projectdeflector",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encypted, err := token.SignedString([]byte("SuperSecretText!!"))
	if err != nil {
		return "", nil
	}

	return encypted, nil
}

type AccessInfo struct {
	UserId string
}

func validateJwt(tokenToValidate string) (AccessInfo, error) {
	token, err := jwt.ParseWithClaims(
		tokenToValidate,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("SuperSecretText!!"), nil
		},
	)
	if err != nil {
		return AccessInfo{}, err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return AccessInfo{}, errors.New("invalid token")
	}

	if !claims.VerifyExpiresAt(time.Now().UTC().Unix(), true) {
		return AccessInfo{}, errors.New("invalid token")
	}

	info := AccessInfo{
		UserId: claims.Subject,
	}

	return info, nil
}
