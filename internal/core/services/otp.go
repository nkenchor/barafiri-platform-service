package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	ports "barafiri-platform-service/internal/ports"
	"github.com/google/uuid"
)

type otpService struct {
	otpRepository ports.OtpRepository
}

func NewOtp(
	otpRepository ports.OtpRepository,
) *otpService {
	return &otpService{

		otpRepository: otpRepository,
	}
}

func (service *otpService) CreateOtp(otp entity.Otp) (interface{}, error) {
	otp.Reference = uuid.New().String()
	helper.LogEvent("INFO", "Creating otp configuration with reference: "+otp.Reference)
	if _, err := otp.Code.CheckOtpEnum(); err != nil {
		return nil, err
	}
	if err := helper.Validate(otp); err != nil {
		return nil, err
	}
	return service.otpRepository.CreateOtp(otp)
}
func (service *otpService) UpdateOtp(reference string, otp entity.Otp) (interface{}, error) {
	helper.LogEvent("INFO", "Updating otp configuration with reference: "+reference)
	_, err := service.GetOtpByRef(reference)

	if err != nil {
		return nil, err
	}
	otp.Reference = reference
	if _, err := otp.Code.CheckOtpEnum(); err != nil {
		return nil, err
	}
	if err := helper.Validate(otp); err != nil {
		return nil, err
	}
	return service.otpRepository.UpdateOtp(reference, otp)
}
func (service *otpService) EnableOtp(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling otp configuration with reference: "+reference)
	_, err := service.GetOtpByRef(reference)
	if err != nil {
		return nil, err
	}
	return service.otpRepository.EnableOtp(reference, enabled)
}

func (service *otpService) GetOtpByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting otp configuration with reference: "+reference)
	otp, err := service.otpRepository.GetOtpByRef(reference)
	if err != nil {
		return nil, err
	}
	return otp, nil
}
func (service *otpService) GetOtpByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting otp configuration with code: "+code)
	otp, err := service.otpRepository.GetOtpByCode(code)
	if err != nil {
		return nil, err
	}
	return otp, nil
}

func (service *otpService) GetAllOtps(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting all otp entries...")
	otps, err := service.otpRepository.GetAllOtps(page)
	if err != nil {
		return nil, err
	}
	return otps, nil
}
