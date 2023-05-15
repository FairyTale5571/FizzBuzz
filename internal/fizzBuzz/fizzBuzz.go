package fizzBuzz

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Provider interface {
	IsFuzzBuzz() fiber.Handler
}

type fizzBuzz struct {
}

func New() Provider {
	return &fizzBuzz{}
}

func (f *fizzBuzz) IsFuzzBuzz() fiber.Handler {
	return func(c *fiber.Ctx) error {
		valueStr := c.Params("number")
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"result": "invalid number",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": verifyFizzBuzzInternal(value),
		})
	}
}

func verifyFizzBuzzInternal(value int) string {
	switch {
	case value%3 == 0 && value%5 == 0:
		return "FizzBuzz"
	case value%3 == 0:
		return "Fizz"
	case value%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(value)
	}
}
