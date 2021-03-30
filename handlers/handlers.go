package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strings"
	"tc-micro-idp/jwt"
	. "tc-micro-idp/utils"
)

func init() {
	switch os.Getenv("ENV") {

	}
}

func TestToken(c *fiber.Ctx) error {
	_, err := jwt.Decrypt(ExtractToken(c), c.Get(ClientKey))
	if err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	return c.SendStatus(http.StatusOK)

}

// ExtractToken read the token from the request header
func ExtractToken(c *fiber.Ctx) string {
	bearToken := c.Get(Authorization)
	if !strings.Contains(bearToken, "Bearer ") {
		return bearToken
	}
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
