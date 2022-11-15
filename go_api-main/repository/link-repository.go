package repository

import (
	"go_api/entity"

	"gorm.io/gorm"
)

type LinkRepository interface {
	InsertApiMsgIn(link entity.Api_Msg_In) entity.Api_Msg_In
}
type linkConnection struct {
	connection *gorm.DB
}

func NewLinkRepository(db *gorm.DB) LinkRepository {
	return &linkConnection{
		connection: db,
	}
}
func (db *partnerConnection) InsertApiMsgIn(link entity.Api_Msg_In) entity.Api_Msg_In {
	db.connection.Save(&link)
	return link
}
