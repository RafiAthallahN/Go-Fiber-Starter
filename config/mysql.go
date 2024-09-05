package config

import (
	"fmt"
	"go-fiber/starter/backend/entities"
	"go-fiber/starter/backend/entities/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MySqlConnect(env dto.MySQLEnv) *gorm.DB {
	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		env.Username, env.Password, env.Host, env.Port, env.DB)

	connect, errConnect := gorm.Open(mysql.Open(address), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if errConnect != nil {
		fmt.Println("Failed to connect to the database")
	} else {
		fmt.Println("Connect Mysql")
	}

	if errMigrate := connect.AutoMigrate(entities.People{}); errMigrate != nil {
		fmt.Printf("Failed to auto migrate: %v", errMigrate)
	}

	return connect
}
