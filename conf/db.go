package conf

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDatabaseConnectionURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", configuration.Database.DbUser, configuration.Database.DbPassword,
		configuration.Database.DbHost, configuration.Database.DbPort, configuration.Database.DbName)
}

func DBConn() *gorm.DB {
	DbUrl := GetDatabaseConnectionURL()
	var db *gorm.DB
	var err error
	db, err = gorm.Open(mysql.Open(DbUrl), GetGormConfig())
	if err != nil {
		panic(err)
	}

	// Using Master/Slave here
	// db.Use(dbresolver.Register(dbresolver.Config{
	// Replicas: []gorm.Dialector{mysql.Open("slave_database")},
	// Policy: dbresolver.RandomPolicy{},
	// }))

	return db
}

func GetGormConfig() *gorm.Config {
	logMode := logger.Silent
	return &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	}
}
