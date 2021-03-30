package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	. "tc-micro-idp/dataManager"
	"tc-micro-idp/models"
	"tc-micro-idp/utils"
	"tc-micro-idp/utils/publicFunctions"
	"time"
)

var Client models.Client

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
	var ok bool
	Client, ok = CanLogin(c.Get(utils.ClientKey))
	if !ok {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16001InvalidClientKey,
			Message: "اپلیکیشن شما ناشناخته است.",
		})
	}

	UserId := GetID(Input.Phone)

	if !UserId.Valid {
		if Client.CanRegister {
			return RegisterChallenge(c, Input)

		}
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16012ClientRegistrationNotAllowed,
			Message: "It is not possible to register user in this client",
		})
	}

	return LoginChallenge(c, Input)

}

func RegisterChallenge(c *fiber.Ctx, Input *models.ChallengeInput) (err error) {
	if IsOtpAttemptExceededAsync(Input) {
		return c.JSON(&models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16008AttemptLimitationExceeded,
			Message: "تعداد درخواست بیش از حد مجاز است.",
		})
	}

	isCodeValid := CheckInviteCode(Input.InviteCode)
	if !isCodeValid {
		return c.JSON("")
	}

	salt, otp := publicFunctions.GenerateOtp(publicFunctions.StringGenerator(Input.Phone))
	utils.SendOtpCode(Input.Phone, otp, Input.AppSignatureHash)

	fmt.Println("ClientId : ", Client.Id)
	var otpAttempt = models.OtpAttempt{
		Id:           publicFunctions.IdGenerator.Generate().Int64(),
		CreationTime: time.Now().UTC(),
		UserId:       sql.NullInt64{},
		Phone:        Input.Phone,
		ClientId:     Client.Id,
		Client:       models.Client{},
		Salt:         salt,
		IssueTime:    time.Now().UTC(),
		ExpireTime:   time.Now().UTC().Add(120 * time.Second),
		UserIp:       c.IP(),
		UserAgent:    c.Get(utils.UserAgent),
		OtpKind:      models.Kind(2),
	}

	if err = InsertOtpAttempt(&otpAttempt); err != nil {
		log.Println("InsertOtpAttempt error : ", err)
		return err
	}

	return c.JSON(models.ResponseModel{
		Data: map[string]interface{}{
			utils.OtpId: otpAttempt.Id,
		},
		Code:    200,
		Message: "",
	})
}

func LoginChallenge(c *fiber.Ctx, Input *models.ChallengeInput) (err error) {
	if IsOtpAttemptExceededAsync(Input) {
		return c.JSON(&models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16008AttemptLimitationExceeded,
			Message: "تعداد درخواست بیش از حد مجاز است.",
		})
	}
	user, _ := FindUserWithRoles(0, Input.Phone)

	salt, otp := publicFunctions.GenerateOtp(publicFunctions.StringGenerator(Input.Phone))
	utils.SendOtpCode(Input.Phone, otp, Input.AppSignatureHash)

	var otpAttempt = models.OtpAttempt{
		Id:           publicFunctions.IdGenerator.Generate().Int64(),
		CreationTime: time.Now().UTC(),
		UserId: sql.NullInt64{
			Int64: user.Id,
		},
		Phone:      user.PhoneNumber,
		ClientId:   Client.Id,
		Salt:       salt,
		IssueTime:  time.Now().UTC(),
		ExpireTime: time.Now().UTC().Add(120 * time.Second),
		UserIp:     c.IP(),
		UserAgent:  c.Get(utils.UserAgent),
		OtpKind:    models.Kind(2),
	}

	if err = InsertOtpAttempt(&otpAttempt); err != nil {
		log.Println("InsertOtpAttempt error : ", err)
		return err
	}

	return c.JSON(models.ResponseModel{
		Data: map[string]interface{}{
			utils.OtpId: otpAttempt.Id,
		},
		Code:    200,
		Message: "",
	})
}
