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

type CountryInfra struct {
	Collection *mongo.Collection
}

func NewCountry(Collection *mongo.Collection) *CountryInfra {
	return &CountryInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.CountryRepository = &CountryInfra{}

func (r *CountryInfra) CreateCountry(country entity.Country) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting country with reference: "+country.Reference)
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"reference":    country.Reference,
			"name":         country.Name,
			"code":         country.Code,
			"dialing_code": country.DialingCode,
			"iso_code_2":   country.ISOCODE2,
			"iso_code_3":   country.ISOCODE3,
			"is_enabled":   country.IsEnabled,
		},
	)
	if err != nil {

		return nil, helper.ErrorMessage(helper.CreateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting country with reference: "+country.Reference+" completed successfully...")
	return country.Reference, nil
}

func (r *CountryInfra) UpdateCountry(reference string, country entity.Country) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting country with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":    country.Reference,
			"name":         country.Name,
			"code":         country.Code,
			"dialing_code": country.DialingCode,
			"iso_code_2":   country.ISOCODE2,
			"iso_code_3":   country.ISOCODE3,
			"is_enabled":   country.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting country with reference: "+reference+" completed successfully...")
	return country.Reference, nil
}
func (r *CountryInfra) EnableCountry(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting country with reference: "+reference)
	e, _ := r.GetCountryByRef(reference)
	country := e.(entity.Country)
	country.IsEnabled = enabled

	_, err := r.Collection.UpdateByID(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":    country.Reference,
			"name":         country.Name,
			"code":         country.Code,
			"dialing_code": country.DialingCode,
			"iso_code_2":   country.ISOCODE2,
			"iso_code_3":   country.ISOCODE3,
			"is_enabled":   country.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting country with reference: "+reference+" completed successfully...")
	return reference, nil
}

func (r *CountryInfra) GetCountryByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving country with reference: "+reference)
	country := entity.Country{} // i removed & from here

	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&country)
	if err != nil || country == (entity.Country{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving country with reference: "+reference+" completed successfully.")
	return country, nil
}

func (r *CountryInfra) GetCountryByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving country with code: "+code)
	country := entity.Country{} // i removed & from here
	filter := bson.M{"code": code}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&country)
	if err != nil || country == (entity.Country{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving country with code: "+code+" completed successfully.")
	return country, nil
}

func (r *CountryInfra) GetAllCountries(page string) (interface{}, error) {

	var countries []entity.Country
	var country entity.Country
	helper.LogEvent("INFO", "Retrieving all country entries...")
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}

	cursor, err := r.Collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&country)
		if err != nil {

			return nil, helper.ErrorMessage(helper.NoRecordError, err.Error())
		}
		countries = append(countries, country)
	}
	if reflect.ValueOf(countries).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Country{}, nil

	}
	helper.LogEvent("INFO", "Retrieving country entries completed successfully.")
	return countries, nil
}
