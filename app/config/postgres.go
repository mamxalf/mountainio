package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mountainio/app/exception"
	"mountainio/domain/entity"
	"strconv"
)

func ConnectPostgres(config Config) *gorm.DB {
	// Parse from ENV
	postgresHost, err := strconv.Atoi(config.Get("POSTGRES_HOST"))
	exception.PanicIfNeeded(err)
	postgresUser, err := strconv.Atoi(config.Get("POSTGRES_USER"))
	exception.PanicIfNeeded(err)
	postgresPassword, err := strconv.Atoi(config.Get("POSTGRES_PASSWORD"))
	exception.PanicIfNeeded(err)
	postgresName, err := strconv.Atoi(config.Get("POSTGRES_DB_NAME"))
	exception.PanicIfNeeded(err)
	postgresPort, err := strconv.Atoi(config.Get("POSTGRES_PORT"))
	exception.PanicIfNeeded(err)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		postgresHost, postgresUser, postgresPassword, postgresName, postgresPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	return db
}
