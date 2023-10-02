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

type CategoryInfra struct {
	Collection *mongo.Collection
}

func NewCategory(Collection *mongo.Collection) *CategoryInfra {
	return &CategoryInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.CategoryRepository = &CategoryInfra{}

func (r *CategoryInfra) CreateCategory(category entity.Category) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting category configurations with reference: "+category.Reference)
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"reference":   category.Reference,
			"name":        category.Name,
			"description": category.Description,
			"industry":    category.Industry,
			"type":        category.Type,
			"is_enabled":  category.IsEnabled,
		},
	)
	if err != nil {

		return nil, helper.ErrorMessage(helper.CreateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting category configurations with reference: "+category.Reference+" completed successfully...")
	return category.Reference, nil
}

func (r *CategoryInfra) UpdateCategory(reference string, category entity.Category) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting category configurations with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":   category.Reference,
			"name":        category.Name,
			"description": category.Description,
			"industry":    category.Industry,
			"type":        category.Type,
			"is_enabled":  category.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting category configurations with reference: "+reference+" completed successfully. ")
	return category.Reference, nil
}
func (r *CategoryInfra) EnableCategory(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling category configurations with reference: "+reference)
	e, _ := r.GetCategoryByRef(reference)
	category := e.(entity.Category)
	category.IsEnabled = enabled
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":   category.Reference,
			"name":        category.Name,
			"description": category.Description,
			"industry":    category.Industry,
			"type":        category.Type,
			"is_enabled":  category.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Enabling category configurations with reference: "+reference+" completed successfully. ")
	return reference, nil
}

func (r *CategoryInfra) GetCategoryByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving category configurations with reference: "+reference)
	category := entity.Category{} // i removed & from here

	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&category)
	if err != nil || category == (entity.Category{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving category configurations with reference: "+reference+" completed successfully. ")
	return category, nil
}

func (r *CategoryInfra) GetCategoryByName(name string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving category configurations with code: "+name)
	category := entity.Category{} // i removed & from here

	filter := bson.M{"code": name}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&category)
	if err != nil || category == (entity.Category{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving category configurations with code: "+name+" completed successfully. ")
	return category, nil
}

func (r *CategoryInfra) GetAllCategories(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving all category configuration entries...")
	var categories []entity.Category
	var category entity.Category
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}
	cursor, err := r.Collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&category)
		if err != nil {

			return nil, helper.ErrorMessage(helper.NoRecordError, err.Error())
		}
		categories = append(categories, category)
	}
	if reflect.ValueOf(categories).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Category{}, nil
	}
	helper.LogEvent("INFO", "Retrieving all category configuration entries completed successfully")
	return categories, nil
}
