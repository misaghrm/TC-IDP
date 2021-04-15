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
		return c.SendStatus(http.StatusBadRequest)
	}
	salt, otp := publicFunctions.GenerateOtp(publicFunctions.StringGenerator(Input.Phone))
	utils.SendOtpCode(Input.Phone, otp, Input.AppSignatureHash)

	fmt.Println("ClientId : ", Client.Id)
	var otpAttempt = models.OtpAttempt{
		Id:           newId().Int64(),
		CreationTime: time.Now().UTC(),
		//UserId:       sql.NullInt64{},
		Phone:      Input.Phone,
		ClientId:   Client.Id,
		Client:     models.Client{},
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
		Code: http.StatusOK,
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
		Code: 200,
	})
}

func Verify(c *fiber.Ctx) error {
	UserRole := new([]models.UserRole)
	Input := new(models.VerifyInput)
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

	if strings.Contains(Client.AccessTokenLifeTime, " day") {
		Client.AccessTokenLifeTime = strings.ReplaceAll(Client.AccessTokenLifeTime, " days", "")
		Client.AccessTokenLifeTime = strings.ReplaceAll(Client.AccessTokenLifeTime, " day", "")
		a, _ := strconv.ParseInt(Client.AccessTokenLifeTime, 10, 64)
		Client.AccessTokenLifeTime = fmt.Sprintf("%vh", 24*a)
		log.Println("AccessTokenLifeTime : ", Client.AccessTokenLifeTime)
	}
	if strings.Contains(Client.RefreshTokenLifeTime, " day") {
		Client.RefreshTokenLifeTime = strings.ReplaceAll(Client.RefreshTokenLifeTime, " days", "")
		Client.RefreshTokenLifeTime = strings.ReplaceAll(Client.RefreshTokenLifeTime, " day", "")
		a, _ := strconv.ParseInt(Client.RefreshTokenLifeTime, 10, 64)
		Client.RefreshTokenLifeTime = fmt.Sprintf("%vh", 24*a)
		log.Println("RefreshTokenLifeTime : ", Client.RefreshTokenLifeTime)
	}

	log.Println(Input.OtpId)
	otpAttempt, err := db.FindOtpAttempt(Input.OtpId)
	if err != nil {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16003OtpIdNotFound,
			Message: "درخواستی برای این کاربر پیدا نشد.",
		})
	}
	if otpAttempt == nil || otpAttempt.Phone != Input.Phone {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16003OtpIdNotFound,
			Message: "درخواستی برای این کاربر پیدا نشد.",
		})
	}
	if otpAttempt.ExpireTime.Unix() < time.Now().UTC().Unix() {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16005OtpCodeExpired,
			Message: "کد شما منقضی شده است.",
		})
	}
	if ok = publicFunctions.IsOtpValid(Input.Code, otpAttempt.Salt); !ok {
		return c.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16006WrongOtpCode,
			Message: "کد اشتباه است.",
		})
	}
	user, roles := db.FindUserWithRoles(otpAttempt.UserId, Input.Phone)
	if !otpAttempt.UserId.Valid && user.Id == 0 {
		user = &models.User{
			Id:            newId().Int64(),
			CreationTime:  time.Now().UTC(),
			PhoneNumber:   Input.Phone,
			LastLoginTime: time.Now().UTC(),
			UserRoles:     *UserRole,
			RefreshTokens: nil,
		}
		err = db.InsertUser(user)
		if err != nil {
			log.Fatalln("InsertUser error : ", err)
			return err
		}
		var userProfile = &models.UserProfile{
			Id:           newId().Int64(),
			CreationTime: time.Now().UTC(),
			UserId:       user.Id,
			InviteCode:   publicFunctions.InviteCodeGenerator(user.PhoneNumber),
		}
		err = db.InsertUserProfile(userProfile)
		if err != nil {
			log.Fatalln("InsertUserProfile error : ", err)
			return err
		}
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
	user.LastLoginTime = time.Now().UTC()
	if len(Client.DefaultRoles) > 0 {
		a := strings.Split(Client.DefaultRoles, ",")
		for _, defaultRole := range a {
			for _, m := range RoleList {
				if defaultRole == m.Name {
					*UserRole = append(*UserRole, models.UserRole{
						UserId: user.Id,
						RoleId: m.Id.Int64,
					})
				}
			}
		}
	}
	var refresh = jwt.GenerateRefreshToken()
	d, err := time.ParseDuration(Client.RefreshTokenLifeTime)
	if err != nil {
		log.Fatalln("ParseDuration error : ", err)
	}
	log.Println("Duration of refresh token life time : ", d)
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

	d, err = time.ParseDuration(Client.AccessTokenLifeTime)
	if err != nil {
		log.Fatalln("ParseDuration error : ", err)
	}

	var accessToken = &models.AccessToken{
		Id:             newId().Int64(),
		CreationTime:   time.Now().UTC(),
		RefreshTokenId: refreshToken.Id,
		IssueTime:      time.Now().UTC(),
		ExpireTime:     time.Now().UTC().Add(d),
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
		Expires:        strconv.FormatInt(time.Now().UTC().Add(d).Unix(), 10),
		NotBefore:      strconv.FormatInt(time.Now().UTC().Unix(), 10),
	}

	var access = jwt.GenerateToken(Tokenclaim, &Client)
	accessToken.Token = access
	err = db.InsertRefresh(refreshToken)
	if err != nil {
		log.Fatalln("InsertRefresh error : ", err)
		return err
	}
	err = db.InsertAccess(accessToken)
	if err != nil {
		log.Fatalln("InsertAccess error : ", err)
		return err
	}
	return c.JSON(&models.ResponseModel{
		Data: map[string]interface{}{
			"accessToken":  access,
			"refreshToken": refresh,
		},
		Code: 200,
	})

}

