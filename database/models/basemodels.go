package models

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy  int            `gorm:"null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

type Movie struct {
	BaseModel
	Name       string   `gorm:"type:string;size:50;not null"`
	Director   Director `gorm:"foreignKey:DirectorId"`
	DirectorId int
}

type Director struct {
	BaseModel
	Firstname    string `gorm:"type:string;size:20;null"`
	Lastname     string `gorm:"type:string;size:30;null"`
	Username     string `gorm:"type:string;size:50;not null;unique"`
	Password     string `gorm:"type:string;not null"`
	MobileNumber string `gorm:"type:string;size:11;not null;unique"`
	Email        string `gorm:"type:string;size:80;null"`
	Enabled      bool   `gorm:"type:bool;default:true"`
	Movies       []Movie
}

type User struct {
	BaseModel
	Firstname    string `gorm:"type:string;size:20;null"`
	Lastname     string `gorm:"type:string;size:30;null"`
	Username     string `gorm:"type:string;size:50;not null;unique"`
	Password     string `gorm:"type:string;not null"`
	MobileNumber string `gorm:"type:string;size:11;not null;unique"`
	Email        string `gorm:"type:string;size:80;null"`
	Enabled      bool   `gorm:"type:bool;default:true"`
}
