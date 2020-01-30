package jobs

import (
	"github.com/aliereno/softwareproject-api/internal/orm/models"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

var SeedRingtone *gormigrate.Migration = &gormigrate.Migration{
	ID: "SEED_RINGTONES",
	Migrate: func(db *gorm.DB) error {
		title := "Animal"
		animal := &models.Category{
			Title: &title,
		}
		tx := db.New().Begin()
		tx = tx.Create(animal).First(animal)
		tx = tx.Commit()

		title = "Nature"
		nature := &models.Category{
			Title: &title,
		}
		tx = db.New().Begin()
		tx = tx.Create(nature).First(nature)
		tx = tx.Commit()

		lion := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "aslan_sesi.mp3",
			Title:      "Lion",
			Price:      15,
			Hit:        234,
		}
		tx = db.New().Begin()
		tx = tx.Create(lion).First(lion)
		tx = tx.Commit()

		horse := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "at_sesi.mp3",
			Title:      "Horse",
			Price:      15,
			Hit:        123,
		}
		tx = db.New().Begin()
		tx = tx.Create(horse).First(horse)
		tx = tx.Commit()

		elephant := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "fil_sesi.mp3",
			Title:      "Elephant",
			Price:      15,
			Hit:        66,
		}
		tx = db.New().Begin()
		tx = tx.Create(elephant).First(elephant)
		tx = tx.Commit()

		cat := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "kedi_sesi.mp3",
			Title:      "Cat",
			Price:      15,
			Hit:        34,
		}
		tx = db.New().Begin()
		tx = tx.Create(cat).First(cat)
		tx = tx.Commit()

		dog := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "kopek_sesi.mp3",
			Title:      "Dog",
			Price:      15,
			Hit:        345,
		}
		tx = db.New().Begin()
		tx = tx.Create(dog).First(dog)
		tx = tx.Commit()

		frog := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "kurbaga_sesi.mp3",
			Title:      "Frog",
			Price:      15,
			Hit:        456,
		}
		tx = db.New().Begin()
		tx = tx.Create(frog).First(frog)
		tx = tx.Commit()

		wolf := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "kurt_sesi.mp3",
			Title:      "Wolf",
			Price:      15,
			Hit:        567,
		}
		tx = db.New().Begin()
		tx = tx.Create(wolf).First(wolf)
		tx = tx.Commit()

		bird := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "kus_sesi.mp3",
			Title:      "Bird",
			Price:      15,
			Hit:        678,
		}
		tx = db.New().Begin()
		tx = tx.Create(bird).First(bird)
		tx = tx.Commit()

		monkey := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "maymun_sesi.mp3",
			Title:      "Monkey",
			Price:      15,
			Hit:        789,
		}
		tx = db.New().Begin()
		tx = tx.Create(monkey).First(monkey)
		tx = tx.Commit()

		snake := &models.Ringtone{
			CategoryID: &animal.ID,
			Category:   animal,
			FileName:   "yilan_sesi.mp3",
			Title:      "Snake",
			Price:      15,
			Hit:        890,
		}
		tx = db.New().Begin()
		tx = tx.Create(snake).First(snake)
		tx = tx.Commit()

		thunder := &models.Ringtone{
			CategoryID: &nature.ID,
			Category:   nature,
			FileName:   "gokgurultusu_sesi.mp3",
			Title:      "Thunder",
			Price:      10,
			Hit:        901,
		}
		tx = db.New().Begin()
		tx = tx.Create(thunder).First(thunder)
		tx = tx.Commit()

		snow := &models.Ringtone{
			CategoryID: &nature.ID,
			Category:   nature,
			FileName:   "kar_sesi.mp3",
			Title:      "Snowing",
			Price:      10,
			Hit:        200,
		}
		tx = db.New().Begin()
		tx = tx.Create(snow).First(snow)
		tx = tx.Commit()

		river := &models.Ringtone{
			CategoryID: &nature.ID,
			Category:   nature,
			FileName:   "nehir_sesi.mp3",
			Title:      "River",
			Price:      10,
			Hit:        300,
		}
		tx = db.New().Begin()
		tx = tx.Create(river).First(river)
		tx = tx.Commit()

		wind := &models.Ringtone{
			CategoryID: &nature.ID,
			Category:   nature,
			FileName:   "ruzgar_sesi.mp3",
			Title:      "Wind",
			Price:      10,
			Hit:        10,
		}
		tx = db.New().Begin()
		tx = tx.Create(wind).First(wind)
		tx = tx.Commit()

		raining := &models.Ringtone{
			CategoryID: &nature.ID,
			Category:   nature,
			FileName:   "yagmur_sesi.mp3",
			Title:      "Raining",
			Price:      10,
			Hit:        5,
		}
		tx = db.New().Begin()
		tx = tx.Create(raining).First(raining)
		tx = tx.Commit()

		return nil
	},
	Rollback: func(db *gorm.DB) error {
		return nil
	},
}
