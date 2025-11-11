package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBUser     string
    DBPass     string
    DBName     string
    DBPort     string
    JWTSecret  string
    AppEnv     string
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ No se pudo cargar el archivo .env (continuando con variables del entorno)")
    }

    return &Config{
        DBHost:    os.Getenv("DB_HOST"),
        DBUser:    os.Getenv("DB_USER"),
        DBPass:    os.Getenv("DB_PASS"),
        DBName:    os.Getenv("DB_NAME"),
        DBPort:    os.Getenv("DB_PORT"),
        JWTSecret: os.Getenv("JWT_SECRET"),
        AppEnv:    os.Getenv("APP_ENV"),
    }
}
