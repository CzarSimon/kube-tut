package main

import "github.com/CzarSimon/util"

const (
	SERVER_PORT_KEY     = "SERVER_PORT"
	DEFAULT_SERVER_PORT = "3000"
	DB_HOST_KEY         = "DB_HOST"
	DB_PASSWORD_KEY     = "DB_PASSWORD"
	DB_USER_KEY         = "DB_USER"
	DB_NAME_KEY         = "DB_NAME"
	DB_PORT_KEY         = "DB_PORT"
	DEFAULT_DB_PORT     = "5432"
)

//Config is the main configuration type
type Config struct {
	server util.ServerConfig
	db     util.PGConfig
}

// getConfig Sets up inital config
func getConfig() Config {
	return Config{
		server: getServerConfig(),
		db:     getDBConfig(),
	}
}

// getServerConfig Sets up server config
func getServerConfig() util.ServerConfig {
	return util.ServerConfig{
		Port: util.GetEnvVar(SERVER_PORT_KEY, DEFAULT_SERVER_PORT),
	}
}

// getDBConfig Sets up database configuration
func getDBConfig() util.PGConfig {
	return util.PGConfig{
		Host:     util.GetEnvVar(DB_HOST_KEY, "localhost"),
		Password: util.GetEnvVar(DB_PASSWORD_KEY, "pwd"),
		User:     util.GetEnvVar(DB_USER_KEY, "kube"),
		DB:       util.GetEnvVar(DB_NAME_KEY, "kube-tut"),
		Port:     util.GetEnvVar(DB_PORT_KEY, DEFAULT_DB_PORT),
	}
}
