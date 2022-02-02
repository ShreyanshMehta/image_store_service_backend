package common

import "strings"

type Response struct {
	Message   string
	Data      []interface{}
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
	if len(resp.Data) > 0 {
		r["data"] = resp.Data
	}
	return r
}

func (resp Response) Success() map[string]interface{} {
	r := make(map[string]interface{})
	r["status"] = true
	if resp.Message != "" {
		r["message"] = resp.Message
	}
	if resp.ErrorCode != 0 {
		r["error_code"] = resp.ErrorCode
	}
	if len(resp.Data) > 0 {
		r["data"] = resp.Data
	}
	return r
}