func LogOut(ctx *fiber.Ctx) error {
	tokenClaims, err := jwt.Decrypt(ExtractToken(ctx), ctx.Get(utils.ClientKey))
	if err != nil {
		return ctx.SendStatus(http.StatusUnauthorized)
	}
	if !tokenClaims.IsLifeTimeValid() {
		if err != nil {
			return ctx.SendStatus(http.StatusUnauthorized)
		}
	}
	err = db.LogOut(tokenClaims)

	if err != nil {
		log.Println(err)
		return ctx.JSON(&models.ResponseModel{
			Data:    map[string]interface{}{"value": false},
			Code:    500,
			Message: "failed",
		})
	}
	return ctx.JSON(&models.ResponseModel{
		Data: map[string]interface{}{"value": true},
		Code: 200,
	})
}

func Refresh(ctx *fiber.Ctx) error {
	var ok bool
	Client, ok = db.CanLogin(ctx.Get(utils.ClientKey))
	if !ok {
		return ctx.JSON(models.ResponseModel{
			Data:    nil,
			Code:    utils.Error16001InvalidClientKey,
			Message: "اپلیکیشن شما ناشناخته است.",
		})
	}

	Input := new(models.RefreshInput)
	if err := ctx.BodyParser(Input); err != nil {
		return err
	}

	tokenClaims, err := jwt.Decrypt(Input.AccessToken, ctx.Get(utils.ClientKey))
	if err != nil {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	if db.IsBlocked(tokenClaims.Phone) {
		return ctx.JSON(models.ResponseModel{
			Code:    utils.Error16014NumberIsBlocked,
			Message: "شماره همراه اشتباه است.",
		})
	}

	userid, err := strconv.ParseInt(tokenClaims.UserId, 10, 64)
	user, roles := db.FindUserWithRoles(sql.NullInt64{Int64: userid, Valid: true}, tokenClaims.Phone)
	if user == nil {
		return ctx.JSON(models.ResponseModel{
			Code:    utils.Error16002UserNotFound,
			Message: "کاربر مورد نظر پیدا نشد.",
		})
	}

	DeviceId, err := strconv.ParseInt(tokenClaims.DeviceId, 10, 64)
	if err != nil {
		log.Fatalln("ParseInt error of tokenClaims.DeviceId : ", err)
	}
	RefreshTokenModel, err := db.FindRefreshToken(Client.Id, user.Id, DeviceId, Input.RefreshToken)
	if err != nil {
		log.Fatalln("FindRefreshToken error : ", err)
	}
	if RefreshTokenModel.Token == "" {
		return ctx.JSON(models.ResponseModel{
			Code:    utils.Error16004RefreshTokenNotFound,
			Message: "درخواستی برای این کاربر پیدا نشد.",
		})
	}

	if RefreshTokenModel.ExpireTime.Unix() < time.Now().UTC().Unix() || RefreshTokenModel.IsRevoked {
		return ctx.JSON(models.ResponseModel{
			Code:    utils.Error16009RefreshTokenExpired,
			Message: "لطفا مجددا وارد شوید.",
		})
	}

	if strings.Contains(Client.AccessTokenLifeTime, " day") {
		Client.AccessTokenLifeTime = strings.ReplaceAll(Client.AccessTokenLifeTime, " days", "")
		Client.AccessTokenLifeTime = strings.ReplaceAll(Client.AccessTokenLifeTime, " day", "")
		a, _ := strconv.ParseInt(Client.AccessTokenLifeTime, 10, 64)
		Client.AccessTokenLifeTime = fmt.Sprintf("%vh", 24*a)
		log.Println("AccessTokenLifeTime : ", Client.AccessTokenLifeTime)
	}
	if strings.Contains(Client.RefreshTokenLifeTime, " day") {
		Client.RefreshTokenLifeTime = strings.ReplaceAll(Client.RefreshTokenLifeTime, " days", "")
		Client.RefreshTokenLifeTime = strings.ReplaceAll(Client.RefreshTokenLifeTime, " day", "")
		a, _ := strconv.ParseInt(Client.RefreshTokenLifeTime, 10, 64)
		Client.RefreshTokenLifeTime = fmt.Sprintf("%vh", 24*a)
		log.Println("RefreshTokenLifeTime : ", Client.RefreshTokenLifeTime)
	}

	var refresh = jwt.GenerateRefreshToken()
	d, err := time.ParseDuration(Client.RefreshTokenLifeTime)
	if err != nil {
		log.Fatalln("ParseDuration error : ", err)
	}
	log.Println("Duration of refresh token life time : ", d)
	var refreshToken = &models.RefreshToken{}
	//refreshtokenid := sql.NullInt64{
	//	Int64: newId().Int64(),
	//	Valid: true,
	//}

	refreshToken = &models.RefreshToken{
		Id:           sql.NullInt64{Int64: newId().Int64()}.Int64,
		CreationTime: time.Now().UTC(),
		UserId:       user.Id,
		ClientId:     Client.Id,
		Token:        refresh,
		IssueTime:    time.Now().UTC(),
		ExpireTime:   time.Now().UTC().Add(d),
		DeviceId:     RefreshTokenModel.DeviceId,
		Device:       RefreshTokenModel.Device,
	}

	d, err = time.ParseDuration(Client.AccessTokenLifeTime)
	if err != nil {
		log.Fatalln("ParseDuration error : ", err)
	}

	var accessToken = &models.AccessToken{
		Id:             newId().Int64(),
		CreationTime:   time.Now().UTC(),
		RefreshTokenId: refreshToken.Id,
		IssueTime:      time.Now().UTC(),
		ExpireTime:     time.Now().UTC().Add(d),
	}

	var Tokenclaim = &models.TokenClaim{
		TokenId:        strconv.FormatInt(accessToken.Id, 10),
		IssuedAt:       time.Now().UTC().String(),
		UserId:         strconv.FormatInt(user.Id, 10),
		Phone:          user.PhoneNumber,
		RefreshVersion: strconv.FormatInt(refreshToken.Id, 10),
		EulaVersion:    utils.EulaVersion,
		Issuer:         Client.Issuer,
		LifeTime:       Client.AccessTokenLifeTime,
		AccessVersion:  strconv.FormatInt(accessToken.Id, 10),
		DeviceId:       strconv.FormatInt(refreshToken.Device.Id.Int64, 10),
		AppSource:      Client.Issuer,
		Roles:          roles,
		Audience:       Client.Audience,
		Expires:        strconv.FormatInt(time.Now().UTC().Add(d).Unix(), 10),
		NotBefore:      strconv.FormatInt(time.Now().UTC().Unix(), 10),
	}

	var access = jwt.GenerateToken(Tokenclaim, &Client)
	accessToken.Token = access
	err = db.InsertRefresh(refreshToken)
	if err != nil {
		log.Fatalln("InsertRefresh error : ", err)
		return err
	}
	err = db.InsertAccess(accessToken)
	if err != nil {
		log.Fatalln("InsertAccess error : ", err)
		return err
	}
	return ctx.JSON(&models.ResponseModel{
		Data: map[string]interface{}{
			"accessToken":  access,
			"refreshToken": refresh,
		},
		Code: 200,
	})

}
