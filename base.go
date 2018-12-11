package uerrors

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/**
 * build code error
 */
type baseCodeError struct {
	code    string
	msgBody string
	prefix  string
}

/**
 * implment error function
 */
func (this *baseCodeError) Error() string {

	// --- check message content ----
	buf := bytes.NewBufferString("[")

	if this.prefix != "" {
		buf.WriteString(this.prefix)
		buf.WriteString(":")
	}

	buf.WriteString(this.code)
	buf.WriteString("]: ")

	buf.WriteString(this.msgBody)
	buf.WriteString("\n ")

	return fmt.Sprintf(buf.String())
}

func (this *baseCodeError) Code() string {
	return this.code
}

func (this *baseCodeError) Prefix() string {
	return this.prefix
}

func (this *baseCodeError) MsgBody() string {
	return this.msgBody
}

/**
 * found get params map
 */
func (this *baseCodeError) getParamsInContent(content string) []string {

	paramSlice := make([]string, 0)

	tmpstring := content
	tmpendInd := len(content)

	var ind, endSymbolInd int
	var matchstring, argmentName string

	for {

		ind = strings.LastIndex(tmpstring, "${")
		if ind == -1 {
			break
		}
		matchstring = tmpstring[ind:tmpendInd]
		endSymbolInd = strings.Index(matchstring, "}")

		argmentName = matchstring[:endSymbolInd+1]

		tmpstring = tmpstring[0:ind]
		tmpendInd = len(tmpstring)

		paramSlice = append(paramSlice, argmentName)
	}

	return paramSlice

}

func (this *baseCodeError) reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func (this *baseCodeError) Build(args ...string) CodeError {

	// --- check the content argument length ---
	paramsInContent := this.getParamsInContent(this.msgBody)
	lenParamsVal := len(args)
	lenParams := len(paramsInContent)

	fmt.Println(this.msgBody)
	fmt.Println(paramsInContent)

	if lenParams != lenParamsVal {
		msg := "The quantity of value inputed  doesn't  match name defined in message. The param names is " + strconv.Itoa(lenParams) + ", and param values is " + strconv.Itoa(lenParamsVal)
		return ErrParamsNotMatch.WithMsgBody(msg)
	}

	this.reverse(paramsInContent)

	tmpstring := this.msgBody
	var val interface{}
	for ind, paramKey := range paramsInContent {
		val = args[ind]
		if val != nil {
			// --- convert string -----
			tmpstring = strings.Replace(tmpstring, paramKey, this.convertToString(val), -1)
		}
	}
	this.msgBody = tmpstring

	return this
}

func (this *baseCodeError) convertToString(value interface{}) string {
	kind := reflect.TypeOf(value).Kind()

	if kind == reflect.Int {
		return strconv.Itoa(value.(int))

	} else if kind == reflect.Int8 {
		result := strconv.FormatInt(int64(value.(int8)), 10)
		return result
	} else if kind == reflect.Int16 {
		result := strconv.FormatInt(int64(value.(int16)), 10)
		return result
	} else if kind == reflect.Int32 {
		result := strconv.FormatInt(int64(value.(int32)), 10)
		return result
	} else if kind == reflect.Int64 {
		result := strconv.FormatInt(value.(int64), 10)
		return result
	} else if kind == reflect.Uint {
		result := strconv.FormatUint(uint64(value.(uint)), 10)
		return result
	} else if kind == reflect.Uint8 {
		result := strconv.FormatUint(uint64(value.(uint8)), 10)
		return result
	} else if kind == reflect.Uint16 {
		result := strconv.FormatUint(uint64(value.(uint16)), 10)
		return result
	} else if kind == reflect.Uint32 {
		result := strconv.FormatUint(uint64(value.(uint32)), 10)
		return result
	} else if kind == reflect.Uint64 {
		result := strconv.FormatUint(value.(uint64), 10)
		return result
	} else if kind == reflect.Float32 {
		result := strconv.FormatFloat(float64(value.(float32)), 'f', 6, 64)
		return result
	} else if kind == reflect.Float64 {
		result := strconv.FormatFloat(value.(float64), 'f', 6, 64)
		return result
	} else if kind == reflect.Bool {
		result := strconv.FormatBool(value.(bool))
		return result
	} else {
		return value.(string)
	}

}

func (this *baseCodeError) BuildByMap(args map[string]interface{}) CodeError {

	// --- check the content argument length ---
	paramsInContent := this.getParamsInContent(this.msgBody)
	lenParamsVal := len(args)
	lenParams := len(paramsInContent)

	if lenParams != lenParamsVal {
		msg := "The quantity of value inputed  doesn't  match name defined in message. The param names is " + strconv.Itoa(lenParams) + ", and param values is " + strconv.Itoa(lenParamsVal)
		return ErrParamsNotMatch.WithMsgBody(msg)
	}

	// --- replace param name with value ---
	tmpstring := this.msgBody
	var val interface{}
	for _, paramKey := range paramsInContent {
		val = args[paramKey[2:len(paramKey)-1]]
		if val != nil {
			// --- convert string -----
			tmpstring = strings.Replace(tmpstring, paramKey, this.convertToString(val), -1)
		}
	}
	this.msgBody = tmpstring

	return this
}

func (this *baseCodeError) WithMsgBody(content string) CodeError {
	buf := bytes.NewBufferString(this.msgBody)
	buf.WriteString(content)

	this.msgBody = buf.String()

	return this
}

/**
 *   create public method code error , use simple code handle --
 */
func NewCodeError(code string, msgContent ...string) CodeError {

	ce := baseCodeError{code, "", ""}

	// --- check message content ----
	buf := bytes.NewBufferString("")

	for _, msg := range msgContent {
		buf.WriteString(msg)
	}

	ce.msgBody = buf.String()

	return &ce
}

/**
 * create new code error
 */
func NewCodeErrorWithPrefix(prefix string, code string, msgContent ...string) CodeError {

	ce := baseCodeError{code, "", prefix}

	// --- check message content ----
	buf := bytes.NewBufferString("")

	for _, msg := range msgContent {
		buf.WriteString(msg)
	}

	ce.msgBody = buf.String()

	return &ce
}
