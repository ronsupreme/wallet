package dto

import "time"

type PartnerUpdateDTO struct {
	ID          uint64    `json:"id" form:"id"`
	Name        string    `json:"name" form:"name" bindding:"required"`
	UserName    string    `json:"username" form:"name" bindding:"required"`
	CompanyName string    `json:"company" form:"name" bindding:"required"`
	CountryCode string    `json:"country" form:"code" bindding:"iso3166_1_alpha2"`
	CountryName string    `json:"conntry" form:"name" bindding:"required"`
	Email       string    `json:"email" form:"email" bindding:"required,email"`
	Password    string    `json:"password,omitempty" form:"password,omitempty"`
	Token       string    `json:"token" form:"name" bindding:"required"`
	Status      int       `json:"status" form:"name" bindding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	CreatedBy   string    `json:"createdBy" form:"name"`
	ModifiedAt  time.Time `json:"modifiedAt" form:"name"`
	ModifiedBy  string    `json:"modifiedBy" form:"name"`
}
