package link

import (
	"fmt"
	"os"

	"shortlink/pkg/common/config"

	"github.com/gofiber/fiber/v2"
	"github.com/skip2/go-qrcode"
)

var (
	BASE_URL = config.GetEnv("BASE_URL")
)

func CreateQR(qr, slug string) (err error){
	if err = qrcode.WriteFile(BASE_URL+slug, qrcode.Medium, 256, "public/"+qr+".png"); err != nil {
		fmt.Println(err)
		fmt.Println("Hi Im here create")
		return fiber.ErrInternalServerError
	}

	return
}

func DeleteQR(qr string) (err error) {
	if e := os.Remove("public/"+qr+".png");e != nil {
		fmt.Println(err)
		fmt.Println("Hi Im here delete")
		return fiber.ErrInternalServerError
	}

	return
}