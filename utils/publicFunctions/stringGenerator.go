package publicFunctions

import (
	"crypto/rand"
	"fmt"
)

func StringGenerator(c int) (str string) {
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return string(b)
}
