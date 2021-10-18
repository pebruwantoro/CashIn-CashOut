package middlewares

import (
	"cash-app/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(user_id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = int(user_id)
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constant.SECRET_JWT))

}

func ExtractToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		user_id := int(claims["user_id"].(float64))
		return user_id
	}
	return 0
}
