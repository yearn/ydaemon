package storage

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

/**************************************************************************************************
** GetDB returns the database connection
**
** This function provides access to the database connection for the application.
** If the connection hasn't been established yet, it will create a new one.
**
** @return *gorm.DB The database connection
** @return error Any error encountered during connection
**************************************************************************************************/
func GetDB() *gorm.DB {
	if db == nil {
		dsn := os.Getenv("KONG_POSTGRES_DSN")

		// If DSN is empty, database features are disabled
		if dsn == "" {
			fmt.Println("KONG_POSTGRES_DSN not set - database features disabled")
			return nil
		}

		// Configure GORM logger
		newLogger := logger.New(
			&gormLogger{}, // Custom logger implementation
			logger.Config{
				SlowThreshold:             time.Second, // Log queries that take more than 1 second
				LogLevel:                  logger.Warn, // Only log warnings and errors
				IgnoreRecordNotFoundError: true,        // Don't log "record not found" errors
				Colorful:                  false,       // Disable color output
			},
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			fmt.Printf("Failed to connect to database: %v\n", err)
			return nil
		}
	}
	return db
}

// gormLogger implements the logger.Interface
type gormLogger struct{}

func (l *gormLogger) Printf(format string, v ...interface{}) {
	// Only log if the query took more than 1 second
	if len(v) > 0 {
		if duration, ok := v[0].(time.Duration); ok && duration > time.Second {
			fmt.Printf(format, v...)
		}
	}
}
