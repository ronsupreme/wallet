package service

import (
	"go_api/dto"
	"go_api/entity"
	"go_api/repository"
)

type LinkService interface {
	Insert(link dto.HeaderLinkDTO) entity.Api_Msg_In
}

type linkService struct {
	linkRepository repository.LinkRepository
}

func NewLinkService(linkRepo repository.LinkRepository) LinkService {
	return &linkService{
		linkRepository: linkRepo,
	}
}

