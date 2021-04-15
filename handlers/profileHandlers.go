package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	db "tc-micro-idp/dataManager"
	"tc-micro-idp/jwt"
	"tc-micro-idp/models"
	"tc-micro-idp/utils"
)

func UpdateAvatar(ctx *fiber.Ctx) error {
	token, err := jwt.Decrypt(ExtractToken(ctx), ctx.Get(utils.ClientKey))
	if !token.IsLifeTimeValid() {
		if err != nil {
			return ctx.SendStatus(http.StatusUnauthorized)
		}
	}
	if err != nil {
		return ctx.SendStatus(http.StatusUnauthorized)
	}
	a := ctx.Query("fileid", "")
	fileId, _ := strconv.ParseInt(a, 10, 64)
	userId, _ := strconv.ParseInt(token.UserId, 10, 64)
	err = db.UpdateAvatarFileId(userId, fileId)
	if err != nil {
		return ctx.JSON(&models.ResponseModel{
			Code:    utils.Error16000AnErrorHasOccurred,
			Message: "خطایی رخ داده است.\n" + err.Error(),
		})
	}
	return ctx.JSON(&models.ResponseModel{
		Data: map[string]interface{}{"value": true},
		Code: 200,
	})

}
