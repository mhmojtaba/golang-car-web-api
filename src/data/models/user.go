package models

type User struct {
	BaseModel
	Username  string `gorm:"type:string;size:20;not null;unique"`
	FirstName string `gorm:"type:string;size:15;null"`
	LastName  string `gorm:"type:string;size:25;null"`
	Mobile    string `gorm:"type:string;size:11;null;unique;default:null"`
	Email     string `gorm:"type:string;size:64;null;unique;default:null"`
	Password  string `gorm:"type:string;size:64;not null"`
	Active    bool   `gorm:"default:true"`
	UserRoles *[]UserRole
}

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;size:16;not null;unique"`
	UserRoles *[]UserRole
}

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId int
	RoleId int
}
