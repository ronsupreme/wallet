package repository

import (
	"go_api/entity"
	"go_api/logger"

	"gorm.io/gorm"
)

type PartnerRepository interface {
	InsertPartner(partner entity.Partner) entity.Partner
	UpdatePartner(partner entity.Partner) entity.Partner
	VerifyCredentialPartner(email string, password string) interface{}
	IsDuplicateEmailPartner(email string) (tx *gorm.DB)
	FindByEmailPartner(email string) entity.Partner
	ProfileUserPartner(userID string) entity.Partner
}
type result struct {
	service_name string
}
type partnerConnection struct {
	connection *gorm.DB
}

// Create new instance of UserRepository
func NewPartnerRepository(db *gorm.DB) PartnerRepository {
	return &partnerConnection{
		connection: db,
	}
}

func (db *partnerConnection) InsertPartner(partner entity.Partner) entity.Partner {
	partner.Password = hashAndSalt([]byte(partner.Password))
	db.connection.Save(&partner)
	return partner
}
func (db *partnerConnection) UpdatePartner(partner entity.Partner) entity.Partner {
	if partner.Password != "" {
		partner.Password = hashAndSalt([]byte(partner.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, partner.ID)
		partner.Password = tempUser.Password
	}

	db.connection.Save(&partner)
	return partner
}
func (db *partnerConnection) VerifyCredentialPartner(email string, password string) interface{} {
	var partner entity.Partner
	logger.InfoLogger.Println("email: ", email)
	res := db.connection.Where("email = ?", email).Take(&partner)
	logger.InfoLogger.Println("res: ")
	if res.Error == nil {
		logger.InfoLogger.Println("partner: ", partner)
		return partner
	}
	return nil
}

func (db *partnerConnection) IsDuplicateEmailPartner(email string) (tx *gorm.DB) {
	var partner entity.Partner
	return db.connection.Where("email = ?", email).Take(&partner)
}

func (db *partnerConnection) FindByEmailPartner(email string) entity.Partner {
	var partner entity.Partner
	db.connection.Where("email = ?", email).Take(&partner)
	return partner
}

func (db *partnerConnection) ProfileUserPartner(userID string) entity.Partner {
	var partner entity.Partner
	db.connection.Find(&partner, userID)
	return partner
}

