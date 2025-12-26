package models

type City struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}

type Country struct {
	BaseModel
	Name   string `gorm:"size:15;type:string;not null"`
	Cities *[]City
}
