package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"

	//"tc-micro-idp/jwt"
	"tc-micro-idp/models"
	. "tc-micro-idp/utils"
	"time"
)

var (
	dbUrl string
	db    *gorm.DB
	err   error
)

//var db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

func init() {
	switch os.Getenv("ENV") {
	case Test:
		dbUrl = PgUrlTest
	case Dev:
		dbUrl = PgUrlDev
	case Prod:
		dbUrl = PgUrlProd
	default:
		dbUrl = PgUrlLocal
	}
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	err = db.Debug().AutoMigrate(&models.Client{}, &models.User{}, &models.Role{}, &models.UserCity{}, &models.OtpAttempt{}, &models.UserRole{}, &models.BlockedPhone{}, &models.UsedInviteCode{}, &models.Device{}, &models.AccessToken{}, &models.RefreshToken{}, &models.UserProfile{})
	if err != nil {
		log.Fatalln("Auto migrating error : ", err)
	}

}

func Initial() {
	ok := SetClientTable(getClientsTable())
	if !ok {
		log.Println("Cannot set the client table")
	}
}
func getClientsTable() (A []models.Client) {
	err = db.Debug().Table(Clients).Order(`"Id" desc`).Model(A).Find(&A).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("Client Table : \n", A)
	return
}

func FindClient(Issuer string) (ClientTable models.Client) {
	a, _ := GetClientsTable()
	for _, client := range a {
		if Issuer == client.Issuer {
			return client
		}
	}
	return
}

func IsBlocked(Number string) bool {
	Blocked := new([]models.BlockedPhone)
	err = db.Debug().Where(`"Number" = ?`, Number).Find(Blocked).Error
	if err != nil {
		log.Print(err)
		return false
	}
	log.Println(Blocked)
	for _, blocked := range *Blocked {
		if blocked.Number == Number {
			log.Println(true, Number)
			return true
		}
	}
	return false
}

func GetID(Phone string) sql.NullInt64 {
	var otpAttempts models.OtpAttempt
	err = db.Debug().Where(`"Phone" = ?`, Phone).Find(&otpAttempts).Error
	if err != nil {
		log.Println(err)
		return sql.NullInt64{}
	}
	return otpAttempts.UserId

}

func IsOtpAttemptExceededAsync(Model *models.ChallengeInput) bool {
	var limitationTime = time.Unix(time.Now().UTC().Unix()-int64(10*time.Minute.Seconds()), 0)
	var AttemptLimitationCount int64 = 3
	otp := new(models.OtpAttempt)
	var count int64
	err = db.Debug().Model(&models.OtpAttempt{}).Where(`"Phone" = ? AND "IssueTime" >= ?`, Model.Phone, limitationTime).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	err = db.Debug().Where(`"Phone" = ?`, Model.Phone).Order(`"IssueTime"`).Find(&otp).Error
	if err != nil {
		log.Println("IssueTime", err)
	}
	log.Println("132")
	if count > AttemptLimitationCount {
		log.Println(AttemptLimitationCount, err)
		return true
	}
	log.Println("137")
	if (time.Unix(time.Now().UTC().Unix()-int64(3*time.Minute.Seconds()), 0).Unix()) < otp.IssueTime.Unix() {
		log.Println("nkijb")
		return true
	}
	log.Println("142")
	return false
}

func FindUserWithRoles(UserId sql.NullInt64, Phone string) (*models.User, []string) {
	user := new(models.User)
	var role []string
	var roles []models.Role

	err = db.Debug().Preload("UserRoles").Model(user).Where(`"PhoneNumber" = ?`, Phone).Find(user).Error
	log.Println("user entities:", user, "err:", err)
	log.Println("user id : ", user.Id)
	if user.Id != 0 {

		err = db.Debug().Preload(`"Roles"."Name"`).Model(&models.Role{}).Where(`"Id" = ?`, -12).Find(&roles).Error
		if err != nil {
			log.Println("user roles error : ", err)
		}
		fmt.Println("user roles ", roles)
	}
	if UserId.Valid {
		err = db.Debug().Preload("Roles.Name").Model(&models.Role{}).Where(models.Role{}.Id, user.UserRoles).Find(&roles).Error
		if err != nil {
			log.Println("roles name error : ", err)
		}
		fmt.Println("b:", roles[0].Name)
	}
	return user, role
}

func CheckInviteCode(inviteCode string) bool {
	var inviterUserProfile string
	if inviteCode == "" {
		return true
	}
	err = db.Debug().Table(models.UserProfile{}.TableName()).Where(`"InviteCode" = ?`).First(&inviterUserProfile).Error
	if err != nil {
		log.Println("user InviteCode error : ", err)
	}
	return inviterUserProfile != ""
}

func InsertOtpAttempt(Model *models.OtpAttempt) (err error) {
	err = db.Debug().Create(Model).Error
	return
}

func FindOtpAttempt(OtpId int64) (OtpAttempt *models.OtpAttempt, err error) {
	OtpAttempt = new(models.OtpAttempt)
	err = db.Debug().Where(`"Id" = ?`, OtpId).Find(OtpAttempt).Error
	return
}

func UpdateOtpAttemptUserId(Model *models.OtpAttempt) (err error) {
	err = db.Debug().Where(`"Id" = ?`, Model.Id).UpdateColumns(Model).Error
	return err
}

func InsertUser(Model *models.User) (err error) {
	err = db.Debug().Create(Model).Error
	return
}

func InsertUserProfile(Model *models.UserProfile) (err error) {
	err = db.Debug().Create(Model).Error
	return
}
func InsertRefresh(Model *models.RefreshToken) (err error) {
	err = db.Debug().Create(Model).Error
	return
}

func InsertAccess(Model *models.AccessToken) (err error) {
	err = db.Debug().Create(Model).Error
	return
}

func LogOut(Model *models.TokenClaim) (err error) {
	var accessToken models.AccessToken
	accId, _ := strconv.Atoi(Model.AccessVersion)
	err = db.Debug().Where(`"AccessTokens"."Id" = ?`, accId).Find(&accessToken).Error
	log.Println("Loaded accessToken entity : ", accessToken)
	err = db.Debug().Where(`"RefreshTokens"."Id" = ?`, accessToken.RefreshTokenId).Find(&accessToken.RefreshToken).Error
	log.Println("Loaded accessToken entity : ", accessToken)
	accessToken.IsRevoked = true
	accessToken.RevokeTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	accessToken.ModifyTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	accessToken.RefreshToken.IsRevoked = true
	accessToken.RefreshToken.RevokeTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	accessToken.RefreshToken.ModifyTime = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	err = db.Debug().Save(&accessToken).Error
	if err != nil {
		log.Fatalln("Error of Saving accessToken for LogOut : ", err)
	}
	err = db.Debug().Save(&accessToken.RefreshToken).Error
	if err != nil {
		log.Fatalln("Error of Saving accessToken for LogOut : ", err)
	}

	return err
}
func UpdateAvatarFileId(UserId int64, FileId int64) (err error) {
	var user models.User
	err = db.Debug().Where(`"Id" = ?`).Find(&user).Error
	if err != nil {
		log.Println(err)
		return err
	}
	if user.Id != UserId {
		return fmt.Errorf("user does not exist")
	}
	var userProfile = &models.UserProfile{
		ModifyTime: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
		ProfileImageFileId: FileId,
	}
	err = db.Debug().Where(`"UserId" = ?`, UserId).Updates(&userProfile).Error
	if err != nil {
		log.Println(err)
		return err
	}
	return
}
