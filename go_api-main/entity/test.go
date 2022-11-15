package entity

type Test struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	UserName    string `gorm:"type:varchar(63)" json:"userName"`
	CompanyName string `gorm:"type:varchar(255)" json:"companyName"`
	CountryCode string `gorm:"type:varchar(63)" json:"countryCode"`
	CountryName string `gorm:"type:varchar(255)" json:"countryName"`
	Email       string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
}
