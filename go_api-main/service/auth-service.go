package service

import (
	"log"

	"go_api/dto"
	"go_api/entity"
	"go_api/logger"
	"go_api/repository"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

// interface use case need to do for authen
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(partner dto.RegisterDTO) entity.Partner
	FindByEmail(email string) entity.Partner
	IsDuplicateEmail(email string) bool
}

type authService struct {
	partnerRepository repository.PartnerRepository
}

func NewAuthService(partner repository.PartnerRepository) AuthService {
	return &authService{
		partnerRepository: partner,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.partnerRepository.VerifyCredentialPartner(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(partner dto.RegisterDTO) entity.Partner {
	userToCreate := entity.Partner{}
	logger.InfoLogger.Println("CreateUser", "userToCreate: ", userToCreate)
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&partner))
	if err != nil {
		log.Println(err)
	}
	res := service.partnerRepository.InsertPartner(userToCreate)
	logger.InfoLogger.Println("CreateUser", "res: ", res)
	return res
}

func (service *authService) FindByEmail(email string) entity.Partner {
	return service.partnerRepository.FindByEmailPartner(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.partnerRepository.IsDuplicateEmailPartner(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, planPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, planPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
