package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
	. "tc-micro-idp/dataManager"
	"tc-micro-idp/models"
	"tc-micro-idp/utils"
	"tc-micro-idp/utils/publicFunctions"
)

var (
	id int64
)

func ChallengeToken(c *fiber.Ctx) error {
	Input := new(models.ChallengeInput)

	if err := c.BodyParser(Input); err != nil {
		return err
	}
	if Input.Phone = utils.CleanUpPhone(Input.Phone); Input.Phone == "" {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16013InvalidPhoneNumber,
			Message: "شماره همراه اشتباه است.",
		})
	}

	if IsBlocked(Input.Phone) {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16014NumberIsBlocked,
			Message: "شماره همراه اشتباه است.",
		})
	}

	Client, ok := CanLogin(c.Get(utils.MNXClient))
	if !ok {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16001InvalidClientKey,
			Message: "اپلیکیشن شما ناشناخته است.",
		})
	}

	ClientId := GetID(Input.Phone)
	if ClientId == 0 {
		if Client.CanRegister {
			return RegisterChallenge(c,Input)
		}
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16012ClientRegistrationNotAllowed,
			Message: "It is not possible to register user in this client",
		})
	}

	return c.SendStatus(500)
}

func RegisterChallenge(c *fiber.Ctx,Input *models.ChallengeInput) error {
	if IsOtpAttemptExceededAsync(Input) {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16008AttemptLimitationExceeded,
			Message: "تعداد درخواست بیش از حد مجاز است.",
		})
	}
	isCodeValid := CheckInviteCode(Input.InviteCode)
	if !isCodeValid {
		return c.JSON("")
	}

	nt, _ := strconv.Atoi(Input.Phone)
	salt, otp := publicFunctions.GenerateOtp(publicFunctions.StringGenerator(nt))
	utils.SendOtpCode(Input.Phone,otp,Input.AppSignatureHash)
	fmt.Println("salt : ",salt)

	return nil
}

func LoginChallenge(c *fiber.Ctx, Phone string) error {

	FindUserWithRoles(0, Phone)
	return nil
}
