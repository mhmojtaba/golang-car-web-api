package models

import (
	"gorm.io/gorm"
)

type UserRole struct {
	BaseModel
	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	RoleId int
} 