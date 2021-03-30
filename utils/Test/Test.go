package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"tc-micro-idp/models"
)

func main() {
	//utils.SendOtpCode("9358378702","6541","kdfjgvbkjdvn")

	db, err := gorm.Open(postgres.Open("postgresql://Misagh:13750620@127.0.0.1:5432/tazminchi"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	//err = db.Debug().AutoMigrate(&models.Client{},&models.User{},&models.Role{},&models.OtpAttempt{},models.BlockedPhone{},models.Client{}, models.OtpAttempt{}, models.Role{}, models.UserRole{}, models.User{}, models.UserProfile{}, models.UserCity{}, models.UsedInviteCode{}, models.Device{},models.RefreshToken{},models.AccessToken{})
	err = db.Debug().AutoMigrate(&models.Client{}, &models.User{}, &models.Role{}, models.UserCity{}, &models.OtpAttempt{}, models.UserRole{}, models.BlockedPhone{}, models.UsedInviteCode{}, &models.Device{}, models.AccessToken{}, models.RefreshToken{})
	if err != nil {
		log.Fatalln("Auto migrating error : ", err)
	}
}
