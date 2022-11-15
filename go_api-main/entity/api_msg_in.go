package entity

import "time"

type Api_Msg_In struct {
	Idrequest   uint64    `gorm:"primary_key:auto_increment" json:"idrequest"`
	Source      string    `gorm:"type:varchar(255)" json:"source"`
	TypeMsg     string    `gorm:"type:varchar(63)" json:"typemsg"`
	ContentMsg  string    `gorm:"type:varchar(255)" json:"contentmsg"`
	PartnerName string    `gorm:"type:varchar(63)" json:"partner_name"`
	CreatedDate time.Time `gorm:"autoUpdateTime;type:datetime; default:current_timestamp" json:"create_date"`
}
