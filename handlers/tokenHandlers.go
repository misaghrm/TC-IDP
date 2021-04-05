package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
	"strings"
	db "tc-micro-idp/dataManager"
	"tc-micro-idp/jwt"
	"tc-micro-idp/models"
	"tc-micro-idp/utils"
	"tc-micro-idp/utils/publicFunctions"
	"time"
)

var RoleList []models.Role

func init() {
	RoleList = utils.RolesConst
}

var Client models.Client
var newId = publicFunctions.IdGenerator.Generate

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

	if db.IsBlocked(Input.Phone) {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16014NumberIsBlocked,
			Message: "شماره همراه اشتباه است.",
		})
	}
	var ok bool
	Client, ok = db.CanLogin(c.Get(utils.ClientKey))
	if !ok {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16001InvalidClientKey,
			Message: "اپلیکیشن شما ناشناخته است.",
		})
	}

	UserId := db.GetID(Input.Phone)

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
	if db.IsOtpAttemptExceededAsync(Input) {
		return c.JSON(&models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16008AttemptLimitationExceeded,
			Message: "تعداد درخواست بیش از حد مجاز است.",
		})
	}

	isCodeValid := db.CheckInviteCode(Input.InviteCode)
	if !isCodeValid {
		return c.JSON("")
	}

	salt, otp := publicFunctions.GenerateOtp(publicFunctions.StringGenerator(Input.Phone))
	utils.SendOtpCode(Input.Phone, otp, Input.AppSignatureHash)

	fmt.Println("ClientId : ", Client.Id)
	var otpAttempt = models.OtpAttempt{
		Id:           newId().Int64(),
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

	if err = db.InsertOtpAttempt(&otpAttempt); err != nil {
		log.Println("InsertOtpAttempt error : ", err)
		return err
	}

	return c.JSON(models.ResponseModel{
		Data: map[string]interface{}{
			utils.OtpId: otpAttempt.Id,
			"otp":       otp,
		},
		Code:    200,
		Message: "",
	})
}

