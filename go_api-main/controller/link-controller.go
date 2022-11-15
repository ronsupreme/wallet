package controller

import (
	"fmt"
	"go_api/dto"
	"go_api/helper"
	"go_api/logger"
	"go_api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type LinkController interface {
	Insert(context *gin.Context)
}
type linkController struct {
	linkService service.LinkService
	jwtService  service.JWTService
}

func NewLinkController(linkService service.LinkService, jwtService service.JWTService) LinkController {
	return &linkController{
		linkService: linkService,
		jwtService:  jwtService,
	}
}
func (c *partnerController) Insert(context *gin.Context) {
	var partnerUpdateDTO dto.PartnerUpdateDTO
	errDTO := context.ShouldBind(&partnerUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	authHeader := context.GetHeader("Authorization")
	logger.InfoLogger.Println("authHeader", authHeader)
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errDTO.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	partnerUpdateDTO.ID = id
	u := c.partnerService.Updatepartner(partnerUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}
