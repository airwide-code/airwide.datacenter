/*
 *  Copyright (c) 2018, https://github.com/airwide-code
 *  All rights reserved.
 *
 *
 *
 */

package auth

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dao"
	"github.com/airwide-code/airwide.datacenter/biz/dal/dataobject"
	"github.com/airwide-code/airwide.datacenter/mtproto"
	"time"
	"github.com/airwide-code/airwide.datacenter/baselib/crypto"
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
)

// TODO(@benqi): 当前测试环境code统一为"12345"
// TODO(@benqi): 限制同一个authKeyId
// TODO(@benqi): 使用redis

const (
	kCodeType_None      = 0
	kCodeType_App       = 1
	kCodeType_Sms       = 2
	kCodeType_Call      = 3
	kCodeType_FlashCall = 4
)

//auth.codeTypeSms#72a3158c = auth.CodeType;
//auth.codeTypeCall#741cd3e3 = auth.CodeType;
//auth.codeTypeFlashCall#226ccefb = auth.CodeType;
//
//auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
//auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
//auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
//auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;

// dataType 实现 lazy create
const (
	kDBTypeNone   = 0
	kDBTypeCreate = 1
	kDBTypeLoad   = 2
	kDBTypeUpdate = 3
	kDBTypeDelete = 3
)

const (
	kCodeStateNone    = 0
	kCodeStateOk      = 1
	kCodeStateSent    = 2
	kCodeStateSignIn  = 3
	kCodeStateSignUp  = 4
	kCodeStateDeleted = -1
	kCodeStateTimeout = -2
)

type sendCodeCallback interface {
	SendCode(string, string, int) error
}

// TODO(@benqi): Add phone region
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
	dataType     	 int
	tableId          int64
	codeCallback     sendCodeCallback
}

// implement sendCodeCallback on phoneCodeData
func (code *phoneCodeData) SendCode(theCode string, string, int) error {
	fmt.Println("SendCode()");

	endpoint := "https://api.twilio.com/2010-04-01/Accounts/AC25c34e873eb0348a2a7b9510f9282319/Messages.json"
        v := url.Values{}
        v.Set("To", code.phoneName)
        v.Add("From", "+14342774779")
        v.Add("Body", "Your Airwide Code is " + theCode)
        payload := strings.NewReader(v.Encode())

        var username string = "AC25c34e873eb0348a2a7b9510f9282319"
        var passwd string = "c052c7e3068c0f1e64ba5067836b10d4"

        req, _ := http.NewRequest("POST", endpoint, payload)
        req.Header.Add("content-type", "application/x-www-form-urlencoded")
        req.Header.Add("cache-control","no-cache")
        req.SetBasicAuth(username, passwd)

        res, err := http.DefaultClient.Do(req)
        if err != nil {
                fmt.Println("Fatal error occured")
        }

        defer res.Body.Close();
        body, _ := ioutil.ReadAll(res.Body);

        fmt.Println(string(body));

        // TODO cater for http redirects. https://stackoverflow.com/questions/16673766/basic-http-auth-in-go

	return nil;
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
// 由params(phoneRegistered, allowFlashCall, currentNumber)确定sentType和nextType
func makeCodeType(phoneRegistered, allowFlashCall, currentNumber bool) (int, int) {

	//if phoneRegistered {
	//	// TODO(@benqi): check other session online
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeApp{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//} else {
	//	// TODO(@benqi): sentCodeTypeFlashCall and sentCodeTypeCall, nextType
	//	// telegramd, we only use sms
	//	authSentCodeType := &mtproto.TLAuthSentCodeTypeSms{Data2: &mtproto.Auth_SentCodeType_Data{
	//		Length: code.GetPhoneCodeLength(),
	//	}}
	//	authSentCode.SetType(authSentCodeType.To_Auth_SentCodeType())
	//
	//	// TODO(@benqi): nextType
	//	// authSentCode.SetNextType()
	//}

	sentCodeType := kCodeType_App
	nextCodeType := kCodeType_None
	return sentCodeType, nextCodeType
}

func makeAuthCodeType(codeType int) *mtproto.Auth_CodeType {
	switch codeType {
	case kCodeType_Sms:
		return mtproto.NewTLAuthCodeTypeSms().To_Auth_CodeType()
	case kCodeType_Call:
		return mtproto.NewTLAuthCodeTypeCall().To_Auth_CodeType()
	case kCodeType_FlashCall:
		return mtproto.NewTLAuthCodeTypeFlashCall().To_Auth_CodeType()
	default:
		return nil
	}
}

func makeAuthSentCodeType(codeType, codeLength int, pattern string) (authSentCodeType *mtproto.Auth_SentCodeType) {
	switch codeType {
	case kCodeType_App:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeApp,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length: int32(codeLength),
			},
		}
	case kCodeType_Sms:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeSms,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length: int32(codeLength),
			},
		}
	case kCodeType_Call:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeCall,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length: int32(codeLength),
			},
		}
	case kCodeType_FlashCall:
		authSentCodeType = &mtproto.Auth_SentCodeType{
			Constructor: mtproto.TLConstructor_CRC32_auth_sentCodeTypeFlashCall,
			Data2: &mtproto.Auth_SentCodeType_Data{
				Length:  int32(codeLength),
				Pattern: pattern,
			},
		}
	default:
		// code bug.
		err := fmt.Errorf("invalid sentCodeType: %d", codeType)
		glog.Error("makeAuthSentCodeType - ", err)
		panic(err)
	}

	return
}

