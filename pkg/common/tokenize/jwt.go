package tokenize

import (
	"fmt"
	"shortlink/pkg/common/config"
	"shortlink/pkg/common/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

func GenereateToken(userId, name string) (token string, err error) {
	mySigningKey := []byte(config.GetEnv("JWT_TOKEN"))

	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err = t.SignedString(mySigningKey)
	if err != nil {
		fmt.Println(err)
		err = fiber.ErrInternalServerError
		return
	}

	return
}

func GetUserId(DB *gorm.DB, tokenString string) (userID string, err error) {
	mySigningKey := []byte(config.GetEnv("JWT_TOKEN"))
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		fmt.Println(err)
	}

	userID = fmt.Sprintf("%v", claims["id"])
	if err = getUserId(DB,userID); err != nil {
		err = fiber.NewError(fiber.StatusUnauthorized, "Invalid or expired")
		return
	}
	return
}

func getUserId(DB *gorm.DB, userId string) (err error) {
	var user models.User

	err = DB.Where("id = ?", userId).First(&user).Error

	return
}