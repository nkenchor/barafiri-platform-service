package repository

import (
	"barafiri-platform-service/internal/core/helper"
	"barafiri-platform-service/internal/ports"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

type MongoRepositories struct {
	Category     ports.CategoryRepository
	Industry     ports.IndustryRepository
	Currency     ports.CurrencyRepository
	Country      ports.CountryRepository
	Notification ports.NotificationRepository
	Otp          ports.OtpRepository
	Ttl          ports.TtlRepository
}

func ConnectToMongo(dbType string, dbUsername string, dbPassword string, dbHost string,
	dbPort string, authdb string, dbname string) (MongoRepositories, error) {
	helper.LogEvent("INFO", "Establishing mongoDB connection with given credentials...")
	var mongoCredentials, authSource string
	if dbUsername != "" && dbPassword != "" {
		mongoCredentials = fmt.Sprint(dbUsername, ":", dbPassword, "@")
		authSource = fmt.Sprint("authSource=", authdb, "&")
	}
	mongoUrl := fmt.Sprint(dbType, "://", mongoCredentials, dbHost, ":", dbPort, "/?", authSource,
		"directConnection=true&serverSelectionTimeoutMS=2000")
	clientOptions := options.Client().ApplyURI(mongoUrl) // Connect to
	helper.LogEvent("INFO", "Connecting to MongoDB...")
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		//log.Println(err)
		//log.Fatal(err)
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.MongoDBError, err.Error()))
		return MongoRepositories{}, err
	}

	// Check the connection
	helper.LogEvent("INFO", "Confirming MongoDB Connection...")
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		//log.Println(err)
		//log.Fatal(err)
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.MongoDBError, err.Error()))
		return MongoRepositories{}, err
	}

	//helper.LogEvent("Info", "Connected to MongoDB!")
	helper.LogEvent("INFO", "Establishing Database collections and indexes...")
	conn := db.Database(dbname)

	categoryCollection := conn.Collection("categories")
	industryCollection := conn.Collection("industries")
	currencyCollection := conn.Collection("currencies")
	bankCollection := conn.Collection("banks")
	countryCollection := conn.Collection("countries")
	notificationCollection := conn.Collection("notification_types")
	otpCollection := conn.Collection("otp_types")
	ttlCollection := conn.Collection("ttl")

	CreateIndex(bankCollection, "name", true)
	CreateIndex(bankCollection, "code", true)
	CreateIndex(bankCollection, "swift_code", true)

	CreateIndex(countryCollection, "name", true)
	CreateIndex(countryCollection, "code", true)
	//CreateIndex(countryCollection, "dialing_code", true)
	CreateIndex(countryCollection, "iso_code_2", true)
	CreateIndex(countryCollection, "iso_code_3", true)

	CreateIndex(currencyCollection, "name", true)
	CreateIndex(currencyCollection, "code", true)
	CreateIndex(currencyCollection, "country_code", true)

	CreateIndex(categoryCollection, "name", true)

	CreateIndex(industryCollection, "name", true)
	repo := MongoRepositories{
		Category:     NewCategory(categoryCollection),
		Industry:     NewIndustry(industryCollection),
		Currency:     NewCurrency(currencyCollection),
		Country:      NewCountry(countryCollection),
		Notification: NewNotification(notificationCollection),
		Otp:          NewOtp(otpCollection),
		Ttl:          NewTtl(ttlCollection),
	}
	return repo, nil
}

// CreateIndex - creates an index for a specific field in a collection
func CreateIndex(collection *mongo.Collection, field string, unique bool) bool {

	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		helper.LogEvent("ERROR", err.Error())
		fmt.Println(err.Error())

		return false
	}
	return true
}

func GetPage(page string) (*options.FindOptions, error) {
	if page == "all" {
		return nil, nil
	}
	var limit, e = strconv.ParseInt(helper.Config.PageLimit, 10, 64)
	var pageSize, ee = strconv.ParseInt(page, 10, 64)
	if e != nil || ee != nil {
		return nil, helper.ErrorMessage(helper.NoRecordError, "Error in page-size or limit-size.")
	}
	findOptions := options.Find().SetLimit(limit).SetSkip(limit * (pageSize - 1))
	return findOptions, nil
}