func MakeCodeData(authKeyId int64, phoneNumber string) *phoneCodeData {
	// TODO(@benqi): Independent Unified Messaging Push System 
	// Check if phpne exists. If it exists online, decide whether to send it via SMS or send it through other clients. 
	// Transparent transmission of AuthId, UserId, terminal type, etc. 
	// Check if the TransactionHash that satisfies the condition exists. Possible conditions: 
	//  1. is_deleted !=0 and now - created_at < 15 分钟
	//

	// sentCodeType, nextCodeType := makeCodeType(phoneRegistered, allowFlashCall, currentNumber)
	code := &phoneCodeData{
		authKeyId:   authKeyId,
		phoneNumber: phoneNumber,
		state:       kCodeStateNone,
		dataType:    kDBTypeCreate,
	}

	code.codeCallBack = code;

	return code
}

func MakeCancelCodeData(authKeyId int64, phoneNumber, codeHash string) *phoneCodeData {
	code := &phoneCodeData{
		authKeyId:   authKeyId,
		codeHash:    codeHash,
		phoneNumber: phoneNumber,
		state:       kCodeStateNone,
		dataType:    kDBTypeDelete,
	}
	return code
}

func MakeCodeDataByHash(authKeyId int64, phoneNumber, codeHash string) *phoneCodeData {
	code := &phoneCodeData{
		authKeyId:   authKeyId,
		codeHash:    codeHash,
		phoneNumber: phoneNumber,
		state:       kCodeStateNone,
		dataType:    kDBTypeLoad,
	}
	return code
}

func (code *phoneCodeData) String() string {
	return fmt.Sprintf("{authKeyId: %d, phoneNumber: %s, codeHash: %s, state: %d}", code.authKeyId, code.phoneNumber, code.codeHash, code.state)
}

