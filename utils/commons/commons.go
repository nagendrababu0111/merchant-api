package commons

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"merchant-api/schema"
	"merchant-api/utils/types"
	"regexp"
	"time"

	"merchant-api/constants"
	"merchant-api/utils/db/mongo"
	"reflect"
	"strconv"
	"strings"
)

func InitIndexes() {
	mongo.CreateUniqueIndex(constants.DATA_BASE, constants.MARCHANT_COLL, []string{"code"})
	mongo.CreateUniqueIndex(constants.DATA_BASE, constants.MEMBER_COLL, []string{"email_id"})
}

//ToJSONString -
func ToJSONString(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes[:])
}

//IsEmptyI -
func IsEmptyI(val interface{}) bool {
	if val == nil {
		return true
	}
	if reflect.TypeOf(val) == reflect.TypeOf("") && len(strings.TrimSpace(val.(string))) == 0 {
		return true
	}
	return false
}

// GenerateID - Generate a new time based ID
func GenerateID(defaultId string) string {
	if IsEmptyI(defaultId) {
		defaultId = GenerateRandomID(10)
	}
	return defaultId
}

func GenerateRandomID(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func InterfaceToMap(input interface{}) (types.Map, error) {
	var output types.Map
	err := json.Unmarshal([]byte(ToJSONString(input)), &output)
	return output, err
}

func ToInt(input string, defaultValue int) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		return defaultValue
	}
	return value
}

func StrToMap(str string) (types.Map, error) {
	var output types.Map
	err := json.Unmarshal([]byte(str), &output)
	return output, err
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func SeedUsers() {
	user := schema.User{
		Id:          GenerateID(""),
		Token:       "3f5ac11913d7e217785c1d3e07d52c86",
		FirstName:   "Nagendra Babu",
		LastName:    "Yandra",
		EmailId:     "nagendrababu0111@gmail.com",
		UserName:    "nagendra.yandra",
		VersionDate: time.Now().UTC(),
		Password:    base64.StdEncoding.EncodeToString([]byte("hello@Nag")),
	}
	mongo.Create(constants.DATA_BASE, constants.USER_COLL, user)
}
