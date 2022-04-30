package middleware

import (
	"github.com/golang-jwt/jwt"
	"mountainio/domain/model"
	"os"
)

func GenerateToken(auth model.AuthClaim) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = auth.UserID
	claims["role"] = auth.Role
	claims["expired"] = auth.Expired

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
