package schema

import (
	"time"
)

// merchant Schema
type Merchant struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Code        string    `json:"code" bson:"code" validate:"nonzero"`
	Category    string    `json:"category" bson:"category" validate:"nonzero"`
	Description string    `json:"description" bson:"description" validate:"nonzero"`
	VersionDate time.Time `json:"version_date" bson:"version_date"`
}

// Team Member Schema
type Member struct {
	Id           string    `json:"id" bson:"_id,omitempty"`
	EmailId      string    `json:"email_id" bson:"email_id" validate:"nonzero"`
	MerchantCode string    `json:"merchant_code" bson:"merchant_code" validate:"nonzero"`
	FirstName    string    `json:"first_name" bson:"first_name" validate:"nonzero"`
	LastName     string    `json:"last_name" bson:"last_name" validate:"nonzero"`
	VersionDate  time.Time `json:"version_date" bson:"version_date"`
}

type User struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	EmailId     string    `json:"email_id" bson:"email_id"`
	Token       string    `json:"token" bson:"token"`
	Password    string    `json:"password" bson:"password"`
	FirstName   string    `json:"first_name" bson:"first_name"`
	LastName    string    `json:"last_name" bson:"last_name"`
	UserName    string    `json:"user_name" bson:"user_name"`
	VersionDate time.Time `json:"version_date" bson:"version_date"`
}
