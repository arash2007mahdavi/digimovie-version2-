package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id string `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy  string         `gorm:"not null;default:null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

type Movie struct {
	BaseModel
	Name string `gorm:"type:string;size:50;not null"`
	Director Director
	DirectorId string
}
