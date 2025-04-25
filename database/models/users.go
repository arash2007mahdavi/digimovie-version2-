package models

type Director struct {
	BaseModel
	Firstname    string `gorm:"type:string;size:20;null"`
	Lastname     string `gorm:"type:string;size:30;null"`
	Username     string `gorm:"type:string;size:50;not null;unique"`
	MobileNumber string `gorm:"type:string;size:11;null"`
	Email        string `gorm:"type:string;size:80;null"`
	Enabled      bool   `gorm:"type:bool;default:true"`
}
