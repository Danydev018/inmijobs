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

	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&user,)

	result_2 := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoNothing: true,
	}).Create(&user_2,)

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

	if result.Error != nil {
		slog.Error("[Database] Failed to seed profile", "error", result.Error)
	} else if result.RowsAffected > 0 {
		slog.Info("[Database] Seeded test profile")
	}

	 reaction:= []model.Reaction{
		
	 }
	db.FirstOrCreate(&reaction, model.Reaction{
		ID:      1,
		Name:    "Me gusta",
		IconURL: "like.png",
	
	},)
	log.Println("INFO [Database] Seeded reactions")
}

func toPtr(s string) *string {
	return &s
}
