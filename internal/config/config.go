package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string

	ServerPort string
	GinMode    string
	JwtSecret  string
	LarekUrl   string

	MlService string
}

var Cfg Config

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	Cfg = Config{
		DbUser:     os.Getenv("POSTGRES_USER"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		DbHost:     os.Getenv("POSTGRES_HOST"),
		DbName:     os.Getenv("POSTGRES_DB"),
		ServerPort: os.Getenv("SERVER_PORT"),
		GinMode:    os.Getenv("GIN_MODE"),
		JwtSecret:  os.Getenv("JWT_SECRET"),
		MlService:  os.Getenv("ML_SERVICE"),
		LarekUrl:   os.Getenv("LAREK_URL"),
	}
	return nil
}
