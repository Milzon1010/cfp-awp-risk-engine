package main

import (
	"cfp-engine/internal/calculator"
	"cfp-engine/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi koneksi ke Database PostgreSQL
	database.ConnectDB()

	app := fiber.New()

	// Menyajikan tampilan antarmuka web HTML dari folder "static"
	app.Static("/", "./static")

	// Endpoint POST untuk menerima data kustom dari frontend / klien
	app.Post("/api/calculate-risk", func(c *fiber.Ctx) error {
		var input calculator.RiskProfileInput

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": "error",
				"error":  "Format input data tidak valid",
			})
		}

		result := calculator.CalculateRiskAndEmergencyFund(input)

		return c.JSON(fiber.Map{
			"status": "success",
			"data":   result,
		})
	})

	// Endpoint GET khusus untuk pengujian data simulasi default
	app.Get("/api/test-simulate", func(c *fiber.Ctx) error {
		sampleInput := calculator.RiskProfileInput{
			MonthlyExpense:    15000000,
			IsMarried:         true,
			NumberOfChildren:  2,
			ExistingInsurance: 500000000,
			AnnualIncome:      300000000,
		}

		result := calculator.CalculateRiskAndEmergencyFund(sampleInput)

		return c.JSON(fiber.Map{
			"status":      "success",
			"description": "Simulasi data default",
			"data":        result,
		})
	})

	// Menjalankan server pada port 8080
	app.Listen(":8080")
}
