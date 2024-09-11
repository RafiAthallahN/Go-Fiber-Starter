package config

import (
	"fmt"
	"go-fiber/starter/backend/entities"
	"go-fiber/starter/backend/entities/dto"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func MySqlConnect(env dto.MySQLEnv) *gorm.DB {
	address := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		env.Username, env.Password, env.Host, env.Port, env.DB)

	connect, errConnect := gorm.Open(mysql.Open(address), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if errConnect != nil {
		fmt.Println("Failed to connect to the database")
	} else {
		fmt.Println("Connect Mysql")
	}

	// Remove it if you don't want to use migration
	if errMigrate := connect.AutoMigrate(entities.People{}); errMigrate != nil {
		fmt.Printf("Failed to auto migrate: %v", errMigrate)
	}

	return connect
}
