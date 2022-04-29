package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mountainio/app/exception"
	"mountainio/domain/entity"
)

func ConnectPostgres(config Config) *gorm.DB {
	// Parse from ENV
	postgresHost := config.Get("POSTGRES_HOST")
	postgresUser := config.Get("POSTGRES_USER")
	postgresPassword := config.Get("POSTGRES_PASSWORD")
	postgresName := config.Get("POSTGRES_DB_NAME")
	postgresPort := config.Get("POSTGRES_PORT")

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
