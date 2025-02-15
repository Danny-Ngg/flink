package common

import (
	"github.com/filswan/go-swan-lib/constants"
)

type BasicResponse struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message,omitempty"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
}

type PageInfo struct {
	PageNumber       string `json:"page_number"`
	PageSize         string `json:"page_size"`
	TotalRecordCount string `json:"total_record_count"`
}

func CreateSuccessResponse(data interface{}) BasicResponse {
	return BasicResponse{
		Status: constants.SWAN_API_STATUS_SUCCESS,
		Data:   data,
	}
}

func CreateErrorResponse(errMsg string) BasicResponse {
	return BasicResponse{
		Status:  constants.SWAN_API_STATUS_FAIL,
		Message: errMsg,
	}
}
