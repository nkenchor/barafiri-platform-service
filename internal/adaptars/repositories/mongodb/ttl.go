package repository

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	"barafiri-platform-service/internal/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TtlInfra struct {
	Collection *mongo.Collection
}

func NewTtl(Collection *mongo.Collection) *TtlInfra {
	return &TtlInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.TtlRepository = &TtlInfra{}

func (r *TtlInfra) UpdateTtl(ttl entity.Ttl) (interface{}, error) {
	helper.LogEvent("INFO", "Updating Ttl..")
	_, _ = r.Collection.DeleteMany(context.TODO(), bson.M{})
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"time_unit": ttl.TimeUnit,
			"value":     ttl.Value,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Ttl updated successfully..")
	return struct {
		TimeUnit string `json:"time_unit"`
		Value    int    `json:"value"`
	}{string(ttl.TimeUnit), ttl.Value}, nil
}

func (r *TtlInfra) GetTtl() (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving Ttl..")
	//var ttls []entity.Ttl
	var ttl entity.Ttl
	cursor, err := r.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	cursor.Next(context.TODO())
	_ = cursor.Decode(&ttl)
	if ttl == (entity.Ttl{}) {
		helper.LogEvent("INFO", "There are no results in this collection..")
		return []entity.Ttl{}, nil
	}
	helper.LogEvent("INFO", "Retrieving Ttl completed successfully.")
	return ttl, nil
}
