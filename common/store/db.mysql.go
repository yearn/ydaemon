package store

import (
	"crypto/sha256"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/yearn/ydaemon/common/logs"
	"gorm.io/driver/postgres"
)

var DATABASE *gorm.DB

func initializeMySQLDatabase() (shouldUseMySQLDB bool) {
	var db *gorm.DB
	var err error
	config := gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
		Logger:                                   logger.Default.LogMode(logger.Error),
	}
	if MYSQL_DSN, existsMy := os.LookupEnv("MYSQL_DSN"); existsMy {
		db, err = gorm.Open(mysql.Open(MYSQL_DSN), &config)
		if err != nil {
			logs.Error(err)
			return false
		}
	}
	if PGSQL_DSN, existsMy := os.LookupEnv("PGSQL_DSN"); existsMy {
		db, err = gorm.Open(postgres.Open(PGSQL_DSN), &config)
		if err != nil {
			logs.Error(err)
			return false
		}
	}
	if db != nil {
		DATABASE = db
		db.AutoMigrate(&DBBlockTime{})
		db.AutoMigrate(&DBHistoricalPrice{})
		db.AutoMigrate(&DBNewVaultsFromRegistry{})
		db.AutoMigrate(&DBVault{})
		db.Table(`db_erc20`).AutoMigrate(&DBERC20{})
		logs.Success(`DB initialized`)
		return true
	}
	return false
}

func getUUID(str string) string {
	hash := sha256.Sum256([]byte(str))
	trimmedHash := hash[:16]
	finalUUID, _ := uuid.FromBytes(trimmedHash)
	return finalUUID.String()
}
