package logging

type Category string
type SubCategory string
type ExtraKey string

const (
	General         Category = "General"
	Internal        Category = "Internal"
	Postgres        Category = "Postgres"
	Redis           Category = "Redis"
	Otp             Category = "Otp"
	User            Category = "User"
	Token           Category = "Token"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
)

const (
	Startup             SubCategory = "Startup"
	Add                 SubCategory = "Add"
	ExternalService     SubCategory = "ExternalService"
	Migration           SubCategory = "Migration"
	New                 SubCategory = "New"
	Edit                SubCategory = "Edit"
	Select              SubCategory = "Select"
	Rollback            SubCategory = "Rollback"
	Update              SubCategory = "Update"
	Delete              SubCategory = "Delete"
	Insert              SubCategory = "Insert"
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"
	MobileValidation    SubCategory = "MobileValidation"
	PasswordValidation  SubCategory = "PasswordValidation"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "LoggerName"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	Body         ExtraKey = "Body"
	ErrorMessage ExtraKey = "ErrorMessage"
)
