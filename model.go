package models

import (
	"database/sql"
	"encoding/json"
	"strings"
	//"tc-micro-idp/utils"
	"time"
)

var Domain string

func init() {
	//switch os.Getenv("ENV") {
	//case utils.Test:
	//	Domain = utils.DomainTest
	//case utils.Dev:
	//	Domain = utils.DomainDev
	//case utils.Prod:
	//	Domain = utils.DomainProd
	//default:
	//	Domain = utils.DomainProd
	//}
}

type ChallengeInput struct {
	Phone            string `json:"phone"`
	AppSignatureHash string `json:"appSignatureHash"`
	InviteCode       string `json:"inviteCode"`
	RecaptchaToken   string `json:"recaptchaToken"`
}

type ResponseModel struct {
	Data    map[string]interface{} `json:"data"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
}

type BlockedPhone struct {
	Id           int64     `gorm:"uniqueIndex;column:Id"`
	CreationTime time.Time `gorm:"column:CreationTime"`
	ModifyTime   time.Time `gorm:"column:ModifyTime"`
	Number       string    `gorm:"column:Number;type:bpchar(10);"`
}

func (BlockedPhone) TableName() string {
	return "BlockedPhones"
}

type Client struct {
	Id                       int64     `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime             time.Time `gorm:"column:CreationTime"`
	ModifyTime               time.Time `gorm:"column:ModifyTime"`
	Issuer                   string    `gorm:"column:Issuer"`
	Audience                 string    `gorm:"column:Audience"`
	ValidateAudience         bool      `gorm:"column:ValidateAudience"`
	ValidateIssuer           bool      `gorm:"column:ValidateIssuer"`
	ValidateIssuerSigningKey bool      `gorm:"column:ValidateIssuerSigningKey"`
	ValidateLifetime         bool      `gorm:"column:ValidateLifetime"`
	CanRegister              bool      `gorm:"column:CanRegister"`
	CanLogin                 bool      `gorm:"column:CanLogin"`
	Alg                      string    `gorm:"column:Alg"`
	Enc                      string    `gorm:"column:Enc"`
	AccessTokenLifeTime      string    `gorm:"column:AccessTokenLifeTime"`
	RefreshTokenLifeTime     string    `gorm:"column:RefreshTokenLifeTime"`
	SupportCompression       bool      `gorm:"column:SupportCompression"`
	SigningKey               string    `gorm:"column:SigningKey"`
	EncryptingKey            string    `gorm:"column:EncryptingKey"`
	RequiredRoles            string    `gorm:"column:RequiredRoles;type:text[]"`
	DefaultRoles             string    `gorm:"column:DefaultRoles;type:text[]"`
}

func (Client) TableName() string {
	return "Clients"
}

type OtpAttempt struct {
	Id           int64         `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime time.Time     `gorm:"column:CreationTime"`
	ModifyTime   *time.Time    `gorm:"column:ModifyTime"`
	UserId       sql.NullInt64 `gorm:"column:UserId"`
	User         *User         `gorm:"foreignKey:UserId"`
	Phone        string        `gorm:"column:Phone"`
	ClientId     int64         `gorm:"column:ClientId"`
	Client       Client        `gorm:"foreignKey:ClientId"`
	Salt         string        `gorm:"column:Salt"`
	IssueTime    time.Time     `gorm:"column:IssueTime"`
	ExpireTime   time.Time     `gorm:"column:ExpireTime"`
	UserIp       string        `gorm:"column:UserIp"`
	UserAgent    string        `gorm:"column:UserAgent"`
	OtpKind      Kind          `gorm:"column:Kind"`
}

func (OtpAttempt) TableName() string {
	return "OtpAttempts"
}

type Kind byte

var OtpKinds = []string{"Register", "Login"}

func (k Kind) String() string {
	switch k {
	case 2:
		return OtpKinds[0]
	case 4:
		return OtpKinds[1]
	default:
		return ""
	}
}

type AppSource int

var appSources = [...]string{"Undefined", "Bazaar", "Google", "Local", "Myket"}

func (appsource AppSource) String() string {
	switch appsource {
	case 0:
		return appSources[0]
	case 2:
		return appSources[1]
	case 4:
		return appSources[2]
	case 8:
		return appSources[3]
	case 16:
		return appSources[4]
	default:
		return appSources[0]
	}
}

func (TokenClaim) TableName() string {
	return "TokenClaims"
}

func GetJson(a *TokenClaim) (payload []byte) {
	payload, _ = json.Marshal(a)
	return
}

//func (t *TokenClaim) GetRoles() []string {
//	s := make([]string, len(t.Roles))
//	for i, v := range t.Roles {
//		s[i] = fmt.Sprintf("%v", v)
//	}
//	return s
//}

type UserRole struct {
	UserId int64 `gorm:"primaryKey;autoIncrement:false;column:UserId;"`
	RoleId int64 `gorm:"primaryKey;autoIncrement:false;column:RoleId;"`
}

func (UserRole) TableName() string {
	return "UserRoles"
}

type Role struct {
	Id           int64      `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime time.Time  `gorm:"column:CreationTime"`
	ModifyTime   *time.Time `gorm:"column:ModifyTime"`
	UserRoles    []UserRole //`gorm:"many2many:UserRoles;"`
	Name         string     `gorm:"column:Name;" ;json:"name"`
	Title        string     `gorm:"column:Title;" ;json:"title"`
	IsInHouse    bool       `gorm:"column:IsInHouse;" ;json:"isInHouse"`
	Visible      bool       `gorm:"column:Visible;" ;json:"visible"`
}

