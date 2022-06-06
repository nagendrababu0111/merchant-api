package auth

import (
	"errors"
	"strings"

	"merchant-api/constants"
	"merchant-api/utils/db/mongo"
	"merchant-api/utils/types"
)

//Authenticate the given service
func Authenticate(authorization string) error {
	query := make(types.Map)
	tokens := strings.Split(authorization, " ")
	if len(tokens) != 2 {
		return errors.New("Invalid Authorization")
	}
	query["token"] = tokens[1]
	count := mongo.GetTableRowsCount(constants.DATA_BASE, constants.USER_COLL, query)
	if count == 0 {
		return errors.New("Invalid Authorization")
	}
	return nil
}
