package publicFunctions

import (
	"crypto/rand"
	"fmt"
	"log"
)

func StringGenerator(c string) (str string) {
	log.Println("StringGenerator 10")
	b := make([]byte, len(c))
	_, err := rand.Read(b)
	log.Println("StringGenerator 13")
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	log.Println("StringGenerator 18")
	return string(b)
}
