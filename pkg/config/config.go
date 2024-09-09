package config

import (
    "log"
    "os"
)

type Config struct {
    ServerAddress    string
    PostgresConn     string
    PostgresJdbcUrl  string
    PostgresUsername string
    PostgresPassword string
    PostgresHost     string
    PostgresPort     string
    PostgresDatabase string
}

func GetEnvVariable(key string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Fatalf("Environment variable %s not set", key)
    }
    return value
}

func LoadConfig() *Config {
    return &Config{
        ServerAddress:    GetEnvVariable("SERVER_ADDRESS"),
        PostgresConn:     GetEnvVariable("POSTGRES_CONN"),
        PostgresJdbcUrl:  GetEnvVariable("POSTGRES_JDBC_URL"),
        PostgresUsername: GetEnvVariable("POSTGRES_USERNAME"),
        PostgresPassword: GetEnvVariable("POSTGRES_PASSWORD"),
        PostgresHost:     GetEnvVariable("POSTGRES_HOST"),
        PostgresPort:     GetEnvVariable("POSTGRES_PORT"),
        PostgresDatabase: GetEnvVariable("POSTGRES_DATABASE"),
    }
}