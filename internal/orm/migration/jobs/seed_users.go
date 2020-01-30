package jobs

import (
	"github.com/aliereno/softwareproject-api/internal/logger"
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

// SeedUsers inserts the first users
var SeedUsers *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_USERS",
	Migrate: func(db *gorm.DB) error {
		name := "admin"
		email := "admin@system.com"
		role := 1
		dbo := &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		tx := db.New().Begin()
		tx = tx.Create(dbo).First(dbo)

		tx = tx.Commit()

		name = "support member"
		email = "support@system.com"
		role = 2
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txx := db.New().Begin()
		txx = txx.Create(dbo).First(dbo)

		txx = txx.Commit()

		name = "customer 1"
		email = "customer@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx := db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Onur"
		email = "onur@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Mete"
		email = "mete@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Emre"
		email = "emre@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Okan"
		email = "okan@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Ali"
		email = "ali@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Oğuz"
		email = "oguz@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()

		name = "Mehmet"
		email = "mehmet@system.com"
		role = 0
		dbo = &models.User{
			Name:  &name,
			Email: email,
			Role:  &role,
		}
		if err := dbo.SetPassword("123"); err != nil {
			logger.Error("password 123")
		}
		// Create scoped clean db interface
		txxx = db.New().Begin()
		txxx = txxx.Create(dbo).First(dbo)

		txxx = txxx.Commit()
		return nil
	},
	Rollback: func(db *gorm.DB) error {
		return nil
	},
}
