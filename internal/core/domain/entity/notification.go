package entity

import "barafiri-platform-service/internal/core/domain/shared"

type Notification struct {
	Reference string                  `bson:"reference" json:"reference"`
	Code      shared.NotificationType `bson:"code" json:"code" validate:"required,min=3"`
	IsEnabled bool                    `bson:"is_enabled" json:"is_enabled"`
}
