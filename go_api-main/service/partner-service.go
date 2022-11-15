package service

import (
	"go_api/dto"
	"go_api/entity"
	"go_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type PartnerService interface {
	Updatepartner(partner dto.PartnerUpdateDTO) entity.Partner
	ProfilePartner(userID string) entity.Partner
}

type partnerService struct {
	partnerRepository repository.PartnerRepository
}

func NewPartnerService(partnerRepo repository.PartnerRepository) PartnerService {
	return &partnerService{
		partnerRepository: partnerRepo,
	}
}
func (service *partnerService) Updatepartner(user dto.PartnerUpdateDTO) entity.Partner {
	partnerToUpdate := entity.Partner{}
	err := smapping.FillStruct(&partnerToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedPartner := service.partnerRepository.UpdatePartner(partnerToUpdate)
	return updatedPartner
}
func (service *partnerService) ProfilePartner(userID string) entity.Partner {
	return service.partnerRepository.ProfileUserPartner(userID)
}
