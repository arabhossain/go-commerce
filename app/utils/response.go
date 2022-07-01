package utils

import (
	"encoding/json"
	"net/http"
)
type DefaultStructure struct {
	Status int `json:"status"`
	Data interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

func JsonResponse(writer http.ResponseWriter, response interface{}, status int)  {
	writer.Header().Set("Content-Type", "application/json")
	rs := DefaultStructure{
		Status: status,
		Data:   response,
	}
	if status != 200 {
		rs.Errors = response
		rs.Data = []string{}
	}

	json.NewEncoder(writer).Encode(rs)
}
