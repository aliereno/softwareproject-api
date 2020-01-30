// Package orm provides `GORM` helpers for the creation, migration and access
// on the project's database
package orm

import (
	"fmt"
	log "github.com/aliereno/softwareproject-api/internal/logger"
	"github.com/aliereno/softwareproject-api/internal/orm/migration"

	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

var autoMigrate, logMode, seedDB bool
var hostDB, userDB, portDB, passwordDB, nameDB, sslDB, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	dialect = "postgres"
	hostDB = "localhost"
	portDB = "5432"
	userDB = "postgres"
	passwordDB = "iü"
	nameDB = "softwareproject"
	sslDB = "disable"
	seedDB = true
	logMode = true
	autoMigrate = true
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	dsn := fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=%s password=%s", hostDB, userDB, portDB, nameDB, sslDB, passwordDB) //Build connection string

	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}
	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(logMode)
	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM] Database connection initialized.")
	return orm, err
}