func (code *phoneCodeData) checkDataType(validType int) {
	// TODO(@benqi): panic
	if code.dataType != validType {
		glog.Fatal("invalid dataType")
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////
//func (code *phoneCodeData) fromDO(do *dataobject.AuthPhoneTransactionsDO) {
//	// TODO(@benqi): 111
//	do := &dataobject.AuthPhoneTransactionsDO{
//		ApiId:           apiId,
//		ApiHash:         apiHash,
//		PhoneNumber:     code.phoneNumber,
//		Code:            code.code,
//		CreatedAt:       time.Now().Format("2006-01-02 15:04:05"),
//		TransactionHash: code.codeHash,
//	}
//}
//
///////////////////////////////////////////////////////////////////////////////////////////////////////////
//func (code *phoneCodeData) toDO() *dataobject.AuthPhoneTransactionsDO {
//	return nil
//}

func (code *phoneCodeData) doSendCodeCallback() error {
	glog.Infof("doSendCodeCallback()")
	if code.codeCallback != nil {
	        glog.Infof("code.codeCallback defined. Sending Code...")
		return code.codeCallback.SendCode(code.code, code.codeHash, code.sentCodeType)
	}

	glog.Infof("code.codeCallback NOT defined. Returning dummy truth. Test Environment.")
	// TODO(@benqi): The test environment is sent successfully by default. 
	return nil
}

// auth.sendCode
func (code *phoneCodeData) DoSendCode(phoneRegistered, allowFlashCall, currentNumber bool, apiId int32, apiHash string) error {
	code.checkDataType(kDBTypeCreate)

	// Use the easiest way to create a new one each time. 
	sentCodeType, nextCodeType := makeCodeType(phoneRegistered, allowFlashCall, currentNumber)
	// TODO(@benqi): gen rand number
	code.code = "12345"
	// code.codeHash = fmt.Sprintf("%20d", helper.NextSnowflakeId())
	code.codeHash = crypto.GenerateStringNonce(16)
	code.codeExpired = int32(time.Now().Unix() + 15*60)
	code.sentCodeType = sentCodeType
	code.nextCodeType = nextCodeType

	err := code.doSendCodeCallback()
	if err != nil {
		glog.Error(err)
		return err
	}

	// save
	do := &dataobject.AuthPhoneTransactionsDO{
		AuthKeyId:        code.authKeyId,
		PhoneNumber:      code.phoneNumber,
		Code:             code.code,
		CodeExpired:      code.codeExpired,
		TransactionHash:  code.codeHash,
		SentCodeType:     int8(code.sentCodeType),
		FlashCallPattern: code.flashCallPattern,
		NextCodeType:     int8(code.nextCodeType),
		State: 			  kCodeStateSent,
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
	//	// TODO(@benqi): FLOOD_WAIT_X, too many attempts, please try later.
	//}

	return nil
}

// auth.resendCode
func (code *phoneCodeData) DoReSendCode() error {
	code.checkDataType(kDBTypeLoad)

	do := dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).SelectByPhoneCodeHash(code.authKeyId, code.phoneNumber, code.codeHash)
	if do == nil {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		glog.Error(err)
		return err
	}

	//// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}
	//
	//// TODO(@benqi): check phone code valid, only number etc.
	//if do.Code == "" {
	//	err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
	//	glog.Error(err)
	//	return err
	//}

	// check state invalid.
	if do.State != kCodeStateSent {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR), "code state error")
		glog.Error(err)
		return err
	}

	now := int32(time.Now().Unix())
	if now > do.CodeExpired {
		// TODO(@benqi): update timeout state?
		// dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).UpdateState(kCodeStateTimeout, do.Id)

		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EXPIRED), "code expired")
		glog.Error(err)
		return err
	}

	code.code = do.Code
	// TODO(@benqi): load from db
	code.codeExpired = do.CodeExpired
	code.sentCodeType = int(do.SentCodeType)
	code.flashCallPattern = do.FlashCallPattern
	code.nextCodeType = int(do.NextCodeType)
	code.state = int(do.State)
	code.tableId = do.Id

	err := code.doSendCodeCallback()
	if err != nil {
		glog.Error(err)
	}

	return err
}

// auth.cancelCode
func (code *phoneCodeData) DoCancelCode() bool {
	master := dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER)
	master.Delete(int8(kCodeStateDeleted), code.authKeyId, code.phoneNumber, code.codeHash)
	return true
}

