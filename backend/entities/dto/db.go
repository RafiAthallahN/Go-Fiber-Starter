package dto

type (

	/* This struct is used for parse several information from .env to MySQLConnect function*/
	MySQLEnv struct {
		Host     string
		Port     string
		DB       string
		Username string
		Password string
	}

	/* This struct is used for parse several information from .env to PostgreSQLConnect function*/
	PostgreSQLEnv struct {
		Host     string
		Port     string
		DB       string
		Username string
		Password string
	}

	/* This struct is used for parse several information from .env to PostgreSQLConnect function*/
	MongoDBEnv struct {
		Host       string
		Port       string
		DB         string
		Collection string
		Username   string
		Password   string
	}

	/* You could add or remove it based on your needs*/
)
