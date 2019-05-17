package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-rae/rae"
)

func main() {
	//use default error type
	rae.LoadFromFile("./errors.yml")
	re, err := rae.Error("db_error")
	if err != nil {
		fmt.Println(err)
	}

	b, err := json.Marshal(re)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[default]rest api err: %+v\njson:%s\n", re, b)
	//use cutom error type
	rae.ErrorType(&IRestAPIError{})
	rae.LoadFromFile("./errors.yml")
	re2, err := rae.Error("db_error")
	if err != nil {
		fmt.Println(err)
	}

	b2, err := json.Marshal(re2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("rest api err: %+v\njson:%s\n", re2, b2)
}

//IRestAPIError custom RestAPIError type
type IRestAPIError struct {
	ErrCode string `json:"code,omitempty"`
	ErrMsg  string `json:"msg,omitempty"`
	ErrDesc string `json:"desc,omitempty"`
}

//Code get code for rae
func (ie IRestAPIError) Code() string {
	return ie.ErrCode
}

//Msg get error msg for rae
func (ie IRestAPIError) Msg() string {
	return ie.ErrMsg
}

//Desc get error description for rae
func (ie IRestAPIError) Desc() string {
	return ie.ErrDesc
}

//Set set error value for rae
func (ie *IRestAPIError) Set(code, msg, desc string) {
	ie.ErrCode = code
	ie.ErrMsg = msg
	ie.ErrDesc = desc
}
