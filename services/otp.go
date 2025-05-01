package services

import (
	"digimovie/src/config"
	"digimovie/src/database"
	"digimovie/src/logging"
	"fmt"
	"math/rand"
	"time"
)

func MakeOtp() string {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	otp := rand.Intn(max - min) + min
	return fmt.Sprint(otp)
}

type OtpService struct {
	cfg *config.Config
	logger logging.Logger
}

func NewOtpService(cfg *config.Config) *OtpService {
	return &OtpService{
		cfg: cfg,
		logger: logging.NewLogger(),
	}
}

type OtpDto struct {
	Value string `josn:"value"`
	Valid bool `json:"valid"`
}

func (s *OtpService) SetOtp(mobileNumber string, otp string, duration time.Duration) error {
	res, err := database.Get[OtpDto](mobileNumber)
	if err == nil && !res.Valid {
		return fmt.Errorf("otp used")
	} else if err == nil && res.Valid {
		return fmt.Errorf("otp exists")
	}
	err = database.Set(mobileNumber, otp, duration)
	if err != nil {
		return err
	}
	s.logger.Info(logging.Otp, logging.Add, "new otp added to redis", nil)
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	res, err := database.Get[OtpDto](mobileNumber)
	if err != nil {
		return fmt.Errorf("doesnt exists")
	} else if !res.Valid {
		return fmt.Errorf("otp used")
	} else if res.Valid && res.Value != otp {
		return fmt.Errorf("invalid otp")
	} else if res.Valid && res.Value == otp {
		res.Valid = false
		err = database.Set(mobileNumber, res, 5)
		if err != nil {
			return err
		}
	}
	return nil
}