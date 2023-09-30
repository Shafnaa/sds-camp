package main

import (
	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	JariJariLingkaran int 	    `json:"jari-jari-lingkaran"`
	SisiPersegi       int       `json:"sisi-persegi"`
	AlasSegitiga      int       `json:"alas-segitiga"`
	TinggiSegitiga    int       `json:"tinggi-segitiga"`
}

type ResponseData struct {
	LuasLingkaran 		float64     `json:"luas-lingkaran"`
	LuasPersegi 		int 		`json:"luas-persegi"`
	LuasSegitiga 		float64 	`json:"luas-segitiga"`
	KelilingLingkaran   float64		`json:"keliling-lingkaran"`
	KelilingPersegi 	int 		`json:"keliling-persegi"`
	KelilingSegitiga 	int 		`json:"keliling-segitiga"`
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
		response.LuasPersegi = request.SisiPersegi * request.SisiPersegi 
		response.LuasSegitiga = float64(request.AlasSegitiga * request.TinggiSegitiga) / 2
		response.KelilingLingkaran = 2 * float64(request.JariJariLingkaran) * 3.14
		response.KelilingPersegi = 4 * request.SisiPersegi
		response.KelilingSegitiga = 3 * request.AlasSegitiga

		return c.JSON(response) 
	})
	
	app.Listen(":3000")
}