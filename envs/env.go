package envs

import (
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	jWTToken        string
	jWTRefreshToken string
	dbHost          string
	dbPassword      string
	dbUser          string
	dbPort          string
	dbName          string
}

func newEnv() *env {
	godotenv.Load()
	return &env{
		jWTToken:        os.Getenv("JWT_SECRET"),
		jWTRefreshToken: os.Getenv("JWT_REFRESH_SECRET"),
		dbHost:          os.Getenv("DB_HOST"),
		dbPassword:      os.Getenv("DB_PASSWORD"),
		dbUser:          os.Getenv("DB_USERNAME"),
		dbPort:          os.Getenv("DB_PORT"),
		dbName:          os.Getenv("DB_DATABASE"),
	}
}

func (env *env) GetJWTToken() string {
	return env.jWTToken
}

func (env *env) GetJWTRefreshToken() string {
	return env.jWTRefreshToken
}

func (env *env) GetDbHost() string {
	return env.dbHost
}

func (env *env) GetDbPassword() string {
	return env.dbPassword
}

func (env *env) GetDbUser() string {
	return env.dbUser
}

func (env *env) GetDbPort() string {
	return env.dbPort
}

func (env *env) GetDbName() string {
	return env.dbName
}

var instance *env

func GetInstance() *env {
	if instance == nil {
		instance = newEnv()
	}
	return instance
}
