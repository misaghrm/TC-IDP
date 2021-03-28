package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tc-micro-idp/models"
	. "tc-micro-idp/utils"
	"time"
)

var (
	dbUrl string
	db *gorm.DB
	err error
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
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	//err = db.Debug().AutoMigrate(&models.Client{},&models.User{},&models.Role{},&models.OtpAttempt{},models.BlockedPhone{},models.Client{}, models.OtpAttempt{}, models.Role{}, models.UserRole{}, models.User{}, models.UserProfile{}, models.UserCity{}, models.UsedInviteCode{}, models.Device{},models.RefreshToken{},models.AccessToken{})
	err = db.Debug().AutoMigrate(&models.Client{},&models.User{},&models.Role{}, models.UserCity{},&models.OtpAttempt{}, models.UserRole{},models.BlockedPhone{}, models.UsedInviteCode{}, &models.Device{},models.AccessToken{},models.RefreshToken{})
	if err != nil {
		log.Println("Auto migrating error : " , err)
	}
	ok := SetClientTable(getClientsTable())
	if !ok {
		log.Println("Cannot set the client table")
	}
}

func getClientsTable() (A []*models.Client) {
	err = db.Table(Clients).Order(`"Id" desc`).Model(A).Find(&A).Error
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("Client Table : \n", A)
	return
}

func FindClient(Issuer string) (ClientTable *models.Client) {
	a := getClientsTable()
	for _, client := range a {
		if Issuer == client.Issuer {
			return client
		}
	}
	return nil
}

func IsBlocked(Number string) bool {
	Blocked := new([]models.BlockedPhone)
	err = db.Where(`"Number" = ?`, Number).Find(Blocked).Error
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

func GetID(Phone string) int64 {
	var otpAttempts models.OtpAttempt
	err = db.Debug().Where(`"Phone" = ?`, Phone).Last(&otpAttempts).Error
	if err != nil {
		log.Println(err)
		return 0
	}
	return otpAttempts.UserId
	//log.Println(otpAttempts)
	//if otpAttempts.UserId == 0 {
	//	otpAttempts = models.OtpAttempt{
	//		Base:       models.Base{
	//			Id:           IdGenerator.Generate().Int64(),
	//			CreationTime: time.Now().UTC(),
	//			ModifyTime:   nil,
	//		},
	//		Phone:      Phone,
	//		ClientId:   0,
	//		Salt:       "",
	//		IssueTime:  time.Now().UTC(),
	//		ExpireTime: time.Now().UTC().Add(120*time.Second),
	//		UserIp:     Ctx.IP(),
	//		UserAgent:  Ctx.Get(UserAgent),
	//		OtpKind:    0,
	//	}
	//	db.Debug().Create(otpAttempts)
	//}
}

func IsOtpAttemptExceededAsync(Model *models.ChallengeInput) bool {
	var limitationTime = time.Unix(time.Now().UTC().Unix()-int64(10*time.Minute.Seconds()),0)
	var AttemptLimitationCount int64 = 3
	otp := new(models.OtpAttempt)
	var count int64
	err = db.Debug().Where(`"PhoneNumber" = ? AND "IssueTime" >= ?`,Model.Phone,limitationTime).Count(&count).Error
	if err != nil {
		log.Println(err)
	}
	err = db.Debug().Where(`"PhoneNumber" = ?`,Model.Phone).Order(otp.IssueTime).Last(&otp).Error
	if err != nil {
		log.Println(err)
	}
	if count >= AttemptLimitationCount {
		return true
	}

	if (time.Unix(time.Now().UTC().Unix()-int64(3*time.Minute.Seconds()),0).Unix()) < otp.IssueTime.Unix() {
		return true
	}
	return false
}

func FindUserWithRoles(UserId int64, Phone string) (*models.User, []string) {
	user := new(models.User)
	var role []string
	var roles []models.Role

	err = db.Debug().Preload("UserRoles").Model(user).Where(`"PhoneNumber" = ?`, Phone).Find(user).Error
	log.Println("user entities:", user, "err:", err)
	log.Println("user id : ",user.Id)
	if user.Id != 0 {

		err = db.Debug().Preload("Roles.Name").Model(&models.Role{}).Where(`"id" = ?`, user.UserRoles).Find(&roles).Error
		if err != nil {
			log.Println("user roles error : ", err)
		}
		fmt.Println("user roles ", roles)
	}
	if UserId > 0 {
		err = db.Debug().Preload("Roles.Name").Model(&models.Role{}).Where(models.Role{}.Id, user.UserRoles).Find(&roles).Error
		if err != nil {
			log.Println("roles name error : ", err)
		}
		fmt.Println("b:",roles[0].Name)
	}
	return user, role
}


func CheckInviteCode(inviteCode string) bool {
	var inviterUserProfile string
	if inviteCode == "" {
		return false
	}
	err = db.Debug().Table(models.UserProfile{}.TableName()).Where(`"InviteCode" = ?`).First(&inviterUserProfile).Error
	if err != nil {
		log.Println("user InviteCode error : ", err)
	}
	return inviterUserProfile != ""
}