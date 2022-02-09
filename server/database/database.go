package database

import (
	"log"
	"os"

	"github.com/GDSC-UIT/job-connection-api/conf"
	"github.com/GDSC-UIT/job-connection-api/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

var DBInstance Database

func ConnectDb() {
	db,err := gorm.Open(postgres.Open(conf.Config.DatabaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if (err !=nil) {
		log.Fatal("Fail to connect database. \n",err)
		os.Exit(2)
	}

	log.Println("Database connected")
	db.AutoMigrate(&models.UserProfile{})

	DBInstance = Database{
		Db: db,
	}
}