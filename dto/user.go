package dto

type UserCreate struct {
	Username     string `json:"username" binding:"min=4,max=75,required"`
	Password     string `json:"password" binding:"password,required"`
	MobileNumber string `json:"mobileNumber" binding:"mobileNumber,required"`
	CreatedBy    int    `json:"createdBy"`
}

type UserUpdate struct {
	Firstname    string `json:"firstname" binding:"alpha,min=2,max=25"`
	Lastname     string `json:"lastname" binding:"alpha,min=2,max=35"`
	Username     string `json:"username" binding:"min=4,max=75"`
	Password     string `json:"password" binding:"password"`
	MobileNumber string `json:"mobileNumber" binding:"mobileNumber"`
	Email        string `josn:"email" binding:"email"`
	Enabled      bool   `json:"enabled"`
}

type UserRes struct {
	Id           int
	Firstname    string `json:",omitempty"`
	Lastname     string `json:",omitempty"`
	Username     string
	Password     string
	MobileNumber string
	Email        string `json:",omitempty"`
	Enabled      bool   `json:",omitempty"`
}
