package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/vjeantet/eastertime"
)

type holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func main() {
	app := fiber.New(fiber.Config{
		GETOnly: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(map[string]string{
				"message": err.Error(),
			})
		},
	})

	app.Get("/", func(c *fiber.Ctx) error {
		queryYear := c.Query("year", carbon.Now().Format("Y"))

		year, err := strconv.Atoi(queryYear)

		if err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "the year must be an integer")
		}

		holidays, err := getHolidays(year)

		if err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
		}

		return c.JSON(holidays)
	})

	err := app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func getHolidays(y int) ([]holiday, error) {
	orthodoxEaster, err := eastertime.OrthodoxByYear(y)
	if err != nil {
		return []holiday{}, errors.New("the year must be greater than 325")
	}

	easter := carbon.CreateFromStdTime(orthodoxEaster)
	secondDayOfEaster := easter.AddDay()
	goodFriday := easter.SubDays(2)

	whitMonday := easter.AddDays(50)
	whitSunday := whitMonday.SubDay()

	return []holiday{
		{
			Name: "Anul nou",
			Date: carbon.CreateFromDate(y, 1, 1).ToDateString(),
		},
		{
			Name: "Anul nou",
			Date: carbon.CreateFromDate(y, 1, 2).ToDateString(),
		},
		{
			Name: "Bobotează",
			Date: carbon.CreateFromDate(y, 1, 6).ToDateString(),
		},
		{
			Name: "Soborul Sfântului Ioan Botezătorul",
			Date: carbon.CreateFromDate(y, 1, 7).ToDateString(),
		},
		{
			Name: "Ziua Unirii",
			Date: carbon.CreateFromDate(y, 1, 24).ToDateString(),
		},
		{
			Name: "Vinerea Mare",
			Date: goodFriday.ToDateString(),
		},
		{
			Name: "Paștele",
			Date: easter.ToDateString(),
		},
		{
			Name: "A doua zi de Paște",
			Date: secondDayOfEaster.ToDateString(),
		},
		{
			Name: "Ziua Muncii",
			Date: carbon.CreateFromDate(y, 5, 1).ToDateString(),
		},
		{
			Name: "Ziua Copilului",
			Date: carbon.CreateFromDate(y, 6, 1).ToDateString(),
		},
		{
			Name: "Rusalii",
			Date: whitSunday.ToDateString(),
		},
		{
			Name: "A doua zi de Rusalii",
			Date: whitMonday.ToDateString(),
		},
		{
			Name: "Adormirea Maicii Domnului",
			Date: carbon.CreateFromDate(y, 8, 15).ToDateString(),
		},
		{
			Name: "Ziua Sfântului Andrei",
			Date: carbon.CreateFromDate(y, 11, 30).ToDateString(),
		},
		{
			Name: "Ziua naţională",
			Date: carbon.CreateFromDate(y, 12, 1).ToDateString(),
		},
		{
			Name: "Crăciunul",
			Date: carbon.CreateFromDate(y, 12, 25).ToDateString(),
		},
		{
			Name: "A doua zi de Crăciun",
			Date: carbon.CreateFromDate(y, 12, 26).ToDateString(),
		},
	}, nil
}
