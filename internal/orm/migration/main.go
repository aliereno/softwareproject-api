package migration

import (
	"fmt"
	log "github.com/aliereno/softwareproject-api/internal/logger"
	"github.com/aliereno/softwareproject-api/internal/orm/migration/jobs"
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Comment{},
		&models.Ringtone{},
		&models.Ticket{},
		&models.Activity{},
	).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
	// Keep a list of migrations here
	m := gormigrate.New(db, gormigrate.DefaultOptions, nil)
	m.InitSchema(func(db *gorm.DB) error {
		log.Info("[Migration.InitSchema] Initializing database schema")
		switch db.Dialect().GetName() {
		//case "postgres":
		//db.Exec("create extension \"uuid-ossp\";")
		}
		if err := updateMigration(db); err != nil {
			return fmt.Errorf("[Migration.InitSchema]: %v", err)
		}
		// Add more jobs, etc here
		return nil
	})
	m.Migrate()

	if err := updateMigration(db); err != nil {
		return err
	}
	// DROP COLUMN
	//err := db.Model(&models.User{}).DropColumn("job_title").Error
	//if err != nil {
	//	// Do whatever you want to do!
	//	log.Info("ERROR: We expect the job_title column to be drop-able")
	//}
	m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		jobs.SeedUsers,
		jobs.SeedRingtone,
	})
	return m.Migrate()
}