func (code *phoneCodeData) DoSignIn(phoneCode string, phoneRegistered bool) error {
	defer func() {
		if code.tableId != 0 {
			// Update attempts
			dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).UpdateAttempts(code.tableId)
		}
	}()

	code.checkDataType(kDBTypeLoad)

	do := dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).SelectByPhoneCodeHash(code.authKeyId, code.phoneNumber, code.codeHash)
	if do == nil {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		glog.Error(code, ", error: ", err)
		return err
	}
	code.tableId = do.Id

	//// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}

	// TODO(@benqi): 重复请求处理...
	// check state invalid.
	if do.State != kCodeStateSent && do.State != kCodeStateSignIn {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR), "code state error")
		glog.Error(err)
		return err
	}

	now := int32(time.Now().Unix())
	if now > do.CodeExpired {
		// TODO(@benqi): update timeout state?
		// dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).UpdateState(kCodeStateTimeout, do.Id)

		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EXPIRED), "code expired")
		glog.Error(err)
		return err
	}

	// TODO(@benqi): check phone code valid, only number etc.
	if do.Code != phoneCode {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
		glog.Error(err)
		return err
	}

	// code.code = do.Code
	// TODO(@benqi): load from db
	// code.codeExpired = do.CodeExpired
	// code.sentCodeType = int(do.SentCodeType)
	// code.flashCallPattern = do.FlashCallPattern
	// code.nextCodeType = int(do.NextCodeType)

	// code.state = kCodeStateSignIn
	if phoneRegistered {
		code.state = kCodeStateOk
	} else {
		code.state = kCodeStateSignIn
	}

	// update state
	dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).UpdateState(int8(code.state), code.tableId)
	return nil
}

// TODO(@benqi): 合并DoSignUp和DoSignIn部分代码
func (code *phoneCodeData) DoSignUp(phoneCode string) error {
	defer func() {
		if code.tableId != 0 {
			// Update attempts
			dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).UpdateAttempts(code.tableId)
		}
	}()

	code.checkDataType(kDBTypeLoad)

	do := dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).SelectByPhoneCodeHash(code.authKeyId, code.phoneNumber, code.codeHash)
	if do == nil {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_NUMBER_INVALID), "invalid phone number")
		glog.Error(err)
		return err
	}
	code.tableId = do.Id

	// TODO(@benqi): attempts
	//if do.Attempts > 3 {
	//	// TODO(@benqi): 输入了太多次错误的phone code
	//	err := mtproto.NewFloodWaitX(15*60, "too many attempts.")
	//	return err
	//}

	// TODO(@benqi): 重复请求处理...
	// check state invalid.
	// TODO(@benqi): remote client error, state is Ok
	if do.State != kCodeStateSignIn && do.State != kCodeStateDeleted {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR), "code state error")
		glog.Error(err)
		return err
	}

	now := int32(time.Now().Unix())
	if now > do.CodeExpired {
		// TODO(@benqi): update timeout state?
		// dao.GetAuthPhoneTransactionsDAO(dao.DB_MASTER).UpdateState(kCodeStateTimeout, do.Id)

		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_EXPIRED), "code expired")
		glog.Error(err)
		return err
	}

	// TODO(@benqi): check phone code valid, only number etc.
	if do.Code != phoneCode {
		err := mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_PHONE_CODE_INVALID), "code invalid")
		glog.Error(err)
		return err
	}

	code.state = kCodeStateOk

	// update state
	dao.GetAuthPhoneTransactionsDAO(dao.DB_SLAVE).UpdateState(int8(code.state), code.tableId)
	return nil
}

// If the mobile number is already registered, check if other devices are online, and use the sendCodeTypeApp 
// Otherwise use sentCodeTypeSms 
// TODO(@benqi): Is there a use of sentCodeTypeFlashCall and entCodeTypeCall? ? 
func (code *phoneCodeData) ToAuthSentCode(phoneRegistered bool) *mtproto.TLAuthSentCode {
	// TODO(@benqi): only use sms

	authSentCode := &mtproto.TLAuthSentCode{Data2: &mtproto.Auth_SentCode_Data{
		PhoneRegistered: phoneRegistered,
		Type:            makeAuthSentCodeType(code.sentCodeType, len(code.code), code.flashCallPattern),
		PhoneCodeHash:   code.codeHash,
		NextType:        makeAuthCodeType(code.nextCodeType),
		Timeout:         60, // TODO(@benqi): 默认60s
	}}
	return authSentCode
}
