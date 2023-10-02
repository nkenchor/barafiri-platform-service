package main

import (
	mongoRepository "barafiri-platform-service/internal/adaptars/repositories/mongodb"
	redisRepository "barafiri-platform-service/internal/adaptars/repositories/redis"
	"barafiri-platform-service/internal/adaptars/routes"
	"barafiri-platform-service/internal/core/helper"
	"fmt"
	"log"
)

func main() {
	//Initialize request Log
	helper.InitializeLog()
	//Start DB Connection
	mongoRepo := startMongo()
	helper.LogEvent("INFO", "MongoDB Initialized!")
	redisRepo := startRedis()
	helper.LogEvent("INFO", "Redis Initialized!")
	//Set up routes
	router := routes.SetupRouter(mongoRepo.Country, mongoRepo.Currency, mongoRepo.Notification,
		mongoRepo.Otp, mongoRepo.Category, mongoRepo.Industry, mongoRepo.Ttl, redisRepo.Configuration)
	//Print custom message for server start

	fmt.Println(helper.ServerStarted)
	//Log server start event
	helper.LogEvent("INFO", helper.ServerStarted)
	//start server
	_ = router.Run(":" + helper.Config.ServicePort)
	//api.SetConfiguration
}

func startMongo() mongoRepository.MongoRepositories {
	helper.LogEvent("INFO", "Initializing Mongo!")
	mongoRepo, err := mongoRepository.ConnectToMongo(helper.Config.DbType, helper.Config.MongoDbUserName,
		helper.Config.MongoDbPassword, helper.Config.MongoDbHost,
		helper.Config.MongoDbPort, helper.Config.MongoDbAuthDb,
		helper.Config.MongoDbName)
	if err != nil {
		fmt.Println(err)
		helper.LogEvent("ERROR", "MongoDB database Connection Error: "+err.Error())
		log.Fatal()
	}
	return mongoRepo
}

func startRedis() redisRepository.RedisRepositories {
	helper.LogEvent("INFO", "Initializing Redis!")
	redisRepo, err := redisRepository.ConnectToRedis(helper.Config.RedisHost, helper.Config.RedisPort)
	if err != nil {
		fmt.Println(err)
		helper.LogEvent("ERROR", "Redis Initialization Error: "+err.Error())
		log.Fatal()
	}
	return redisRepo
}
