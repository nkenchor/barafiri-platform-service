package repository

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"

	"barafiri-platform-service/internal/ports"
)

type OtpInfra struct {
	Collection *mongo.Collection
}

func NewOtp(Collection *mongo.Collection) *OtpInfra {
	return &OtpInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.OtpRepository = &OtpInfra{}

func (r *OtpInfra) CreateOtp(otp entity.Otp) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting otp with reference: "+otp.Reference)
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"reference":  otp.Reference,
			"code":       otp.Code,
			"is_enabled": otp.IsEnabled,
		},
	)
	if err != nil {

		return nil, helper.ErrorMessage(helper.CreateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting otp with reference: "+otp.Reference+" completed successfully...")
	return otp.Reference, nil
}

func (r *OtpInfra) UpdateOtp(reference string, otp entity.Otp) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting otp with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":  otp.Reference,
			"code":       otp.Code,
			"is_enabled": otp.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting otp with reference: "+reference+" completed successfully...")
	return otp.Reference, nil
}
func (r *OtpInfra) EnableOtp(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling otp with reference: "+reference)
	e, _ := r.GetOtpByRef(reference)
	otp := e.(entity.Otp)
	otp.IsEnabled = enabled

	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":  otp.Reference,
			"code":       otp.Code,
			"is_enabled": otp.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Enabling otp with reference: "+reference+" completed successfully...")
	return reference, nil
}

func (r *OtpInfra) GetOtpByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving otp with reference: "+reference)
	otp := entity.Otp{} // i removed & from here
	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&otp)
	if err != nil || otp == (entity.Otp{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving otp with reference: "+reference+" completed successfully...")
	return otp, nil
}

func (r *OtpInfra) GetOtpByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling otp with code: "+code)
	otp := entity.Otp{} // i removed & from here
	filter := bson.M{"code": code}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&otp)
	if err != nil || otp == (entity.Otp{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Enabling otp with code: "+code+" completed successfully...")
	return otp, nil
}

func (r *OtpInfra) GetAllOtps(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving all otp entries...")
	var otps []entity.Otp
	var otp entity.Otp

	findOptions, err := GetPage(page)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}

	cursor, err := r.Collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&otp)
		if err != nil {

			return nil, helper.ErrorMessage(helper.NoRecordError, err.Error())
		}
		otps = append(otps, otp)
	}
	if reflect.ValueOf(otps).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection")
		return []entity.Otp{}, nil
	}
	helper.LogEvent("INFO", "Retrieving all otp entries completed successfully...")
	return otps, nil
}
