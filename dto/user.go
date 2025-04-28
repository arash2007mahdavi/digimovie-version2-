package dto

type UserDto struct {
	Firstname    string `json:"firstname" binding:"alpha,min=2,max=25"`
	Lastname     string `json:"lastname" binding:"alpha,min=2,max=35"`
	Username     string `json:"username" binding:"min=4,max=75,required"`
	MobileNumber string `json:"mobileNumber" binding:"mobileNumber"`
	Email        string `josn:"email" binding:"email"`
	Enabled      bool   `json:"enabled"`
}

type UserRes struct {
	
}