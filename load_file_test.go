package rae

import (
	"encoding/json"
	"fmt"
	"testing"
)

type IRestAPIError struct {
	ErrCode string `json:"err_code,omitempty"`
	ErrMsg  string `json:"err_msg,omitempty"`
	ErrDesc string `json:"err_desc,omitempty"`
}

func (ie IRestAPIError) Code() string {
	return ie.ErrCode
}
func (ie IRestAPIError) Msg() string {
	return ie.ErrMsg
}
func (ie IRestAPIError) Desc() string {
	return ie.ErrDesc
}
func (ie *IRestAPIError) Set(code, msg, desc string) {
	ie.ErrCode = code
	ie.ErrMsg = msg
	ie.ErrDesc = desc
}
func TestLoadFromFile(t *testing.T) {
	if err := LoadFromFile("./examples/errors.yml"); err == nil {
		if er, err := restErrorList.get("db_error"); err == nil {
			b, err := json.Marshal(er)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("json:%s", b)

		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Print(err)
	}

}
