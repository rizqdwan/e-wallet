package middlerware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type New

func generateToken(id string, secret string)(tokenStr string, error){
	claims := jwt.MapClaims{
		"id" : id,
		"exp" : jwt.NumericDate(time.Now().Add(time.Hour * 24)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenStr, err = token.SigningString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nill 
}

func validateToken(tokenStr string) ()