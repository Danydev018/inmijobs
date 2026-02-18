package model

type Job struct {
	ID          int     `gorm:"primaryKey"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Status      string  `gorm:"not null"`
	CompanyID   int     `gorm:"not null"`
	Company     Company `gorm:"foreignKey:CompanyID"`
	RecruiterID string  `gorm:"not null" json:"recruiter_id"`
	Recruiter   User    `gorm:"foreignKey:RecruiterID"`

	Posts []Post `gorm:"foreignKey:JobID"`
}
