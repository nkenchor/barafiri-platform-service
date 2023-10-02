package services

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	port "barafiri-platform-service/internal/ports"
	"github.com/google/uuid"
)

type notificationService struct {
	notificationRepository port.NotificationRepository
}

func NewNotification(
	notificationRepository port.NotificationRepository,
) *notificationService {
	return &notificationService{

		notificationRepository: notificationRepository,
	}
}
func (service *notificationService) CreateNotification(notification entity.Notification) (interface{}, error) {
	notification.Reference = uuid.New().String()
	helper.LogEvent("INFO", "Creating notification configuration with reference: "+notification.Reference)
	if _, err := notification.Code.CheckNotificationEnum(); err != nil {
		return nil, err
	}
	if err := helper.Validate(notification); err != nil {
		return nil, err
	}
	return service.notificationRepository.CreateNotification(notification)
}

func (service *notificationService) UpdateNotification(reference string, notification entity.Notification) (interface{}, error) {
	helper.LogEvent("INFO", "Updating notification configuration with reference: "+reference)
	_, err := service.GetNotificationByRef(reference)
	notification.Reference = reference

	if err != nil {
		return nil, err
	}
	if _, err := notification.Code.CheckNotificationEnum(); err != nil {
		return nil, err
	}
	if err := helper.Validate(notification); err != nil {
		return nil, err
	}
	return service.notificationRepository.UpdateNotification(reference, notification)
}
func (service *notificationService) EnableNotification(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling notification configuration with reference: "+reference)
	_, err := service.GetNotificationByRef(reference)
	if err != nil {
		return nil, err
	}
	return service.notificationRepository.EnableNotification(reference, enabled)
}

func (service *notificationService) GetNotificationByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting notification configuration with reference: "+reference)
	notification, err := service.notificationRepository.GetNotificationByRef(reference)
	if err != nil {
		return nil, err
	}
	return notification, nil
}
func (service *notificationService) GetNotificationByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Updating notification configuration with code: "+code)
	notification, err := service.notificationRepository.GetNotificationByCode(code)
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (service *notificationService) GetAllNotifications(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Getting all notification configurations...")
	notifications, err := service.notificationRepository.GetAllNotifications(page)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
