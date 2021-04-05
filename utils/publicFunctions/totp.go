package publicFunctions

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"log"
	"time"
)

func GenerateOtp(secret string) (salt, Code string) {
	log.Println("GenerateOtp 12")
	salt = base32.StdEncoding.EncodeToString([]byte(secret))
	log.Println("GenerateOtp 14")

	Code, err := totp.GenerateCodeCustom(salt, time.Now().UTC(), totp.ValidateOpts{
		Period:    120,
		Skew:      20,
		Digits:    4,
		Algorithm: otp.AlgorithmSHA1,
	})
	log.Println("GenerateOtp 22")

	if err != nil {
		log.Println(err)
		return "", ""
	}
	log.Println("GenerateOtp 28")

	return

}

func IsOtpValid(Code, Salt string) (ok bool) {
	var err error
	log.Println("Code : ", Code, "\nSalt : ", Salt)
	ok, err = totp.ValidateCustom(Code, Salt, time.Now().UTC(), totp.ValidateOpts{
		Period:    120,
		Skew:      20,
		Digits:    4,
		Algorithm: otp.AlgorithmSHA1,
	})
	log.Println("ValidateCustom = ", ok)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
