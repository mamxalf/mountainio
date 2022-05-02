package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"mountainio/domain/model"
	"os"
)

// TODO: Ini harusnya tidak masuk middleware
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

// TODO: Ini harusnya tidak masuk middleware
func GetClaimToken(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)
}

func AuthProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("SECRET_JWT")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
	}
}
