package models

import "time"

type City struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}

type Country struct {
	BaseModel
	Name    string `gorm:"size:15;type:string;not null"`
	Cities  []City
	Company []Company
}

type CountryYear struct {
	BaseModel
	Title         string    `gorm:"size:15;type:string;not null; unique"`
	Year          int       `gorm:"type:int;uniqueindex;not null"`
	StartAt       time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	EndAt         time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	CarModelYears []CarModelYear
}

type Color struct {
	BaseModel
	Name           string `gorm:"size:15;type:string;not null"`
	HexCode        string `gorm:"size:7;type:string;not null"`
	CarModelColors []CarModelColor
}

type File struct {
	BaseModel
	Name        string `gorm:"size:100;type:string;not null"`
	Directory   string `gorm:"size:255;type:string;not null"`
	Description string `gorm:"size:255;type:string;not null"`
	MediaType   string `gorm:"size:50;type:string;not null"`
}
