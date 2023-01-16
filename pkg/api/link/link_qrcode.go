package link

import (
	"fmt"
	"shortlink/pkg/common/config"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

var (
	BASE_URL = config.GetEnv("BASE_URL")
)

func CreateQR(id, slug string) (err error){
	if err = qrcode.WriteFile(BASE_URL+slug, qrcode.Medium, 256, "public/"+id+".png"); err != nil {
		fmt.Println(err)
		return fiber.ErrInternalServerError
	}

	return
}