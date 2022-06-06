package merchant_test

import (
	"merchant-api/services/merchant"
	"merchant-api/utils/commons"
	"merchant-api/utils/config"
	"merchant-api/utils/types"
	"testing"
)

var id string

func init() {
	id = commons.GenerateID("")
	config.LoadJSON("../../config/development.json")
	commons.InitIndexes()
}

func TestMembersCreate(t *testing.T) {
	validData := types.Map{
		"id":          id,
		"code":        "MRC_001",
		"category":    "Goods & Supply",
		"description": "Goods and Supply",
	}

	_, err := merchant.Create(validData)
	if err != nil {
		t.Error("error while create", err)
	}
}

func TestMembersUpdate(t *testing.T) {
	validData := types.Map{
		"id":          id,
		"code":        "MRC_001",
		"category":    "Goods & Supply",
		"description": "Goods And Supply",
	}

	err := merchant.Update(id, validData)
	if err != nil {
		t.Error("error while create", err)
	}
}

func TestMembersFind(t *testing.T) {
	_, _, err := merchant.Find(types.Map{}, types.Page{})
	if err != nil {
		t.Error("error while get", err)
	}
}

func TestMembersFindOne(t *testing.T) {
	_, err := merchant.FindOne(types.Map{"_id": id})
	if err != nil {
		t.Error("error while get", err)
	}
}

func TestMembersDelete(t *testing.T) {
	err := merchant.Delete(types.Map{"_id": id})
	if err != nil {
		t.Error("error while get", err)
	}
}
