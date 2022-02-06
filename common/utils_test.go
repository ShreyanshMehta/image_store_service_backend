package common

import (
	"reflect"
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Testing current formatted, time",
			want: time.Now().Format("01-02-2006 15:04:05 Monday"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetCurrentTime(); got != tt.want {
				t.Errorf("GetCurrentTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Testing for a single word name",
			args: args{"Testcase"},
			want: "testcase",
		},
		{
			name: "Testing for a multiple words",
			args: args{"Multiple Word Test Cases"},
			want: "multiplewordtestcases",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashName(tt.args.name); got != tt.want {
				t.Errorf("HashName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Error(t *testing.T) {
	type fields struct {
		Message   string
		ErrorCode int
		Status    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "When status true is passed in constructor for error message",
			fields: fields{
				Message:   "This is a message",
				ErrorCode: 34222,
				Status:    true,
			},
			want: map[string]interface{}{
				"message":    "This is a message",
				"error_code": 34222,
				"status":     false,
			},
		},
		{
			name: "When status false is passed in constructor for error message",
			fields: fields{
				Message:   "This is a message",
				ErrorCode: 34222,
				Status:    false,
			},
			want: map[string]interface{}{
				"message":    "This is a message",
				"error_code": 34222,
				"status":     false,
			},
		},
		{
			name: "When error code is not passed in constructor",
			fields: fields{
				Message: "This is a message",
				Status:  false,
			},
			want: map[string]interface{}{
				"message": "This is a message",
				"status":  false,
			},
		},
		{
			name: "When message is not passed in constructor",
			fields: fields{
				Status: false,
			},
			want: map[string]interface{}{
				"status": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := Response{
				Message:   tt.fields.Message,
				ErrorCode: tt.fields.ErrorCode,
				Status:    tt.fields.Status,
			}
			if got := resp.Error(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_Success(t *testing.T) {
	type fields struct {
		Message   string
		ErrorCode int
		Status    bool
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]interface{}
	}{
		{
			name: "When status true is passed in constructor for error message",
			fields: fields{
				Message: "This is a message",
				Status:  true,
			},
			args: args{},
			want: map[string]interface{}{
				"message": "This is a message",
				"status":  true,
			},
		},
		{
			name: "When status false is passed in constructor for error message",
			fields: fields{
				Message: "This is a message",
				Status:  false,
			},
			args: args{},
			want: map[string]interface{}{
				"message": "This is a message",
				"status":  true,
			},
		},
		{
			name: "When object is passed",
			fields: fields{
				Message: "This is a message",
				Status:  true,
			},
			args: args{
				map[string]string{
					"key": "value",
				},
			},
			want: map[string]interface{}{
				"message": "This is a message",
				"status":  true,
				"data": map[string]string{
					"key": "value",
				},
			},
		},
		{
			name: "When list is passed",
			fields: fields{
				Message: "This is a message",
				Status:  true,
			},
			args: args{
				[]string{"Item1", "Item2", "Item3"},
			},
			want: map[string]interface{}{
				"message": "This is a message",
				"status":  true,
				"data":    []string{"Item1", "Item2", "Item3"},
			},
		},
		{
			name: "When message is not passed in constructor",
			fields: fields{
				Status: true,
			},
			want: map[string]interface{}{
				"status": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := Response{
				Message:   tt.fields.Message,
				ErrorCode: tt.fields.ErrorCode,
				Status:    tt.fields.Status,
			}
			if got := resp.Success(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Success() = %v, want %v", got, tt.want)
			}
		})
	}
}
