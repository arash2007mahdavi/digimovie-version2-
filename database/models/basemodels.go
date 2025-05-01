package models

import (
	"context"
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

func (base *BaseModel) BeforeCreate(ctx context.Context) error {
	base.CreatedAt = time.Now().UTC()
	base.CreatedBy = int(ctx.Value("Userid").(float64))
	return nil
}

type Movie struct {
	BaseModel
	Name       string   `gorm:"type:string;size:50;not null"`
	Director   Director `gorm:"foreignKey:DirectorId"`
	DirectorId string
}

type Director struct {
	BaseModel
	Firstname    string `gorm:"type:string;size:20;null"`
	Lastname     string `gorm:"type:string;size:30;null"`
	Username     string `gorm:"type:string;size:50;not null;unique"`
	Password     string `gorm:"type:string;size:50;not null"`
	MobileNumber string `gorm:"type:string;size:11;not null;unique"`
	Email        string `gorm:"type:string;size:80;null"`
	Enabled      bool   `gorm:"type:bool;default:true"`
}

type User struct {
	BaseModel
	Firstname    string `gorm:"type:string;size:20;null"`
	Lastname     string `gorm:"type:string;size:30;null"`
	Username     string `gorm:"type:string;size:50;not null;unique"`
	Password     string `gorm:"type:string;size:50;not null"`
	MobileNumber string `gorm:"type:string;size:11;not null;unique"`
	Email        string `gorm:"type:string;size:80;null"`
	Enabled      bool   `gorm:"type:bool;default:true"`
}
