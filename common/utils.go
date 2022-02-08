package common

import (
	"strings"
	"time"
)

type Response struct {
	Message   string
	ErrorCode int
	Status    bool
}

func HashName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ToLower(name)
	return name
}

func (resp Response) Error() map[string]interface{} {
	r := make(map[string]interface{})
	r["status"] = false
	if resp.Message != "" {
		r["message"] = resp.Message
	}
	if resp.ErrorCode != 0 {
		r["error_code"] = resp.ErrorCode
	}
	return r
}

func (resp Response) Success(data interface{}) map[string]interface{} {
	r := make(map[string]interface{})
	r["status"] = true
	if resp.Message != "" {
		r["message"] = resp.Message
	}
	if data != nil {
		r["data"] = data
	}
	return r
}

func GetCurrentTime() string {
	currentTime := time.Now()
	currentTimeFormatted := currentTime.Format("01-02-2006 15:04:05")
	return currentTimeFormatted
}
