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
)

var DATABASE *gorm.DB

func initializeMySQLDatabase() (shouldUseMySQLDB bool) {
	DSN, exists := os.LookupEnv("MYSQL_DSN")
	if !exists {
		logs.Info("DSN not found, skipping database usage")
		return false
	}

	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
		Logger:                                   logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		logs.Error(err)
		return false
	}
	DATABASE = db
	db.AutoMigrate(&DBBlockTime{})
	db.AutoMigrate(&DBHistoricalPrice{})
	logs.Success(`DB initialized`)
	return true
}

func getUUID(str string) string {
	hash := sha256.Sum256([]byte(str))
	trimmedHash := hash[:16]
	finalUUID, _ := uuid.FromBytes(trimmedHash)
	return finalUUID.String()
}
