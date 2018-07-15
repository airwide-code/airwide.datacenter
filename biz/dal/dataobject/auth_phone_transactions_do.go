/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package dataobject

type AuthPhoneTransactionsDO struct {
	Id               int64  `db:"id"`
	AuthKeyId        int64  `db:"auth_key_id"`
	PhoneNumber      string `db:"phone_number"`
	Code             string `db:"code"`
	CodeExpired      int32  `db:"code_expired"`
	TransactionHash  string `db:"transaction_hash"`
	SentCodeType     int8   `db:"sent_code_type"`
	FlashCallPattern string `db:"flash_call_pattern"`
	NextCodeType     int8   `db:"next_code_type"`
	State            int8   `db:"state"`
	ApiId            int32  `db:"api_id"`
	ApiHash          string `db:"api_hash"`
	Attempts         int32  `db:"attempts"`
	CreatedTime      int64  `db:"created_time"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
	IsDeleted        int8   `db:"is_deleted"`
}
