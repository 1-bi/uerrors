package uerrors

const PREFIX = "uerrors"

type CodeError interface {
	error
	Code() string
	Prefix() string
	MsgBody() string

	WithMsgBody(content string) CodeError

	/**
	 * build err object
	 */
	Build(args ...string) CodeError

	/**
	 * build err object
	 */
	BuildByMap(args map[string]interface{}) CodeError
}

type ValidatorError interface {
	CodeError
}
