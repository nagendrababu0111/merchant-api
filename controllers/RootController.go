package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"merchant-api/utils/commons"
	"merchant-api/utils/types"

	"github.com/gin-gonic/gin"
)

// RootController -
type RootController struct {
}

// GetParams -
func (rc RootController) GetParams(g *gin.Context) types.Map {
	params := make(types.Map, 0)
	query := g.Request.URL.Query()
	for k, v := range query {
		if len(v) == 0 {
			continue
		}
		params[k] = v[0]
	}
	return params
}

// GetReqBody -
func (rc RootController) GetReqBody(ctx *gin.Context) types.Map {
	reqBody := make(types.Map, 0)
	content, _ := ioutil.ReadAll(ctx.Request.Body)
	json.Unmarshal(content, &reqBody)
	return reqBody
}

// GetHeadrs
func (rc RootController) GetHeadrs(ctx *gin.Context) map[string][]string {
	headersMap := make(map[string][]string)
	for k, v := range ctx.Request.Header {
		headersMap[k] = v
	}
	return headersMap
}

// // MakeCustomSuccessResponse -
// func (rc RootController) MakeCustomSuccessResponse(ctx *gin.Context, result types.Map) {
// 	result["status"] = "success"
// 	ctx.Header("Content-Type", "application/json")
// 	ctx.JSON(http.StatusOK, result)
// }

// MakeSingleEntityResponse -
func (rc RootController) MakeSingleEntityResponse(ctx *gin.Context, result interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, types.Map{"status": "success", "row": result})
}

// MakeListEntityResponse -
func (rc RootController) MakeListEntityResponse(ctx *gin.Context, result interface{}, total interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, types.Map{"status": "success", "rows": result, "total": total})
}

// MakeSuccessResponse -
func (rc RootController) MakeSuccessResponse(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, types.Map{"status": "success"})
}

// MakeAnErrorResponse -
func (rc RootController) MakeAnErrorResponse(ctx *gin.Context, err error) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, types.Map{"status": "failed", "message": err.Error()})
}

func (rc RootController) GetQuery(c *gin.Context) (types.Map, types.Page, error) {
	query := make(types.Map, 0)
	for k := range c.Request.URL.Query() {
		query[k] = c.Query(k)
	}

	if !commons.IsEmptyI(query["inQuery"]) {
		temp, err := commons.InterfaceToMap(query["inQuery"])
		if err != nil {
			log.Println("error in parsing in query ...", err)
		}
		enrichInQuery(query, temp)
		delete(query, "inQuery")
	}

	paging, err := rc.Paging(c)
	if err != nil {
		return nil, types.Page{}, err
	}
	delete(query, "page")
	delete(query, "limit")
	delete(query, "sortBy")
	return query, paging, nil
}

func enrichInQuery(query types.Map, input types.Map) types.Map {
	for k, v := range input {
		query[k] = types.Map{"$in": v}
	}
	return query
}

//Paging ...
func (rc RootController) Paging(c *gin.Context) (types.Page, error) {

	page := commons.ToInt(c.Query("page"), 1)
	limit := commons.ToInt(c.Query("limit"), 15)

	p := types.Page{
		Page:  page,
		Limit: limit,
		Skip:  (page - 1) * limit,
	}

	return p, nil
}
