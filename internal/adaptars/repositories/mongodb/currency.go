package repository

import (
	"barafiri-platform-service/internal/core/domain/entity"
	"barafiri-platform-service/internal/core/helper"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
)

type CurrencyInfra struct {
	Collection *mongo.Collection
}

func NewCurrency(Collection *mongo.Collection) *CurrencyInfra {
	return &CurrencyInfra{Collection}
}
func (r *CurrencyInfra) CreateCurrency(currency entity.Currency) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting currency with reference: "+currency.Reference)
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"reference":    currency.Reference,
			"name":         currency.Name,
			"code":         currency.Code,
			"country_code": currency.CountryCode,
			"is_enabled":   currency.IsEnabled,
		},
	)
	if err != nil {

		return nil, helper.ErrorMessage(helper.CreateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisted currency with reference: "+currency.Reference+" completed successfully...")
	return currency.Reference, nil
}

func (r *CurrencyInfra) UpdateCurrency(reference string, currency entity.Currency) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting currency with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":    currency.Reference,
			"name":         currency.Name,
			"code":         currency.Code,
			"country_code": currency.CountryCode,
			"is_enabled":   currency.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting currency with reference: "+reference+" completed successfully...")
	return currency.Reference, nil
}
func (r *CurrencyInfra) EnableCurrency(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling currency with reference: "+reference)
	e, _ := r.GetCurrencyByRef(reference)
	currency := e.(entity.Currency)
	currency.IsEnabled = enabled

	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":    currency.Reference,
			"name":         currency.Name,
			"code":         currency.Code,
			"country_code": currency.CountryCode,
			"is_enabled":   currency.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Enabling currency with reference: "+reference+" completed successfully...")
	return reference, nil
}

func (r *CurrencyInfra) GetCurrencyByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving currency with reference: "+reference)
	currency := entity.Currency{} // i removed & from here

	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&currency)
	if err != nil || currency == (entity.Currency{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving currency with reference: "+reference+" completed successfully...")
	return currency, nil
}

func (r *CurrencyInfra) GetCurrencyByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving currency with code: "+code)
	currency := entity.Currency{} // i removed & from here

	filter := bson.M{"code": code}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&currency)
	if err != nil || currency == (entity.Currency{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving currency with code: "+code+" completed successfully...")
	return currency, nil
}

func (r *CurrencyInfra) GetAllCurrencies(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving all currency entries...")
	var currencies []entity.Currency
	var currency entity.Currency
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}
	cursor, err := r.Collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&currency)
		if err != nil {

			return nil, helper.ErrorMessage(helper.NoRecordError, err.Error())
		}
		currencies = append(currencies, currency)
	}
	if reflect.ValueOf(currencies).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Currency{}, nil

	}
	helper.LogEvent("INFO", "Retrieving currency entries completed successfully...")
	return currencies, nil
}
