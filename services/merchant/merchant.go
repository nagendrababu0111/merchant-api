package merchant

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
	row, err := mongo.FindOne(constants.DATA_BASE, constants.MARCHANT_COLL, params)
	if err != nil || len(row) == 0 {
		log.Println("error occurred while getting records from ", constants.MARCHANT_COLL, params, err)
		return row, errors.New(constants.NOT_FOUND)
	}
	return row, err
}

// Find  -
func Find(query types.Map, paging types.Page) ([]schema.Merchant, int, error) {
	rows := make([]schema.Merchant, 0)
	count, err := mongo.FindRecords(constants.DATA_BASE, constants.MARCHANT_COLL, query, paging, &rows)
	if err != nil {
		log.Println("error occurred while getting records from ", constants.MARCHANT_COLL, query, err)
		return rows, 0, errors.New(constants.NOT_FOUND)
	}

	return rows, count, err
}

// Create  -
func Create(data types.Map) (interface{}, error) {
	var merchant = schema.Merchant{}

	if !isValidMerchant(data, &merchant) {
		return nil, errors.New(constants.INV_INP)
	}
	merchant.Id = commons.GenerateID(merchant.Id)
	merchant.VersionDate = time.Now().UTC()

	row, err := mongo.Create(constants.DATA_BASE, constants.MARCHANT_COLL, merchant)
	if err != nil {
		log.Println("error occurred while creating records from ", constants.MARCHANT_COLL, data, err)
		err = HandleDuplicateError(err)
		return nil, err
	}
	return row, err
}

// Update -
func Update(id string, data types.Map) error {
	var merchant = schema.Merchant{}

	if !isValidMerchant(data, &merchant) {
		return errors.New(constants.INV_INP)
	}
	log.Println("mer", merchant, data)

	merchant.VersionDate = time.Now().UTC()
	err := mongo.UpdateOne(constants.DATA_BASE, constants.MARCHANT_COLL, id, merchant)
	if err != nil {
		log.Println("error occurred while updating records from ", constants.MARCHANT_COLL, data, err)
		return HandleDuplicateError(err)
	}
	return nil
}

// Delete -
func Delete(data types.Map) error {
	if len(data) == 0 {
		return errors.New(constants.INV_INP)
	}
	err := mongo.Delete(constants.DATA_BASE, constants.MARCHANT_COLL, data)
	if err != nil {
		log.Println("error occurred while deleting records from ", constants.MARCHANT_COLL, data, err)
		return errors.New(constants.NOT_FOUND)
	}
	return nil
}

func isValidMerchant(inp types.Map, merchant *schema.Merchant) bool {
	b, err := json.Marshal(inp)
	err = json.Unmarshal(b, merchant)
	if (schema.Merchant{} == *merchant || err != nil) {
		log.Println("erreor while unmarshal", err, merchant)
		return false
	}

	err = validator.Validate(merchant)
	if err != nil {
		log.Println("validation error ", err)
		return false
	}

	return true

}

func HandleDuplicateError(err error) error {
	if mgo.IsDup(err) {
		return errors.New(constants.DUPLICATE_MRC_CODE)
	}
	return errors.New(constants.INTERNAL_ERR)
}
