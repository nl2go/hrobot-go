package models

type CancellationResponse struct {
	Cancellation Cancellation `json:"cancellation"`
}
type Cancellation struct {
	ServerIP                 string      `json:"server_ip"`
	ServerNumber             int         `json:"server_number"`
	ServerName               string      `json:"server_name"`
	EarliestCancellationDate string      `json:"earliest_cancellation_date"`
	Cancelled                bool        `json:"cancelled"`
	CancellationDate         string      `json:"cancellation_date"`
	CancellationReason       interface{} `json:"cancellation_reason"`
}
