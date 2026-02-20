package database

import (
	"log"
	"log/slog"

	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/model"
	"github.com/Gabo-div/bingo/inmijobs/backend-core/internal/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Seed(db *gorm.DB) {
	// Create Test User
	userID := "user_test_123"
	userID_2 := "user_test_456"
	ownerID := "owner_99"

	user := model.User{
		ID:            userID,
		Name:          "Test User",
		Email:         "test@example.com",
		EmailVerified: true,
	}

	user_2 := model.User{
		ID:            userID_2,
		Name:          "Test User 2",
		Email:         "test2@example.com",
		EmailVerified: true,
	}

	owner := model.User{
        ID:    "owner_99",
        Name:  "Admin Owner",
        Email: "admin@empresa.com",
    }
    
    recruiter := model.User{
        ID:    "user_test_123",
        Name:  "Test Recruiter",
        Email: "recruiter@test.com",
    }

    db.Clauses(clause.OnConflict{DoNothing: true}).Create(&owner)
    db.Clauses(clause.OnConflict{DoNothing: true}).Create(&recruiter)

	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&user)

	result_2 := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&user_2)

	if result.Error != nil {
		slog.Error("[Database] Failed to seed user", "error", result.Error)
	} else if result.RowsAffected > 0 {
		slog.Info("[Database] Seeded test user")
	}

	if result_2.Error != nil {
		slog.Error("[Database] Failed to seed user", "error", result.Error)
	} else if result.RowsAffected > 0 {
		slog.Info("[Database] Seeded test user")
	}

	// Create Test Profile
	profileID := utils.NewID()
	profile := model.Profile{
		ID:        profileID,
		UserID:    userID,
		Biography: toPtr("This is a test biography."),
		Location:  toPtr("Test City, Country"),
	}

	result = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}},
		DoNothing: true,
	}).Create(&profile)

	company := model.Company{
		ID:       1,
		Name:     "Global Tech",
		Location: "Madrid, España",
		OwnerID:  ownerID,
	}
	if err := db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoNothing: true}).Create(&company).Error; err != nil {
		slog.Error("[Database] Failed to seed company", "error", err)
	} else {
		slog.Info("[Database] Seeded test company")
	}

	job := model.Job{
		ID:          1,
		Title:       "Senior Go Developer",
		Description: "Estamos buscando un experto en Go y Microservicios.",
		Status:      "Active",
		CompanyID:   1, 
		RecruiterID: userID,     // "user_test_123"
	}

	if err := db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoNothing: true}).Create(&job).Error; err != nil {
		slog.Error("[Database] Failed to seed job", "error", err)
	} else {
		slog.Info("[Database] Seeded test job")
	}

	if result.Error != nil {
		slog.Error("[Database] Failed to seed profile", "error", result.Error)
	} else if result.RowsAffected > 0 {
		slog.Info("[Database] Seeded test profile")
	}

	reaction := model.Reaction{ID: 1, Name: "Me gusta", IconURL: "like.png"}
	db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "id"}}, DoNothing: true}).Create(&reaction)

	log.Println("INFO [Database] Seed process finished")
}

func toPtr(s string) *string {
	return &s
}
