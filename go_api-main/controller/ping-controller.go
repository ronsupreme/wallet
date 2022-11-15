package controller

import (
	"go_api/logger"
	"go_api/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PingController interface {
	Ping(ctx *gin.Context)
}
type pingController struct {
	pingService service.PingService
}

func NewPingController(pingService service.PingService) PingController {
	return &pingController{
		pingService: pingService,
	}
}

func (c *pingController) Ping(ctx *gin.Context) {
	log.Println("log vo log file")

	ctx.JSON(http.StatusOK, gin.H{
		"Mess": "HI",
	})
	a := 5
	logger.InfoLogger.Println("MONO" + strconv.Itoa(a))
	logger.WarningLogger.Println("MONO" + strconv.Itoa(a+1))

}
