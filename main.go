package main

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/vjeantet/eastertime"
	"log"
	"strconv"
)

func main() {
	app := fiber.New()

	app.All("/", func(c *fiber.Ctx) error {
		queryYear := c.Query("year", carbon.Now().Format("Y"))

		year, err := strconv.Atoi(queryYear)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, getJSONError("the year must be an integer"))
		}

		holidays, err := getHolidays(year)

		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, getJSONError(err.Error()))
		}

		holidaysJson, _ := json.Marshal(holidays)

		return c.Send(holidaysJson)
	})

	err := app.Listen(":8080")

	if err != nil {
		log.Fatal(err)
	}
}

func getHolidays(year int) ([15]Holiday, error) {
	orthodoxEaster, err := eastertime.OrthodoxByYear(year)
	if err != nil {
		return [15]Holiday{}, errors.New("the year must be greater than 325")
	}

	easter := carbon.Time2Carbon(orthodoxEaster)
	secondDayOfEaster := easter.AddDay()
	goodFriday := easter.SubDays(2)
	whitMonday := easter.AddDays(50)
	whitSunday := whitMonday.SubDay()

	holidays := [15]Holiday{
		{
			Name: "Anul nou",
			Date: carbon.CreateFromDate(year, 1, 1).ToDateString(),
		},
		{
			Name: "Anul nou",
			Date: carbon.CreateFromDate(year, 1, 2).ToDateString(),
		},
		{
			Name: "Ziua Unirii",
			Date: carbon.CreateFromDate(year, 1, 24).ToDateString(),
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
			Date: carbon.CreateFromDate(year, 5, 1).ToDateString(),
		},
		{
			Name: "Ziua Copilului",
			Date: carbon.CreateFromDate(year, 6, 1).ToDateString(),
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
			Date: carbon.CreateFromDate(year, 8, 15).ToDateString(),
		},
		{
			Name: "Ziua Sfântului Andrei",
			Date: carbon.CreateFromDate(year, 11, 30).ToDateString(),
		},
		{
			Name: "Ziua naţională",
			Date: carbon.CreateFromDate(year, 12, 1).ToDateString(),
		},
		{
			Name: "Crăciunul",
			Date: carbon.CreateFromDate(year, 12, 25).ToDateString(),
		},
		{
			Name: "A doua zi de Crăciun",
			Date: carbon.CreateFromDate(year, 12, 26).ToDateString(),
		},
	}

	return holidays, nil
}

func getJSONError(message string) string {
	err, _ := json.Marshal(Error{message})
	return string(err)
}

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type Error struct {
	Message string `json:"message"`
}
