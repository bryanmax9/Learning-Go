package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getSecretKey() []byte {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

	
	return []byte(secretKey)
}

func GenerateToken(email string, userId int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(getSecretKey())
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		
		//check if token was signed
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return getSecretKey(), nil
	})

	if err != nil {
		return 0,errors.New("could not parse token")
	}

	//check if its a valid token signed by our secret key
	if !parsedToken.Valid {
		return 0,errors.New("invalid token")
	}

	//it should be true this below statement if it was signed by our serer key
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0,errors.New("invalid token claims")
	}

	// //if its a valid statement than we are doing type hecking of the email and user id of the right format
	// email := claims["email"].(string)


	userId := int64(claims["userId"].(float64))

	return  userId, nil
}