package repository

import (
	"gorm.io/gorm"
)

// What can to do db
type PingRepository interface {
	Ping() string
}

type pingConnection struct {
	connection *gorm.DB
}

// Create new instance of UserRepository
func NewPingRepository(db *gorm.DB) PingRepository {
	return &pingConnection{
		connection: db,
	}
}

func (db *pingConnection) Ping() string {
	return "nghilc"
}
