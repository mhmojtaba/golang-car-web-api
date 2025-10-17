package models

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:16;not null;unique"`
	UserRoles *[]UserRole
}
