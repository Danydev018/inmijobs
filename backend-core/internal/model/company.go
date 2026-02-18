package model

type Rol string

const (
	RolAdmin      Rol = "Admin"
	RolEmployment Rol = "Employment"
	RolUser       Rol = "User"
)

type Company struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Location string `gorm:"not null"`
	OwnerID  string `gorm:"index"`
	Owner    User   `gorm:"foreignKey:OwnerID"`
}

type Employee struct {
	ID        int     `gorm:"primaryKey"`
	UserID    string  `gorm:"not null" `
	User      User    `gorm:"foreignKey:UserID"`
	CompanyID int     `gorm:"not null"`
	Company   Company `gorm:"foreignKey:CompanyID"`
	Rol       Rol     `gorm:"index;type:rol;not null"`
}
