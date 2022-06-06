package controllers

import (
	"errors"
	// "strings"

	"merchant-api/utils/types"

	"merchant-api/services/member"

	"github.com/gin-gonic/gin"
)

// MemberController -
type MemberController struct {
	RootController
}

// FindOne -
func (mc *MemberController) FindOne(ctx *gin.Context) {
	// mer
	result, err := member.FindOne(types.Map{"_id": ctx.Param("id")})
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSingleEntityResponse(ctx, result)
}

// FindMembersByMerchantCode
// http://localhost:8080/merchant/members/:code?limit=1&page=2
func (mc *MemberController) FindMembersByMerchantCode(ctx *gin.Context) {

	query := types.Map{"merchant_code": ctx.Param("code")}
	p, _ := mc.Paging(ctx)
	result, count, err := member.FindMembersByMerchantCode(query, p)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeListEntityResponse(ctx, result, count)
}

// Find -
func (mc *MemberController) Find(ctx *gin.Context) {
	query, paging, err := mc.GetQuery(ctx)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	result, count, err := member.Find(query, paging)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeListEntityResponse(ctx, result, count)
}

// Create -
func (mc *MemberController) Create(ctx *gin.Context) {
	data := mc.GetReqBody(ctx)
	if len(data) == 0 {
		mc.MakeAnErrorResponse(ctx, errors.New("Invalid input received"))
		return
	}
	row, err := member.Create(data)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSingleEntityResponse(ctx, row)
}

// Update -
func (mc *MemberController) Update(ctx *gin.Context) {
	data := mc.GetReqBody(ctx)
	err := member.Update(ctx.Param("id"), data)
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSuccessResponse(ctx)
}

// Delete -
func (mc *MemberController) Delete(ctx *gin.Context) {
	err := member.Delete(types.Map{"_id": ctx.Param("id")})
	if err != nil {
		mc.MakeAnErrorResponse(ctx, err)
		return
	}
	mc.MakeSuccessResponse(ctx)
}
