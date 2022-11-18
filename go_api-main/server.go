package main

import (
	"fmt"
	"go_api/config"
	"go_api/controller"
	"go_api/dto"
	"go_api/helper"
	"go_api/logger"
	"go_api/middleware"
	"go_api/repository"
	"go_api/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	pingRepository    repository.PingRepository    = repository.NewPingRepository(db)
	partnerRepository repository.PartnerRepository = repository.NewPartnerRepository(db)

	jwtService     service.JWTService     = service.NewJWTService()
	authService    service.AuthService    = service.NewAuthService(partnerRepository)
	userService    service.UserService    = service.NewUserService(userRepository)
	pingService    service.PingService    = service.NewPingService(pingRepository)
	partnerService service.PartnerService = service.NewPartnerService(partnerRepository)

	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	pingController    controller.PingController    = controller.NewPingController(pingService)
	partnerController controller.PartnerController = controller.NewPartnerController(partnerService, jwtService)
)

func main() {
	logger.InfoLogger.Println("Start")
	defer config.CloseDatabaseConnection(db)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}
	r.POST("/partner", authController.RegisterPartner)

	r.GET("/ping", pingController.Ping)
	r.GET("/call", callapi)
	r.POST("/link", Link)
	r.Run(":8080")
}

func callapi(ctx *gin.Context) {
	response, err := http.Get("https://mocki.io/v1/16993188-338c-4667-a4eb-164723b765c8")
	logger.InfoLogger.Println("response: ", response.Body)
	logger.InfoLogger.Println("err: ", err)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	ctx.String(http.StatusOK, string(responseData))
	fmt.Println(string(responseData))
}

func Link(ctx *gin.Context) {
	var request dto.Request
	var response helper.WalletResponse
	errDTO := ctx.ShouldBind(&request)
	logger.InfoLogger.Println("rq: ", request)

	authHeader := ctx.GetHeader("Authorization")
	logger.InfoLogger.Println("authorization: ", authHeader)

	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token, errToken := jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errDTO.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims.VerifyExpiresAt(time.Now().Unix(), true) {

	}
	logger.InfoLogger.Println("claims: ", claims.VerifyExpiresAt(time.Now().UnixMicro(), true))
	logger.InfoLogger.Println("cmp: ", claims)

	var ress string
	db.Raw("SELECT  service_name FROM  wallet.api_provider_mapping WHERE Status= 'Y' and provider =? and operation = ? ", request.Header.Providerno, "wallet2Bank").Scan(&ress)
	logger.InfoLogger.Println("ress: ", ress)
	response.Body.BankRef = "001"
	// response := helper.BuildLinkRequest(headerLinkDTO, bodyLinkDTO)
	ctx.JSON(http.StatusOK, response)
}
