package utils

const (
	ENV  = "ENV"
	Test = "test"
	Dev  = "dev"
	Prod = "prod"

	//Postman            = "postman"
	//PostmanIndex       = 0
	//Pwa                = "app-pwa"
	//PwaIndex           = 1
	//Android            = "app-android"
	//AndroidIndex       = 2
	//AndroidBazaar      = "app-android-bazaar"
	//AndroidBazaarIndex = 3
	//PanelAdmin         = "panel-admin"
	//PanelAdminIndex    = 4
	//AndroidMyket       = "app-android-myket"
	//AndroidMyketIndex  = 5
	//AndroidGoogle      = "app-android-google"
	//AndroidGoogleIndex = 6
	//AndroidLocal       = "app-android-local"
	//AndroidLocalIndex  = 7
	//PanelPartner       = "panel-partner"
	//PanelPartnerIndex  = 8

	Clients          = "Clients"
	UserAgent        = "User-Agent"
	OtpId            = "otpId"
	UserCities       = "UserCities"
	UsedInvitedCodes = "UsedInvitedCodes"
	UserProfiles     = "UserProfiles"
	Devices          = "Devices"
	AccessTokens     = "AccessTokens"
	RefreshTokens    = "RefreshTokens"
	ClientID         = "ClientID"

	PgUrlDev      = "host=10.1.10.69 user=m_ramezani password=S7B-C=aUzt7h@HCv dbname=auth port=54327 sslmode=disable"
	PgUrlTest     = "host=10.1.10.59 user=m_ramezani password=j3cNeRarL72?D$nd*X%Z dbname=auth port=54327 sslmode=disable"
	PgUrlProd     = "host=10.1.10.105 user=m_ramezani password=2?BA5+H6jzmKL*q2%s*cqL2u_RXR dbname=auth port=54327 sslmode=disable"
	PgUrlLocal    = "postgresql://Misagh:13750620@127.0.0.1:5432/tazminchi"
	EncryptingKey = `"EncryptingKey"`
	SigningKey    = `"SigningKey"`
	Issuer        = `"Issuer"`
	ClientTable   = "ClientTable"
	TokenVersion  = "tc:tv"
	Authorization = "Authorization"
	BlockedPhones = "BlockedPhones"
	V2            = "v1"
	ClientKey     = "client-key"
	OtpAttempts   = "OtpAttempts"
	Users         = "Users"
	UserRoles     = "UserRoles"
	Roles         = "Roles"
	IssuerHeader  = "tc:iss"
	EulaVersion   = "1"

	DomainProd = "https://napi.tazminchi.com/"
	DomainDev  = "https://dev.tazminchi.com/"
	DomainTest = "https://test.tazminchi.com/"

	EncryptingKeyIndex = 0
	SigningKeyIndex    = 1

	Jti        = "jti"
	Iat        = "iat"
	NameId     = "nameid"
	UniqueName = "unique_name"
	Refv       = "tc:refv"
	Eulav      = "tc:eulav"
	Iss        = "iss"
	Tlt        = "tc:tlt"
	Accv       = "tc:accv"
	Dvid       = "tc:dvid"
	Aps        = "tc:aps"
	Aud        = "aud"
	Exp        = "exp"
	Nbf        = "nbf"
)
