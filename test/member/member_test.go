package member_test

import (
	"merchant-api/services/member"
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
		"id":            id,
		"email_id":      "nagendrababu0111@gmail.com",
		"merchant_code": "MRC_001",
		"first_name":    "Ngendra",
		"last_name":     "Yandra",
	}

	_, err := member.Create(validData)
	if err != nil {
		t.Error("error while create", err)
	}
}

func TestMembersUpdate(t *testing.T) {
	// inValidData := types.Map{}
	validData := types.Map{
		"id":            id,
		"email_id":      "nagendrababu0111@gmail.com",
		"merchant_code": "MRC_001",
		"first_name":    "Nagendra Babu",
		"last_name":     "Yandra",
	}

	err := member.Update(id, validData)
	if err != nil {
		t.Error("error while create", err)
	}
}

func TestMembersFind(t *testing.T) {
	_, _, err := member.Find(types.Map{}, types.Page{})
	if err != nil {
		t.Error("error while get", err)
	}
}

func TestMembersFindByMerchantCode(t *testing.T) {
	_, _, err := member.FindMembersByMerchantCode(types.Map{"merchant_code": "MRC_001"}, types.Page{})
	if err != nil {
		t.Error("error while get", err)
	}
}

func TestMembersFindOne(t *testing.T) {
	_, err := member.FindOne(types.Map{"_id": id})
	if err != nil {
		t.Error("error while get", err)
	}
}

func TestMembersDelete(t *testing.T) {
	err := member.Delete(types.Map{"_id": id})
	if err != nil {
		t.Error("error while get", err)
	}
}
