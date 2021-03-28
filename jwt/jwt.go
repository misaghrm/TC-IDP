package jwt

import (
	"encoding/base64"
	"fmt"
	jose "github.com/dvsekhvalnov/jose2go"
	//"golang.org/x/oauth2/jwt"
	"log"
	db "tc-micro-idp/dataManager"
	"tc-micro-idp/models"
	. "tc-micro-idp/utils"
)

var ClientsTable []*models.Client

func init() {

	var OK bool

	ClientsTable, OK = db.GetClientsTable()
	if !OK {
		log.Println("Client table does not exist in cache")
	}
	log.Println(ClientsTable)

	for _, client := range ClientsTable {
		var Key [][]byte
		e, s := SigningAndEncryptionKeyFinder(client.Issuer)
		Key = append(Key, e, s)
		db.Cache.Set(client.Issuer, Key, 1)
	}
}

func Decrypt(token, issuer string) (tokenClaims *models.TokenClaim, err error) {
	if token == "" {
		return nil, fmt.Errorf("no valid token")
	}
	return decrypt(token, issuer)
}

func decrypt(token, issuer string) (tokenClaims *models.TokenClaim, err error) {
	var newVersion bool
	var signingKey, encryptionKey []byte
	key, found := db.Cache.Get(issuer)
	if !found {
		return nil, fmt.Errorf("the %s issuer does not exist", issuer)
	}

	a, b, err := jose.Decode(token, func(header map[string]interface{}, payload string) interface{} {
		if header[IssuerHeader] != issuer {
			return nil
		}
		encryptionKey, signingKey = (key).([][]byte)[EncryptingKeyIndex], (key).([][]byte)[SigningKeyIndex]
		log.Println("encryptionKey:", encryptionKey)
		if i, ok := header[TokenVersion]; ok {
			if i.(string) == "v2" {
				newVersion = true
				return encryptionKey
			}
		}
		token, _, err = jose.Decode(token, encryptionKey)
		log.Println("error decryptttttt:", err, "\n", token)
		if err != nil {
			log.Println("error decrypt:", err, "\n", token)
			return encryptionKey
		}
		return encryptionKey
	})
	log.Println("a,b : ", a, "\n", b)
	log.Println("decrypted Token : ", token)
	if !newVersion {
		tokenClaims, err = decode(token, signingKey)
		log.Println("error decrypt:", err)
	}
	//tokenClaims, err = verifyToken(a, issuer)
	if err != nil {
		log.Println(err)
	}
	return
}

func decode(decryptedToken string, SigningKey []byte) (tokenClaims *models.TokenClaim, err error) {
	if len(decryptedToken) <= 0 {
		return nil, fmt.Errorf("token is invalidd %s", decryptedToken)
	}
	//return ExtractTokenMetadata(decryptedToken[3:], SigningKey)
	return
}

func SigningAndEncryptionKeyFinder(h string) (encryptingKey, signingKey []byte) {
	var err error
	ClientsTable, _ = db.GetClientsTable()
	for _, client := range ClientsTable {
		if h == client.Issuer {
			signingKey, err = base64.StdEncoding.DecodeString(client.SigningKey)
			if err != nil {
				return nil, nil
			}
			encryptingKey, err = base64.StdEncoding.DecodeString(client.EncryptingKey)
			if err != nil {
				return nil, nil
			}
			return
		}
	}
	return
}

func GenerateToken(payload []byte) string {

	encryptingKey, err := base64.StdEncoding.DecodeString(ClientsTable[PwaIndex].EncryptingKey)
	EncryptedToken, err := jose.EncryptBytes(payload, jose.A128KW, jose.A128CBC_HS256, encryptingKey, jose.Zip(jose.DEF), jose.Headers(map[string]interface{}{"typ": "JWT", "bmn:iss": "app-pwa", TokenVersion: "v2"}))
	if err != nil {
		log.Println(err)
		return ""
	}
	log.Println(EncryptedToken)
	return EncryptedToken
}