func (Role) TableName() string {
	return "Roles"
}

type User struct {
	Id            int64          `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime  time.Time      `gorm:"column:CreationTime"`
	ModifyTime    *time.Time     `gorm:"column:ModifyTime"`
	PhoneNumber   string         `gorm:"column:PhoneNumber;type:bpchar(10);" ;json:"phoneNumber"`
	UserProfileId int64          `gorm:"column:UserProfileId" ;json:"userProfileId"`
	LastLoginTime time.Time      `gorm:"column:LastLoginTime" ;json:"lastLoginTime"`
	UserRoles     []UserRole     //`gorm:"many2many:UserRoles;"`
	RefreshTokens []RefreshToken //`gorm:"many2many:RefreshTokens" ;json:"refreshTokens"`
	OtpAttempts   []*OtpAttempt  //`gorm:"many2many:OtpAttempts" ;json:"otpAttempts"`
	UserCity      UserCity       //`gorm:"foreignKey:FK_UserCities_Users_UserId"`
}

func (User) TableName() string {
	return "Users"
}

type RefreshToken struct {
	Id           int64         `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id;"`
	CreationTime time.Time     `gorm:"column:CreationTime"`
	ModifyTime   time.Time     `gorm:"column:ModifyTime"`
	UserId       int64         `gorm:"column:UserId" ;json:"userId"`
	User         User          `gorm:"foreignKey:UserId"`
	ClientId     int64         `gorm:"column:ClientId" ;json:"clientId"`
	Client       Client        `gorm:"foreignKey:ClientId"`
	Token        string        `gorm:"column:Token" ;json:"token"`
	IssueTime    time.Time     `gorm:"column:IssueTime" ;json:"issueTime"`
	ExpireTime   time.Time     `gorm:"column:ExpireTime" ;json:"expireTime"`
	AccessTokens []AccessToken //`gorm:"many2many:AccessTokens;"`
	DeviceId     int64         `gorm:"column:DeviceId" ;json:"deviceId"`
	Device       Device        `gorm:"foreignKey:DeviceId"`
	IsRevoked    bool          `gorm:"column:IsRevoked" ;json:"isRevoked"`
	RevokeTime   *time.Time    `gorm:"column:RevokeTime" ;json:"revokeTime"`
}

func (RefreshToken) TableName() string {
	return "RefreshTokens"
}

type AccessToken struct {
	Id             int64        `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime   time.Time    `gorm:"column:CreationTime"`
	ModifyTime     *time.Time   `gorm:"column:ModifyTime"`
	RefreshTokenId int64        `gorm:"column:RefreshTokenId" ;json:"refreshTokenId"`
	RefreshToken   RefreshToken `gorm:"foreignKey:RefreshTokenId"`
	Token          string       `gorm:"column:Token" ;json:"token"`
	IssueTime      time.Time    `gorm:"column:IssueTime" ;json:"issueTime"`
	ExpireTime     time.Time    `gorm:"column:ExpireTime" ;json:"expireTime"`
	IsRevoked      bool         `gorm:"column:IsRevoked" ;json:"isRevoked"`
	RevokeTime     *time.Time   `gorm:"column:RevokeTime" ;json:"revokeTime"`
}

func (AccessToken) TableName() string {
	return "AccessTokens"
}

type Device struct {
	Id             int64          `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime   time.Time      `gorm:"column:CreationTime"`
	ModifyTime     *time.Time     `gorm:"column:ModifyTime"`
	RefreshTokenId int64          `gorm:"column:RefreshTokenId" ;json:"RefreshTokenId"`
	RefreshToken   []RefreshToken `gorm:"foreignKey:Id"`
	UserIp         string         `gorm:"column:UserIp" ;json:"userIp"`
	UserAgent      string         `gorm:"column:UserAgent" ;json:"userAgent"`
	FireBaseId     string         `gorm:"column:FireBaseId" ;json:"fireBaseId"`
	YandexId       string         `gorm:"column:YandexId" ;json:"yandexId"`
	Imei           string         `gorm:"column:Imei" ;json:"imei"`
	PhoneModel     string         `gorm:"column:PhoneModel" ;json:"phoneModel"`
	AndroidVersion string         `gorm:"column:AndroidVersion" ;json:"androidVersion"`
	ScreenSize     string         `gorm:"column:ScreenSize" ;json:"screenSize"`
	AppVersion     string         `gorm:"column:AppVersion" ;json:"appVersion"`
	SimOperator    string         `gorm:"column:SimOperator" ;json:"simOperator"`
	AppSource      AppSource      `gorm:"column:AppSource" ;json:"appSource"`
}

