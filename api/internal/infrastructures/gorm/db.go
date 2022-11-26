package gorm

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDatabase(host string, database string, username string, password string) *gorm.DB {
    Logger := logger.Default
    if os.Getenv("ENV") == "development" {
        Logger = Logger.LogMode(logger.Info)
    }
    db, err := gorm.Open(postgres.Open("host=" + host + " user=" + username + " password=" + password + " dbname=" + database), &gorm.Config{
        Logger: Logger,
    })
    if err != nil {
        panic("failed to connect database")
    }
    DB, err := db.DB()
    if err != nil {
        panic("failed to connect database")
    }
    driver, err := migratePostgres.WithInstance(DB, &migratePostgres.Config{})
    m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }
    m.Up()
    return db
}
