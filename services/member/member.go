package member

import (
	"encoding/json"
	"errors"
	"log"
	"merchant-api/constants"
	"merchant-api/schema"
	"merchant-api/utils/commons"
	"time"

	"merchant-api/utils/db/mongo"
	"merchant-api/utils/types"

	"gopkg.in/mgo.v2"
	validator "gopkg.in/validator.v2"
)

// FindOne  -
func FindOne(params types.Map) (types.Map, error) {
	row, err := mongo.FindOne(constants.DATA_BASE, constants.MEMBER_COLL, params)
	if err != nil || len(row) == 0 {
		log.Println("error occurred while getting records from ", constants.MEMBER_COLL, params, err)
		return row, errors.New(constants.NOT_FOUND)
	}
	return row, err
}

// FindMembersByMerchantCode  -
func FindMembersByMerchantCode(params types.Map, paging types.Page) ([]schema.Member, int, error) {
	rows := make([]schema.Member, 0)
	count, err := mongo.FindRecords(constants.DATA_BASE, constants.MEMBER_COLL, params, paging, &rows)
	if err != nil || len(rows) == 0 {
		log.Println("error occurred while getting records from ", constants.MEMBER_COLL, params, err)
		return rows, 0, errors.New(constants.NOT_FOUND)
	}
	return rows, count, err
}

// Find  -
func Find(query types.Map, paging types.Page) ([]schema.Member, int, error) {
	rows := make([]schema.Member, 0)
	count, err := mongo.FindRecords(constants.DATA_BASE, constants.MEMBER_COLL, query, paging, &rows)
	if err != nil {
		log.Println("error occurred while getting records from ", constants.MEMBER_COLL, query, err)
		return rows, 0, errors.New(constants.NOT_FOUND)
	}

	return rows, count, err
}

// Create  -
func Create(data types.Map) (interface{}, error) {
	var member = schema.Member{}

	if !isValidMember(data, &member) {
		return nil, errors.New(constants.INV_INP)
	}
	if !commons.IsEmailValid(member.EmailId) {
		return nil, errors.New(constants.INV_EMAIL)
	}
	if !isValidMerchantCode(member.MerchantCode) {
		return nil, errors.New(constants.INV_MERCHANT_CODE)
	}

	member.Id = commons.GenerateID(member.Id)
	member.VersionDate = time.Now().UTC()

	row, err := mongo.Create(constants.DATA_BASE, constants.MEMBER_COLL, member)
	if err != nil {
		log.Println("error occurred while creating records from ", constants.MEMBER_COLL, data, err)
		err = HandleDuplicateError(err)
		return nil, err
	}
	return row, err
}

// Update -
func Update(id string, data types.Map) error {
	var member = schema.Member{}

	if !isValidMember(data, &member) {
		return errors.New(constants.INV_INP)
	}
	if !commons.IsEmailValid(member.EmailId) {
		return errors.New(constants.INV_EMAIL)
	}
	if !isValidMerchantCode(member.MerchantCode) {
		return errors.New(constants.INV_MERCHANT_CODE)
	}
	member.VersionDate = time.Now().UTC()
	err := mongo.UpdateOne(constants.DATA_BASE, constants.MEMBER_COLL, id, member)
	if err != nil {
		log.Println("error occurred while updating records from ", constants.MEMBER_COLL, data, err)
		return HandleDuplicateError(err)
	}
	return nil
}

// Delete -
func Delete(data types.Map) error {
	if len(data) == 0 {
		return errors.New(constants.INV_INP)
	}
	err := mongo.Delete(constants.DATA_BASE, constants.MEMBER_COLL, data)
	if err != nil {
		log.Println("error occurred while deleting records from ", constants.MEMBER_COLL, data, err)
		return errors.New(constants.NOT_FOUND)
	}
	return nil
}

func isValidMember(inp types.Map, member *schema.Member) bool {
	b, err := json.Marshal(inp)
	err = json.Unmarshal(b, member)
	if (schema.Member{} == *member || err != nil) {
		return false
	}

	err = validator.Validate(member)
	if err != nil {
		log.Println("validation error ", err)
		return false
	}

	return true

}

func isValidMerchantCode(code string) bool {
	return mongo.GetTableRowsCount(constants.DATA_BASE, constants.MARCHANT_COLL, types.Map{"code": code}) > 0
}

func HandleDuplicateError(err error) error {
	if mgo.IsDup(err) {
		return errors.New(constants.DUPLICATE_EMAIL)
	}
	return errors.New(constants.INTERNAL_ERR)
}
