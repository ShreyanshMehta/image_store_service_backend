package common

import "strings"

func HashName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
	name = strings.ToLower(name)
	return name
}

func ErrorMsg(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  false,
		"message": message,
	}
}

func SuccessMsg(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  true,
		"message": message,
	}
}
