package main

import (
	"fmt"
)

type sendCodeCallback interface {
	SendCode(string, string, int) error
}

type phoneCodeData struct {
        authKeyId        int64
        phoneNumber      string
        code             string
        codeHash         string
        codeExpired      int32
        sentCodeType     int
        flashCallPattern string
        nextCodeType     int
        state            int
        // dataType: kDBTypeCreate, kDBTypeLoad
        dataType         int
        tableId          int64
        codeCallback     sendCodeCallback
}

// implement sendCodeCallback on phoneCodeData
func (code *phoneCodeData) SendCode(string, string, int) error {
        fmt.Println("SendCode()");

        return nil
}

func MakeCodeData(authKeyId int64, phoneNumber string) *phoneCodeData {
        // TODO(@benqi): Independent Unified Messaging Push System
        // Check if phpne exists. If it exists online, decide whether to send it via SMS or send it through other clients.
        // Transparent transmission of AuthId, UserId, terminal type, etc.
        // Check if the TransactionHash that satisfies the condition exists. Possible conditions:
        //  1. is_deleted !=0 and now - created_at < 15 分钟
        //


        code := &phoneCodeData{
                authKeyId:   authKeyId,
                phoneNumber: phoneNumber,
                // state:       kCodeStateNone,
                // dataType:    kDBTypeCreate,
        }

	code.codeCallback = code

        return code
}

func (code *phoneCodeData) doSendCodeCallback() error {
        fmt.Println("doSendCodeCallback()")
        if code.codeCallback != nil {
                fmt.Println("code.codeCallback defined. Sending Code...")
                return code.codeCallback.SendCode(code.code, code.codeHash, code.sentCodeType)
        }

        fmt.Println("code.codeCallback NOT defined. Returning dummy truth. Test Environment.")
        // TODO(@benqi): The test environment is sent successfully by default.
        return nil
}

// implement 
// auth.sendCode
func (code *phoneCodeData) DoSendCode(phoneRegistered, allowFlashCall, currentNumber bool, apiId int32, apiHash string) error {

	/*
        code.checkDataType(kDBTypeCreate)

        // Use the easiest way to create a new one each time.
        sentCodeType, nextCodeType := makeCodeType(phoneRegistered, allowFlashCall, currentNumber)
        // TODO(@benqi): gen rand number
        code.code = "12345"
        code.codeHash = crypto.GenerateStringNonce(16)
        code.codeExpired = int32(time.Now().Unix() + 15*60)
        code.sentCodeType = sentCodeType
        code.nextCodeType = nextCodeType
        */

        err := code.doSendCodeCallback()
        if err != nil {
                // glog.Error(err)
		fmt.Println("error doSendCodeCallback ")
                return err
        }

	/*
        // save
        do := &dataobject.AuthPhoneTransactionsDO
                AuthKeyId:        code.authKeyId,
                PhoneNumber:      code.phoneNumber,
                Code:             code.code,
                CodeExpired:      code.codeExpired,
                TransactionHash:  code.codeHash,
                SentCodeType:     int8(code.sentCodeType),
                FlashCallPattern: code.flashCallPattern,
                NextCodeType:     int8(code.nextCodeType),
                State:                    kCodeStateSent,
                ApiId:            apiId,
                ApiHash:          apiHash,
                CreatedTime:      time.Now().Unix(),
        }
        code.tableId = dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).Insert(do)
        //// TODO(@benqi):
        //lastCreatedAt := time.Unix(time.Now().Unix()-15*60, 0).Format("2006-01-02 15:04:05")
        //do := dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).SelectByPhoneAndApiIdAndHash(code.phoneNumber, apiId, apiHash, lastCreatedAt)
        //if do == nil {
        //} else {
        //      // TODO(@benqi): FLOOD_WAIT_X, too many attempts, please try later.
        //}
        */
        return nil
}

func main() {

	fmt.Println("Send Code ...");

        var authId int64
	var phoneNumber string
	var phoneRegistered bool
	var allowFlashcall bool
	var currentNumber bool
        var apiId int32
	var apiHash string

	authId = 123456
	phoneNumber = "+27729745087"
	phoneRegistered = false
	allowFlashcall = true
	currentNumber = false
	apiId = 1234567890
	apiHash = "1234567890123456"

        code := MakeCodeData(authId, phoneNumber)
	fmt.Println("To Phone Number: ", code.phoneNumber)

        var err = code.DoSendCode(phoneRegistered, allowFlashcall, currentNumber, apiId, apiHash)
        if err != nil {
                // glog.Error(err)
		fmt.Println("error occured: %v", err)
                // return nil, err
        }

	fmt.Println("Auth Code Successfully Sent.");
}
