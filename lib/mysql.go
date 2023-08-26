package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MYSQL_CONFIG    = "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	POSTGRES_CONFIG = "host=%s user=%s password=%s dbname=%s port=%v sslmode=%s TimeZone=Asia/Shanghai"
)

func DatabaseConnection(configFile string) *gorm.DB {
	var dialect gorm.Dialector
	var dbConnection string

	if err := godotenv.Load(configFile); err != nil {
		log.Fatalf("error while read configuration, %s", err.Error())
	}

	switch os.Getenv("db_driver") {
	case "postgres":
		dbConnection = fmt.Sprintf(POSTGRES_CONFIG, readEnvironment("db_host"), readEnvironment("db_username"), readEnvironment("db_password"), readEnvironment("db_name"), readEnvironment("db_port"), readEnvironment("db_sslmode"))
		dialect = postgres.Open(dbConnection)
	case "mysql":
		dbConnection = fmt.Sprintf(MYSQL_CONFIG, readEnvironment("db_username"), readEnvironment("db_password"), readEnvironment("db_host"), readEnvironment("db_port"), readEnvironment("db_name"))
		dialect = mysql.Open(dbConnection)
	}

	db, errConnect := gorm.Open(dialect, &gorm.Config{})
	if errConnect != nil {
		log.Fatalf("Error while connect to database :[%s]", errConnect)
		os.Exit(1)
	}

	return db

}

func readEnvironment(key string) string {
	return os.Getenv(key)
}
