package apps

import (
	"go-fiber/starter/backend/entities/dto"
	"go-fiber/starter/backend/routes"
	"go-fiber/starter/config"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func StartApps() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type, Authorization",
	}))

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006/01/02 15:04:05",
		DisableTimestamp: false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		},
	})

	log := logrus.New()
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql := config.MySqlConnect(dto.MySQLEnv{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		DB:       os.Getenv("MYSQL_DB"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
	})

	peopleRoutes := routes.PeopleRoutes{
		App:    app,
		People: SetupPeople(mysql, log),
	}
	peopleRoutes.SetupPeopleRoutes()

	errApp := app.Listen(":" + os.Getenv("APP_PORT"))
	if errApp != nil {
		logrus.Fatalf("Error starting Fiber app: %v", errApp)
	}
}
