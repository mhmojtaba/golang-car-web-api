package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id         int          `gorm:"primaryKey"`
	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone; not null;"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone; null;"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone; null;"`

	CreatedBy  int            `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId") // get userId from context
	var UserId = -1
	// TODO check userId type
	if value != nil {
		UserId = int(value.(float64))
	}
	b.CreatedAt = time.Now().UTC()
	b.CreatedBy = UserId

	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}
	if value != nil {
		UserId = &sql.NullInt64{Int64: int64(value.(float64)), Valid: true}
	}
	b.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	b.ModifiedBy = UserId

	return nil
}

func (b *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}
	if value != nil {
		UserId = &sql.NullInt64{Int64: int64(value.(float64)), Valid: true}
	}
	b.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	b.DeletedBy = UserId

	return nil
}
