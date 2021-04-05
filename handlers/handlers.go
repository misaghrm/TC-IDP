package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
	"strings"
	"tc-micro-idp/jwt"
	"tc-micro-idp/models"
	. "tc-micro-idp/utils"
)

func init() {
	switch os.Getenv("ENV") {

	}
}

func TestToken(c *fiber.Ctx) error {
	token, err := jwt.Decrypt(ExtractToken(c), c.Get(ClientKey))
	if err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}
	fmt.Println(token)
	return c.JSON(&models.ResponseModel{
		Data: map[string]interface{}{"tokenClaims": token},
		Code: 200,
	})

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
