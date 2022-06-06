package routers

import (
	"merchant-api/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init - Initing routers  -
func Init() *gin.Engine {
	routingEngine := gin.Default()
	// CORS
	routingEngine.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowHeaders:     []string{"content-type", "Authorization", "clientID", "sessionEnabled", "X-Requested-With", "Cache-Control", "Pragma", "X-SESSION-TOKEN", "If-Modified-Since", "timeZone", "Timestamp"},
		AllowCredentials: true,
		MaxAge:           5 * time.Minute,
	}))

	// Recovery middleware
	// routingEngine.Use(gin.Recovery())
	auth := &controllers.AuthController{}
	// userControlledRouters := routingEngine.Group("/v1/user")
	// userActionsRouters := routingEngine.Group("/v1/user/profile")
	// defaultRouters := routingEngine.Group("/")
	// defaultRouters.Use(auth.Authenticate)
	// userActionsRouters.Use(auth.Authenticate)
	// commonController := routingEngine.Group("/v1/commons")
	// locationController := routingEngine.Group("/v1/locations")
	// /******************************************************************************
	//                             UserController
	// ******************************************************************************/
	// uc := &controlleres.UserController{}
	// userControlledRouters.POST("/signup", uc.Signup)
	// userControlledRouters.POST("/login", uc.Login)
	// uac := &controlleres.UserActionsController{}
	// userActionsRouters.PUT("/update", uac.Update)
	// userActionsRouters.PUT("/update/:id", uac.Update)
	// userActionsRouters.GET("/view", uc.UserProfile)
	routers := routingEngine.Group("/")
	routers.Use(auth.Authenticate)
	/******************************************************************************
	                            DefaultController
	******************************************************************************/
	//Enabling the token based authentication
	mrc := &controllers.MerchantController{}
	routers.GET("/merchants", mrc.Find)
	routers.GET("/merchant/:id", mrc.FindOne)
	routers.POST("/merchant", mrc.Create)
	routers.PUT("/merchant/:id", mrc.Update)
	routers.DELETE("/merchant/:id", mrc.Delete)

	/******************************************************************************
	                            CommonController
	******************************************************************************/
	//Enabling the token based authentication
	mmc := &controllers.MemberController{}
	routers.GET("/members", mmc.Find)
	routers.GET("/member/:id", mmc.FindOne)
	routers.POST("/member", mmc.Create)
	routers.PUT("/member/:id", mmc.Update)
	routers.DELETE("/member/:id", mmc.Delete)
	routers.GET("/merchant/members/:code", mmc.FindMembersByMerchantCode)

	return routingEngine
}
