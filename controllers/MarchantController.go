package controllers

import (
	"errors"
	// "strings"

	"merchant-api/utils/types"

	"merchant-api/services/merchant"

	"github.com/gin-gonic/gin"
)

// AuthController -
type MerchantController struct {
	RootController
}

// FindOne -
func (mc *MerchantController) FindOne(ctx *gin.Context) {
	// mer
	result, err := merchant.FindOne(types.Map{"_id": ctx.Param("id")})
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSingleEntityResponse(ctx, result)
}

// Find -
func (mc *MerchantController) Find(ctx *gin.Context) {
	query, paging, err := mc.GetQuery(ctx)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	result, count, err := merchant.Find(query, paging)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeListEntityResponse(ctx, result, count)
}

// Create -
func (mc *MerchantController) Create(ctx *gin.Context) {
	data := mc.GetReqBody(ctx)
	if len(data) == 0 {
		mc.MakeAnErrorResponse(ctx, errors.New("Invalid input received"))
		return
	}
	row, err := merchant.Create(data)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSingleEntityResponse(ctx, row)
}

// Update -
func (mc *MerchantController) Update(ctx *gin.Context) {
	data := mc.GetReqBody(ctx)
	err := merchant.Update(ctx.Param("id"), data)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSuccessResponse(ctx)
}

// Delete -
func (mc *MerchantController) Delete(ctx *gin.Context) {
	err := merchant.Delete(types.Map{"_id": ctx.Param("id")})
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSuccessResponse(ctx)
}
