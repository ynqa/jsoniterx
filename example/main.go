package main

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/ynqa/jsoniterx"
)

type Example struct {
	Time time.Time `json:"time" format:"2006-01-02 15:04:05" location:"Asia/Tokyo"`
}

func main() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	json.RegisterExtension(jsoniterx.TimePlugin())

	var e Example
	str := `{"time": "2019-01-01 12:00:00"}`
	if err := json.Unmarshal([]byte(str), &e); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e.Time)
}
