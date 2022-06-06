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
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           5 * time.Minute,
	}))

	auth := &controllers.AuthController{}
	routers := routingEngine.Group("/")

	//Enabling the token based authentication
	routers.Use(auth.Authenticate)
	/******************************************************************************
	                            Merchant Controller
	******************************************************************************/
	mrc := &controllers.MerchantController{}
	routers.GET("/merchants", mrc.Find)
	routers.GET("/merchant/:id", mrc.FindOne)
	routers.POST("/merchant", mrc.Create)
	routers.PUT("/merchant/:id", mrc.Update)
	routers.DELETE("/merchant/:id", mrc.Delete)

	/******************************************************************************
	                            Member Controller
	******************************************************************************/
	mmc := &controllers.MemberController{}
	routers.GET("/members", mmc.Find)
	routers.GET("/member/:id", mmc.FindOne)
	routers.POST("/member", mmc.Create)
	routers.PUT("/member/:id", mmc.Update)
	routers.DELETE("/member/:id", mmc.Delete)
	routers.GET("/merchant/members/:code", mmc.FindMembersByMerchantCode)

	return routingEngine
}
