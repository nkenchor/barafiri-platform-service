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

type NotificationInfra struct {
	Collection *mongo.Collection
}

func NewNotification(Collection *mongo.Collection) *NotificationInfra {
	return &NotificationInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.NotificationRepository = &NotificationInfra{}

func (r *NotificationInfra) CreateNotification(notification entity.Notification) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting notification with reference: "+notification.Reference)
	_, err := r.Collection.InsertOne(
		context.TODO(),

		bson.M{
			"reference":  notification.Reference,
			"code":       notification.Code,
			"is_enabled": notification.IsEnabled,
		},
	)
	if err != nil {

		return nil, helper.ErrorMessage(helper.CreateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting notification with reference: "+notification.Reference+" completed successfully..")
	return notification.Reference, nil
}

func (r *NotificationInfra) UpdateNotification(reference string, notification entity.Notification) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting notification with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":  notification.Reference,
			"code":       notification.Code,
			"is_enabled": notification.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Persisting notification with reference: "+reference+" completed successfully...")
	return notification.Reference, nil
}

func (r *NotificationInfra) EnableNotification(reference string, enabled bool) (interface{}, error) {
	helper.LogEvent("INFO", "Enabling notification with reference: "+reference)
	e, _ := r.GetNotificationByRef(reference)
	notification := e.(entity.Notification)
	notification.IsEnabled = enabled

	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": reference},
		bson.M{
			"reference":  notification.Reference,
			"code":       notification.Code,
			"is_enabled": notification.IsEnabled,
		},
	)
	if err != nil {
		return nil, helper.ErrorMessage(helper.UpdateError, err.Error())
	}
	helper.LogEvent("INFO", "Enabling notification with reference: "+reference+" completed successfully...")
	return reference, nil
}

func (r *NotificationInfra) GetNotificationByRef(reference string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving notification with reference: "+reference)
	notification := entity.Notification{} // i removed & from here

	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil || notification == (entity.Notification{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Retrieving notification with reference: "+reference+" completed successfully...")
	return notification, nil
}

func (r *NotificationInfra) GetNotificationByCode(code string) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting notification with code: "+code)
	notification := entity.Notification{} // i removed & from here

	filter := bson.M{"code": code}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil || notification == (entity.Notification{}) {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	helper.LogEvent("INFO", "Persisting notification with code: "+code+" completed successfully...")
	return notification, nil
}

func (r *NotificationInfra) GetAllNotifications(page string) (interface{}, error) {
	helper.LogEvent("INFO", "Retrieving all notifications...")
	var notifications []entity.Notification
	var notification entity.Notification
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}
	cursor, err := r.Collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, helper.NoRecordFound)
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&notification)
		if err != nil {

			return nil, helper.ErrorMessage(helper.NoRecordError, err.Error())
		}
		notifications = append(notifications, notification)
	}
	if reflect.ValueOf(notifications).IsNil() {
		helper.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Notification{}, nil
	}
	helper.LogEvent("INFO", "Persisting notifications")
	return notifications, nil
}