func LoginChallenge(c *fiber.Ctx, Input *models.ChallengeInput) (err error) {
	if db.IsOtpAttemptExceededAsync(Input) {
		return c.JSON(&models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16008AttemptLimitationExceeded,
			Message: "تعداد درخواست بیش از حد مجاز است.",
		})
	}
	user, _ := db.FindUserWithRoles(sql.NullInt64{}, Input.Phone)

	salt, otp := publicFunctions.GenerateOtp(publicFunctions.StringGenerator(Input.Phone))
	utils.SendOtpCode(Input.Phone, otp, Input.AppSignatureHash)

	var otpAttempt = models.OtpAttempt{
		Id:           newId().Int64(),
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

	if err = db.InsertOtpAttempt(&otpAttempt); err != nil {
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

func Verify(c *fiber.Ctx) error {
	UserRole := new([]models.UserRole)
	Input := new(models.VerifyInput)
	log.Println("verify 1")
	if err := c.BodyParser(Input); err != nil {
		return err
	}
	log.Println("verify 2")
	if Input.Phone = utils.CleanUpPhone(Input.Phone); Input.Phone == "" {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16013InvalidPhoneNumber,
			Message: "شماره همراه اشتباه است.",
		})
	}
	log.Println("verify 3")
	if db.IsBlocked(Input.Phone) {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16014NumberIsBlocked,
			Message: "شماره همراه اشتباه است.",
		})
	}
	log.Println("verify 4")
	var ok bool
	Client, ok = db.CanLogin(c.Get(utils.ClientKey))
	log.Println("verify 5")
	if !ok {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16001InvalidClientKey,
			Message: "اپلیکیشن شما ناشناخته است.",
		})
	}
	log.Println("verify 6")
	log.Println(Input.OtpId)
	otpAttempt, err := db.FindOtpAttempt(Input.OtpId)
	if err != nil {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16003OtpIdNotFound,
			Message: "درخواستی برای این کاربر پیدا نشد.",
		})
	}
	log.Println("verify 7")
	if otpAttempt == nil || otpAttempt.Phone != Input.Phone {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16003OtpIdNotFound,
			Message: "درخواستی برای این کاربر پیدا نشد.",
		})
	}
	log.Println("verify 8")
	if otpAttempt.ExpireTime.Unix() < time.Now().UTC().Unix() {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16005OtpCodeExpired,
			Message: "کد شما منقضی شده است.",
		})
	}
	log.Println("verify 9")
	if ok := publicFunctions.IsOtpValid(Input.Code, otpAttempt.Salt); !ok {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16006WrongOtpCode,
			Message: "کد اشتباه است.",
		})
	}
	log.Println("verify 10")
	user, roles := db.FindUserWithRoles(otpAttempt.UserId, Input.Phone)
	log.Println("verify 11")
	if !otpAttempt.UserId.Valid && user.Id == 0 {
		log.Println("verify 12")
		user = &models.User{
			Id:            newId().Int64(),
			CreationTime:  time.Now().UTC(),
			PhoneNumber:   Input.Phone,
			LastLoginTime: time.Now().UTC(),
			UserRoles:     *UserRole,
			RefreshTokens: nil,
		}
		log.Println("verify 13")
		err = db.InsertUser(user)
		if err != nil {
			log.Fatalln("InsertUser error : ", err)
			return err
		}
		log.Println("verify 14")
		var userProfile = &models.UserProfile{
			Id:           newId().Int64(),
			CreationTime: time.Now().UTC(),
			UserId:       user.Id,
			InviteCode:   publicFunctions.InviteCodeGenerator(user.PhoneNumber),
		}
		log.Println("verify 15")
		err = db.InsertUserProfile(userProfile)
		if err != nil {
			log.Fatalln("InsertUserProfile error : ", err)
			return err
		}
		log.Println("verify 16")
		otpAttempt.UserId = sql.NullInt64{
			Int64: user.Id,
			Valid: true,
		}
		otpAttempt.ModifyTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}
		err = db.UpdateOtpAttemptUserId(otpAttempt)
		if err != nil {
			log.Fatalln("UpdateOtpAttemptUserId : ", err)
		}
	}

	log.Println("verify 17")
	user.LastLoginTime = time.Now().UTC()

	if len(Client.DefaultRoles) > 0 {
		log.Println("verify 18")
		a := strings.Split(Client.DefaultRoles, ",")
		for _, defaultRole := range a {
			for _, m := range RoleList {

				if defaultRole == m.Name {
					*UserRole = append(*UserRole, models.UserRole{
						UserId: user.Id,
						RoleId: m.Id,
					})
				}
			}
		}
		log.Println("verify 19")
	}
	log.Println("verify 20")
	var refresh = jwt.GenerateRefreshToken()
	d, err := time.ParseDuration("720h")
	if err != nil {
		log.Fatalln("ParseDuration error : ", err)
	}
	log.Println("Duration of refresh token life time : ", d)
	log.Println("verify 21")
	var refreshToken = &models.RefreshToken{}
	refreshtokenid := sql.NullInt64{
		Int64: newId().Int64(),
		Valid: true,
	}
	deviceid := sql.NullInt64{
		Int64: newId().Int64(),
		Valid: true,
	}
	refreshToken = &models.RefreshToken{
		Id:           refreshtokenid.Int64,
		CreationTime: time.Now().UTC(),
		ModifyTime: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		UserId:     user.Id,
		ClientId:   Client.Id,
		Token:      refresh,
		IssueTime:  time.Now().UTC(),
		ExpireTime: time.Now().UTC().Add(d),
		//DeviceId: deviceid,
		Device: models.Device{
			Id:             deviceid,
			CreationTime:   time.Now().UTC(),
			UserIp:         c.IP(),
			UserAgent:      c.Get(utils.UserAgent),
			FireBaseId:     Input.FirebaseId,
			YandexId:       Input.YandexId,
			Imei:           Input.Imei,
			PhoneModel:     Input.PhoneModel,
			AndroidVersion: Input.AndroidVersion,
			ScreenSize:     Input.ScreenSize,
			AppVersion:     Input.AppVersion,
			SimOperator:    Input.SimOperator,
			//RefreshTokenId: refreshtokenid.Int64,
			//AppSource:      models.AppSource(Input.AppSource),
		},
	}
	//device :=
	var accessToken = &models.AccessToken{
		Id:             newId().Int64(),
		CreationTime:   time.Now().UTC(),
		RefreshTokenId: refreshToken.Id,
		IssueTime:      time.Now().UTC(),
		ExpireTime:     time.Now().UTC(),
	}
	d, err = time.ParseDuration("24h")
	if err != nil {
		log.Fatalln("ParseDuration error : ", err)
	}
	var Tokenclaim = &models.TokenClaim{
		TokenId:        strconv.FormatInt(accessToken.Id, 10),
		IssuedAt:       time.Now().UTC().String(),
		UserId:         strconv.FormatInt(user.Id, 10),
		Phone:          user.PhoneNumber,
		RefreshVersion: strconv.FormatInt(refreshToken.Id, 10),
		EulaVersion:    utils.EulaVersion,
		Issuer:         "",
		LifeTime:       Client.AccessTokenLifeTime,
		AccessVersion:  strconv.FormatInt(accessToken.Id, 10),
		DeviceId:       strconv.FormatInt(refreshToken.Device.Id.Int64, 10),
		AppSource:      Input.AppSource,
		Roles:          roles,
		Audience:       Client.Audience,
		Expires:        time.Now().UTC().Add(d).String(),
		NotBefore:      time.Now().UTC().String(),
	}

	var access = jwt.GenerateToken(Tokenclaim, &Client)
	accessToken.Token = access
	log.Println("verify 22")

	err = db.InsertRefresh(refreshToken)
	if err != nil {
		log.Fatalln("InsertRefresh error : ", err)
		return err
	}
	log.Println("verify 23")

	err = db.InsertAccess(accessToken)
	if err != nil {
		log.Fatalln("InsertAccess error : ", err)
		return err
	}
	log.Println("verify 24")

	return c.JSON(&models.ResponseModel{
		Data: map[string]interface{}{
			"accessToken":  access,
			"refreshToken": refresh,
		},
		Code: 200,
	})

}

func LogOut(c *fiber.Ctx) error {
	tokenClaims, err := jwt.Decrypt(ExtractToken(c), c.Get(utils.ClientKey))
	if err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}
	err = db.LogOut(tokenClaims)
	if err != nil {
		log.Println()
		return c.JSON(&models.ResponseModel{
			Data:    map[string]interface{}{"value": false},
			Code:    500,
			Message: "failed",
		})
	}
	return c.JSON(&models.ResponseModel{
		Data: map[string]interface{}{"value": true},
		Code: 200,
	})
}
