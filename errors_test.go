package rae

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	if err := LoadFromFile("./examples/errors.yml"); err == nil {
		if er, err := restErrorList.get("db_error"); err == nil {
			b, err := json.Marshal(er)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("json:%s\n", b)

		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Print(err)
	}

}
func TestError(t *testing.T) {
	re, err := Error("db_error")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("err:%+v\n", re)
}
