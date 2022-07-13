package main

import (
	"encoding/json"
	"fmt"
)

type ResumeSendStatementRequest struct {
	SendStatementID int64 `json:"sendStatementID,omitempty"`
	Limit int64 `json:"limit,omitempty"`
}

func main() {
	var result ResumeSendStatementRequest

	str := "{\"sendStatementID\": 31}"
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		fmt.Println("unjson 失败")
	}
	fmt.Println("unjson after： ",result)
}