func (Device) TableName() string {
	return "Devices"
}

type UserProfile struct {
	Id                 int64      `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime       time.Time  `gorm:"column:CreationTime"`
	ModifyTime         *time.Time `gorm:"column:ModifyTime"`
	UserId             int64      `gorm:"column:UserId"`
	User               User       `gorm:"foreignKey:UserId"`
	FirstName          string     `gorm:"column:FirstName"`
	LastName           string     `gorm:"column:LastName"`
	Gender             Gender     `gorm:"column:Gender"`
	BirthDate          *time.Time `gorm:"column:BirthDate"`
	JobTitle           string     `gorm:"column:JobTitle"`
	Email              string     `gorm:"column:Email"`
	ProfileImageFileId int64      `gorm:"column:ProfileImageFileId"`
	InviteCode         string     `gorm:"column:InviteCode"`
	Address            string     `gorm:"column:Address"`
	PostalCode         string     `gorm:"column:PostalCode"`
	Latitude           float64    `gorm:"column:Latitude"`
	Longitude          float64    `gorm:"column:Longitude"`
}

func (UserProfile) TableName() string {
	return "UserProfiles"
}

type Gender byte

var genderKinds = []string{"Male", "Female"}

func (g Gender) String() string {
	switch g {
	case 1:
		return genderKinds[0]
	case 2:
		return genderKinds[1]
	default:
		return ""
	}
}

type UsedInviteCode struct {
	InviterUserId               int64      `gorm:"column:InviterUserId"`
	RequestDate                 time.Time  `gorm:"column:RequestDate"`
	InvitedPhoneNumber          string     `gorm:"column:InvitedPhoneNumber"`
	InvitedUserId               int64      `gorm:"column:InvitedUserId"`
	VerifyDate                  *time.Time `gorm:"column:VerifyDate"`
	IsVerified                  bool       `gorm:"column:IsVerified"`
	InvitedFirstPurchaseMade    bool       `gorm:"column:InvitedFirstPurchaseMade"`
	InvitedFirstInvoiceId       int64      `gorm:"column:InvitedFirstInvoiceId"`
	InvitedFirstPurchaseTimeUtc *time.Time `gorm:"column:InvitedFirstPurchaseTimeUtc"`
	InviterBonusSent            bool       `gorm:"column:InviterBonusSent"`
	InviterBonusSentAmount      float64    `gorm:"column:InviterBonusSentAmount"`
	InviterBonusSentTimeUtc     *time.Time `gorm:"column:InviterBonusSentTimeUtc"`
}

func (UsedInviteCode) TableName() string {
	return "UsedInvitedCodes"
}

type UserCity struct {
	Id           int64      `gorm:"uniqueIndex;primaryKey;autoIncrement:false;column:Id"`
	CreationTime time.Time  `gorm:"column:CreationTime"`
	ModifyTime   *time.Time `gorm:"column:ModifyTime"`
	UserId       int64      `gorm:"primaryKey;autoIncrement:false;column:UserId;"`
	CityId       int64      `gorm:"primaryKey;autoIncrement:false;column:CityId"`
}

func (UserCity) TableName() string {
	return "UserCities"
}

// Trim Domain trims the domain prefix and query params if exists and implemented for Authorizer.pb.go
func (x *Request) TrimDomain() {
	x.URL = strings.TrimPrefix(x.URL, Domain)
	x.URL = strings.Split(x.URL, "?")[0]
}
