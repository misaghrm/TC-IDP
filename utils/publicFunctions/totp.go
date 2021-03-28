package publicFunctions

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"log"
	"time"
)

func GenerateOtp(secret string) (salt ,Code string) {
	salt = base32.StdEncoding.EncodeToString([]byte(secret))

	Code, err := totp.GenerateCodeCustom(salt, time.Now().UTC(), totp.ValidateOpts{
		Period:    140,
		Skew:      20,
		Digits:    4,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		log.Println(err)
		return "",""
	}
	return

}

func IsOtpValid (Code,Salt string) (ok bool) {
	var err error
	ok, err = totp.ValidateCustom(Code, Salt, time.Now().UTC(), totp.ValidateOpts{
		Period:    140,
		Skew:      20,
		Digits:    4,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		log.Println(err)
		return
	}
	return
}
