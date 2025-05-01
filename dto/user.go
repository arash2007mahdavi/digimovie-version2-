package dto

type UserCreate struct {
	Username     string `json:"username" binding:"min=4,max=75,required"`
	Password     string `json:"password" binding:"password,required"`
	MobileNumber string `json:"mobileNumber" binding:"mobileNumber,required"`
	CreatedBy    int    `json:"createdBy"`
}

type UserUpdate struct {
	Firstname    string `json:"firstname,omitempty" binding:"omitempty,alpha,min=2,max=25"`
	Lastname     string `json:"lastname,omitempty" binding:"omitempty,alpha,min=2,max=35"`
	Username     string `json:"username,omitempty" binding:"omitempty,min=4,max=75"`
	Password     string `json:"password,omitempty" binding:"omitempty,password"`
	MobileNumber string `json:"mobileNumber,omitempty" binding:"omitempty,mobileNumber"`
	Email        string `josn:"email,omitempty" binding:"omitempty,email"`
	Enabled      bool   `json:"enabled" binding:"omitempty"`
}

type UserRes struct {
	Id           int
	Firstname    string `json:",omitempty"`
	Lastname     string `json:",omitempty"`
	Username     string
	Password     string
	MobileNumber string
	Email        string `json:",omitempty"`
	Enabled      bool
}
