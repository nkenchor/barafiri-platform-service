package entity

import "barafiri-platform-service/internal/core/domain/shared"

type Ttl struct {
	TimeUnit shared.TimeUnit `bson:"time_unit" json:"time_unit" validate:"required,min=3"`
	Value    int             `bson:"value" json:"value" validate:"required,numeric,min=1"`
}
