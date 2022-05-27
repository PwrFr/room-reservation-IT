package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// env, _ := godotenv.Read()
//generate jwt token
func InitEnv() map[string]string {
	_ = godotenv.Load()
	env, _ := godotenv.Read()

	return env
}

func GenerateJwtToken(role string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(InitEnv()["SECRET_KEY"]))
	if err != nil {
		return "", err
	}

	return t, nil

}

func verifyToken(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(InitEnv()["SECRET_KEY"]), nil
	})
	if err != nil {
		return false
	}

	return token.Valid
}

func extractToken(tokenStr string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(InitEnv()["SECRET_KEY"]), nil
		})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims, nil
}
