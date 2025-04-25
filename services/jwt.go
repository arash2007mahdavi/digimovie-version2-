package services

import (
	"digimovie/src/config"
	"digimovie/src/logging"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
	cfg    *config.Config
	logger logging.Logger
}

func NewJwtService(cfg *config.Config) *JwtService {
	return &JwtService{
		cfg:    cfg,
		logger: logging.NewLogger(),
	}
}

type JwtDto struct {
	Id           string
	Firstname    string
	Lastname     string
	Username     string
	MobileNumber string
	Email        string
	Enabled      bool
	Roles        []string
}

type JwtAccess struct {
	AccessToken string
	AccessExpireTime int
}

func (s JwtService) GenerateToken(dto JwtDto, expireTime time.Duration) (*JwtAccess, error) {
	access := &JwtAccess{}
	access.AccessExpireTime = int(time.Now().Add(expireTime * time.Minute).Unix())

	claims := jwt.MapClaims{}
	claims["id"] = dto.Id
	claims["firstname"] = dto.Firstname
	claims["lastname"] = dto.Lastname
	claims["username"] = dto.Username
	claims["mobilenumber"] = dto.MobileNumber
	claims["email"] = dto.Email
	claims["enables"] = dto.Enabled
	claims["roles"] = dto.Roles

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	access.AccessToken, err = at.SignedString([]byte("@rash2007"))
	if err != nil {
		return nil, err
	}
	return access, nil
}