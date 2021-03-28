package utils

const (
	Test = "test"
	Dev  = "dev"
	Prod = "prod"

	Postman            = "postman"
	PostmanIndex       = 0
	Pwa                = "app-pwa"
	PwaIndex           = 1
	Android            = "app-android"
	AndroidIndex       = 2
	AndroidBazaar      = "app-android-bazaar"
	AndroidBazaarIndex = 3
	PanelAdmin         = "panel-admin"
	PanelAdminIndex    = 4
	AndroidMyket       = "app-android-myket"
	AndroidMyketIndex  = 5
	AndroidGoogle      = "app-android-google"
	AndroidGoogleIndex = 6
	AndroidLocal       = "app-android-local"
	AndroidLocalIndex  = 7
	PanelPartner       = "panel-partner"
	PanelPartnerIndex  = 8
	Clients            = "Clients"
	UserAgent = "User-Agent"

	PgUrlDev   = "host=10.1.10.69 user=m_ramezani password=S7B-C=aUzt7h@HCv dbname=auth port=54327 sslmode=disable"
	PgUrlTest  = "host=10.1.10.59 user=m_ramezani password=j3cNeRarL72?D$nd*X%Z dbname=auth port=54327 sslmode=disable"
	PgUrlProd  = "host=10.1.10.105 user=m_ramezani password=2?BA5+H6jzmKL*q2%s*cqL2u_RXR dbname=auth port=54327 sslmode=disable"
	PgUrlLocal = "postgresql://Misagh:13750620@127.0.0.1:5432/tazminchi"
	//"host=127.0.0.1 user=Misagh password=13750620 dbname=auth port=5432 sslmode=disable" http://127.0.0.1:50093/?key=b3c47530-6939-44bd-8d37-d0698dc950c0

	EncryptingKey = `"EncryptingKey"`
	SigningKey    = `"SigningKey"`
	Issuer        = `"Issuer"`
	ClientTable   = "ClientTable"
	TokenVersion  = "bmn:tv"
	Authorization = "Authorization"
	BlockedPhones = "BlockedPhones"
	MNXClient     = "MNX-Client"
	OtpAttempts   = "OtpAttempts"
	Users         = "Users"
	UserRoles     = "UserRoles"
	Roles         = "Roles"
	XScopeToken   = "X-Scope-Token"
	IssuerHeader  = "bmn:iss"

	DomainProd = "https://napi.tazminchi.com/"
	DomainDev  = "https://dev.api.baman.club/"
	DomainTest = "https://test.api.baman.club/"

	EncryptingKeyIndex = 0
	SigningKeyIndex    = 1

	Jti        = "jti"
	Iat        = "iat"
	NameId     = "nameid"
	UniqueName = "unique_name"
	Refv       = "bmn:refv"
	Eulav      = "bmn:eulav"
	Iss        = "iss"
	Tlt        = "bmn:tlt"
	Accv       = "bmn:accv"
	Dvid       = "bmn:dvid"
	Aps        = "bmn:aps"
	Role       = "http://schemas.microsoft.com/ws/2008/06/identity/claims/role"
	Aud        = "aud"
	Exp        = "exp"
	Nbf        = "nbf"
)
