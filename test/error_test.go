package test

import (
	"fmt"
	"github.com/1-bi/uerrors"
	"testing"
)

/**
 * case without argument
 */
func Test_Errorcode_Case1(t *testing.T) {

	errorCode := uerrors.NewCodeErrorWithPrefix("test", "test0001")

	errorCode.WithMsgBody("this is error message content.")

	//log.Printf()

	fmt.Println(errorCode.Error())

}

/**
 *  case 2 for with parameter define
 */
func Test_Errorcode_Build(t *testing.T) {

	errorCode := uerrors.NewCodeErrorWithPrefix("test", "test0001")

	errorCode.WithMsgBody("this is error message content with param.")
	errorCode.WithMsgBody("params: ${p1} , ${p2} , ${p3}.")

	//log.Printf()

	errorCode.Build("hello-message ", "my deal-other  ", "define")

	fmt.Println(errorCode.Error())

}

/**
 *  case 2 for with parameter define
 */
func Test_Errorcode_BuildMap(t *testing.T) {

	errorCode := uerrors.NewCodeErrorWithPrefix("test", "test0001")

	errorCode.WithMsgBody("this is error message content with param.")
	errorCode.WithMsgBody("params: ${p1} , ${p2} , ${p3}.")

	paramValue := make(map[string]interface{}, 0)
	paramValue["p1"] = 122345
	paramValue["p2"] = "hello message "
	paramValue["p3"] = "go ahead"

	errorCode.BuildByMap(paramValue)

	fmt.Println(errorCode.Error())

}
