package domain

import "time"

type ApiResponse struct {
	TraceID   string      `json:"trace_id"`
	Timestamp time.Time   `json:"timestamp"`
	Code      string      `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}
