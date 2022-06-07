package constants

// Database
const DATA_BASE = "merchant"

// Tables
const MARCHANT_COLL = "merchant"
const MEMBER_COLL = "member"
const USER_COLL = "users"

// Error Codes

const (
	NOT_FOUND          = "Not Found"
	INV_INP            = "Hey! Can you please check and provide the valid input?"
	INV_EMAIL          = "Hey! Can you provide the valid email id?"
	DUPLICATE_MRC_CODE = "Hey! Merchant exists with same Code"
	DUPLICATE_EMAIL    = "Hey! Team Member exists with same email"
	INTERNAL_ERR       = "Hey! Internal Server error occurred, Please contact your administrator"
	INV_MERCHANT_CODE  = "Hey! Can you please check the merchant code that you provided?"
)
