package rae

import (
	"errors"
	"sync"
)

//
type inputError struct {
	Code string
	Key  string
	Msg  string
	Desc string
}
type inputErrorConfig struct {
	App     string
	Version string
	Errors  []inputError `yaml:",flow"`
}

//RestAPIError restful api error
type RestAPIError interface {
	Code() string
	Msg() string
	Desc() string
	Set(code, msg, desc string)
}

//restErrors restful api error
type restErrors struct {
	sync.Map
}

var restErrorList *restErrors

//Get error by key
func (re *restErrors) get(key string) (RestAPIError, error) {
	v, ok := re.Load(key)
	if ok {
		return v.(RestAPIError), nil
	}

	return nil, errors.New("error not found")

}

//AddError add error to restErrors
func (re *restErrors) add(key string, err RestAPIError) error {
	if key == "" {
		return errors.New("key must be not nil")
	}
	//check
	re.Store(key, err)
	return nil
}

//AddError add error to restErrors
// func (re *restErrors) delete(key string) error {
// 	if key == "" {
// 		return errors.New("key must be not nil")
// 	}
// 	re.Delete(key)
// 	return nil
// }

//Error get rest api error by key
func Error(key string) (RestAPIError, error) {
	if re, err := restErrorList.get(key); err == nil {
		return re, nil
	}
	return nil, errors.New("error not found")
}

var restAPIErrorType RestAPIError

func init() {
	restAPIErrorType = &DefaultRestAPIError{}
}

// ErrorType set error type
func ErrorType(raeType RestAPIError) {
	restAPIErrorType = raeType
}
