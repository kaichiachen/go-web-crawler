package errs

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type ErrorType string

const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	ERROR_TYPE_ANALYZER   ErrorType = "analyzer error"
	ERROR_TYPE_PIPELINE   ErrorType = "pipeline error"
	ERROR_TYPE_SCHEDULER  ErrorType = "scheduler error"
)

type CrawlerError interface {
	Type() ErrorType
	Error() string
}

type myCrawlerError struct {
	errType    ErrorType
	errMsg     string
	fullErrMsg string
}

func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{
		errType: errType,
		errMsg:  strings.TrimSpace(errMsg),
	}
}

func NewCrawlerErrorBy(errType ErrorType, err error) CrawlerError {
	return NewCrawlerError(errType, err.Error())
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}

	return ce.fullErrMsg
}

func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}

	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}

type IllegalParameterError struct {
	msg string
}

// NewIllegalParameterError 会创建一个IllegalParameterError类型的实例。
func NewIllegalParameterError(errMsg string) error {
	// return IllegalParameterError{
	// 	msg: fmt.Sprintf("illegal parameter: %s",
	// 		strings.TrimSpace(errMsg)),
	// }
	return errors.New(fmt.Sprintf("illegal parameter: %s",
		strings.TrimSpace(errMsg)))
}
