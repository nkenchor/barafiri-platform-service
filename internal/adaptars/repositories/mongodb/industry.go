package repository

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	"barafiri-platform-service/internal/ports"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type IndustryInfra struct {
	Collection *mongo.Collection
}

func NewIndustry(Collection *mongo.Collection) *IndustryInfra {
	return &IndustryInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.IndustryRepository = &IndustryInfra{}

func (r *IndustryInfra) CreateIndustry(industry entity.Industry) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting industry configurations with reference: "+industry.Reference)
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"reference":   industry.Reference,
			"name":        industry.Name,
			"description": industry.Description,
			"is_enabled":  industry.IsEnabled,
		},
	)
	if err != nil {

		return nil, helper.ErrorMessage(helper.CreateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting industry configurations with reference: "+industry.Reference+" completed successfully...")
	return industry.Reference, nil
}

func (r *IndustryInfra) UpdateIndustry(reference string, industry entity.Industry) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting industry configurations with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":   industry.Reference,
			"name":        industry.Name,
			"description": industry.Description,
			"is_enabled":  industry.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting industry configurations with reference: "+reference+" completed successfully. ")
	return industry.Reference, nil
}
func (r *IndustryInfra) EnableIndustry(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling industry configurations with reference: "+reference)
	e, _ := r.GetIndustryByRef(reference)
	industry := e.(entity.Industry)
	industry.IsEnabled = enabled
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":   industry.Reference,
			"name":        industry.Name,
			"description": industry.Description,
			"is_enabled":  industry.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Enabling industry configurations with reference: "+reference+" completed successfully. ")
	return reference, nil
}

func (r *IndustryInfra) GetIndustryByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving industry configurations with reference: "+reference)
	industry := entity.Industry{} // i removed & from here

	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&industry)
	if err != nil || industry == (entity.Industry{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving industry configurations with reference: "+reference+" completed successfully. ")
	return industry, nil
}

func (r *IndustryInfra) GetIndustryByName(name string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving industry configurations with code: "+name)
	industry := entity.Industry{} // i removed & from here

	filter := bson.M{"code": name}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&industry)
	if err != nil || industry == (entity.Industry{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving industry configurations with code: "+name+" completed successfully. ")
	return industry, nil
}

func (r *IndustryInfra) GetAllIndustries(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving all industry configuration entries...")
	var industrys []entity.Industry
	var industry entity.Industry
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}
	cursor, err := r.Collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&industry)
		if err != nil {

			return nil, helper.ErrorMessage(helper.NoRecordError, err.Error())
		}
		industrys = append(industrys, industry)
	}
	if reflect.ValueOf(industrys).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Industry{}, nil
	}
	helper.LogEvent("INFO", "Retrieving all industry configuration entries completed successfully")
	return industrys, nil
}
