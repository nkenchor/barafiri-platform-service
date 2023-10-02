package shared

import (
	"barafiri-platform-service/internal/core/helper"
)

func (c NotificationType) CheckNotificationEnum() (string, error) {
	helper.LogEvent("INFO", "Validating notification type...")
	var response string
	notificationtype := [...]string{"RING_AND_VIBRATE", "RING_ONLY", "VIBRATE_ONLY"}

	x := string(c)
	for _, v := range notificationtype {
		if v == x {

			response = x

			return response, nil
		}
	}

	response = ""

	return response, helper.ErrorMessage(helper.ValidationError, "invalid notification code")
}

func (c OtpType) CheckOtpEnum() (string, error) {
	helper.LogEvent("INFO", "Validating otp type...")
	var response string
	otptype := [...]string{"SEND_AS_SMS_ONLY", "SEND_AS_EMAIL_ONLY", "SEND_AS_SMS_AND_EMAIL"}

	x := string(c)
	for _, v := range otptype {
		if v == x {

			response = x

			return response, nil
		}
	}

	response = ""
	return response, helper.ErrorMessage(helper.ValidationError, "invalid otp code")
}
func (c TimeUnit) CheckTimeUnitEnum() (string, error) {
	helper.LogEvent("INFO", "Validating time unit...")
	var response string
	otptype := [...]string{"HOURS", "MINUTES", "SECONDS"}

	x := string(c)
	for _, v := range otptype {
		if v == x {

			response = x

			return response, nil
		}
	}

	response = ""
	return response, helper.ErrorMessage(helper.ValidationError, "invalid time unit")
}
