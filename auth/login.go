package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-sql-api/env"
	"github.com/gofiber/fiber/v2"
)

type UserForm struct {
	Email    string `json:"email" xml:"name" form:"name"`
	Password string `json:"password" xml:"password" form:"password"`
}

func Login(c *fiber.Ctx) error {

	body := new(UserForm)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse Json",
		})
	}

	if body.Email != "test@gmail.com" || body.Password != "password" {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})

	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "1"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["email"] = body.Email
	claims["password"] = body.Password

	s, err := token.SignedString([]byte(env.ENV("SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user": struct {
			Id    int    `json:"id"`
			Email string `json:"email"`
		}{
			Id:    1,
			Email: "test@gmail.com",
		},
	})
}
