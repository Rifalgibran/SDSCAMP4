package main

import (
	"math"

	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	JariJariLingkaran float64 `json:"jari-jari-lingkaran"`
	SisiPersegi       float64 `json:"sisi-persegi"`
	AlasSegitiga      float64 `json:"alas-segitiga"`
	TinggiSegitiga    float64 `json:"tinggi-segitiga"`
}

type ResponseData struct {
	LuasLingkaran     float64 `json:"luas-Lingkaran"`
	LuasPersegi       float64 `json:"luas-Persegi"`
	LuasSegitiga      float64 `json:"luas-Segitiga"`
	KelilingLingkaran float64 `json:"keliling-Lingkaran"`
	KelilingPersegi   float64 `json:"keliling-Persegi"`
	KelilingSegitiga  float64 `json:"keliling-Segitiga"`
}

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		request := new(RequestData)

		if err := c.BodyParser(request); err != nil {
			return err
		}

		response := ResponseData{}

		response.LuasLingkaran = math.Pi * request.JariJariLingkaran * request.JariJariLingkaran
		response.KelilingLingkaran = 2 * math.Pi * request.JariJariLingkaran
		response.LuasPersegi = math.Pow(request.SisiPersegi, 2)
		response.LuasSegitiga = 0.5 * request.AlasSegitiga * request.TinggiSegitiga
		response.KelilingPersegi = 4 * request.SisiPersegi
		response.KelilingSegitiga = request.AlasSegitiga + 2*math.Sqrt(math.Pow(request.TinggiSegitiga, 2)+math.Pow(request.AlasSegitiga/2, 2))

		return c.JSON(response)
	})

	app.Listen(":3000")
}
