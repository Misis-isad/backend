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
	}
	return nil

	// if err := env.Parse(&Cfg); err != nil {
	// 	return err
	// }
	// return nil
}
