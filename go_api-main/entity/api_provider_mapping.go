package entity

import "time"

type api_provider_mapping struct {
	ID          uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Name        string    `gorm:"type:varchar(255)" json:"name"`
	UserName    string    `gorm:"type:varchar(63)" json:"userName"`
	CompanyName string    `gorm:"type:varchar(255)" json:"companyName"`
	CountryCode string    `gorm:"type:varchar(63)" json:"countryCode"`
	CountryName string    `gorm:"type:varchar(255)" json:"countryName"`
	Email       string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password    string    `gorm:"->;<-;not null" json:"-"`
	Token       string    `gorm:"-" json:"token"`
	Status      int       `gorm:"type:int" json:"status"`
	CreatedAt   time.Time `gorm:"autoUpdateTime;type:datetime; default:current_timestamp" json:"createdAt"`
	CreatedBy   string    `gorm:"type:varchar(63)" json:"createdBy"`
	ModifiedAt  time.Time `gorm:"autoUpdateTime;type:datetime; default:current_timestamp on update current_timestamp" json:"modifiedAt"`
	ModifiedBy  string    `gorm:"type:varchar(63)" json:"modifiedBy"`
}