package main

import (
	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	JariJariLingkaran int 	    `json:"jari-jari-lingkaran"`
}

type ResponseData struct {
	LuasLingkaran 		float64     `json:"luas-lingkaran"`
	KelilingLingkaran   float64		`json:"keliling-lingkaran"`
}

func main() {
	app := fiber.New()

	app.Post("/", func (c *fiber.Ctx) error {
		request := new(RequestData)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		response := new(ResponseData)

		response.LuasLingkaran = 3.14 * float64(request.JariJariLingkaran * request.JariJariLingkaran)
		response.KelilingLingkaran = 2 * float64(request.JariJariLingkaran) * 3.14

		return c.JSON(response) 
	})
	
	app.Listen(":3000")
}