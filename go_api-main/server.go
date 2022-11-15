package main

import (
	"fmt"
	"go_api/config"
	"go_api/controller"
	"go_api/logger"
	"go_api/middleware"
	"go_api/repository"
	"go_api/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	r.Run(":8080")
}

type foo struct {
	name string
	age  int
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
