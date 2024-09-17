package services

import (
	"github.com/omidhaqi/clean-web-api/api/dto"
	"github.com/omidhaqi/clean-web-api/common"
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/infra/persistence/database"
	"github.com/omidhaqi/clean-web-api/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     *logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := database.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		cfg: cfg,
		database: database,
		logger: &logger,
		otpService: NewOtpService(cfg),
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {

	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber,otp)
	if err != nil {
		return err
	}
	return nil 
	
}
