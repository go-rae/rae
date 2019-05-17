package rae

import (
	"errors"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"
)

// LoadFromFile load errors from yaml file
func LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.New("open error config yaml file error")
	}
	defer file.Close()
	var iErrs inputErrorConfig
	if err := yaml.NewDecoder(file).Decode(&iErrs); err == nil {
		if len(iErrs.Errors) > 0 {
			restErrorList = new(restErrors)
			for _, iErr := range iErrs.Errors {
				var rae = restAPIErrorType
				//clone a new error convert *RestAPIError to RestAPIError
				raeType := reflect.TypeOf(rae).Elem()
				//rae pointer
				newRae := reflect.New(raeType).Interface().(RestAPIError)
				newRae.Set(iErr.Code, iErr.Msg, iErr.Desc)
				restErrorList.add(iErr.Key, newRae)
			}
			return nil
		}

	} else {
		return errors.New("read error from yaml failed")
	}
	return errors.New("unkown error")
}
