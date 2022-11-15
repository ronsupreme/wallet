package controller

import (
	"go_api/service"

	"github.com/gin-gonic/gin"
)

type LinkController interface {
	UpdatePartner(context *gin.Context)
	ProfilePartner(context *gin.Context)
}
type linkController struct {
	partnerService service.PartnerService
	jwtService     service.JWTService
}

func NewLinkController(partnerService service.PartnerService, jwtService service.JWTService) PartnerController {
	return &partnerController{
		partnerService: partnerService,
		jwtService:     jwtService,
	}
}
