package server

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func SetupAndListen() {
	router := fiber.New()

	fmt.Println(router)

}
