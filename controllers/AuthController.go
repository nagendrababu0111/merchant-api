package controllers

import (
	"merchant-api/services/auth"
	"merchant-api/utils/types"

	"github.com/gin-gonic/gin"
)

// AuthController -
type AuthController struct {
	RootController
}

//Authenticate
func (uc *AuthController) Authenticate(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")
	// auth.
	err := auth.Authenticate(authorization)
	if err != nil {
		ctx.AbortWithStatusJSON(401, types.Map{"status": "Fail", "message": "401 Unauthorized"})
		// uc.Send401(ctx)
	}
	// ctx.Request.Header.Add("authInfo", commons.ToJSONString(userData))
}
