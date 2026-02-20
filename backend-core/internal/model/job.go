package model

/* Esto es una modelo PRUEBA para ver si se puede relacionar con atributos post
   no nos hacemos responsable del modelo original ni este
   */
type Status string

const (
	On  Status = "Active"
	Off Status = "Inactive"
)

type Job struct {
	ID          int     `gorm:"primaryKey"`
	Title       string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Status      string  `gorm:"index;type:status;not null"`
	CompanyID   int     //`gorm:"not null"`
	Company     Company `gorm:"foreignKey:CompanyID"`
	RecruiterID string  //`gorm:"not null" json:"recruiter_id"`
	Recruiter   User    `gorm:"foreignKey:RecruiterID"`

	Posts []Post `gorm:"foreignKey:JobID"`
}
