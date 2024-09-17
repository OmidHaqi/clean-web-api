package services

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/omidhaqi/clean-web-api/config"
	"github.com/omidhaqi/clean-web-api/constants"
	"github.com/omidhaqi/clean-web-api/infra/cache"
	"github.com/omidhaqi/clean-web-api/pkg/logging"
	"github.com/omidhaqi/clean-web-api/pkg/service_errors"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redisClient  *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redisClient: redis}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{Value: otp, Used: false}

	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExistsError}

	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsedError}

	}
	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)

	if err != nil {
		return err
	}
	return nil
}
func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {

	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)

	res, err := cache.Get[OtpDto](s.redisClient, key)

	if err != nil {
		return err
	} else if  res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsedError}

	} else if  !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValidError}

	} else if  !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
