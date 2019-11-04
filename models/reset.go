package models

const ResetTypePower = "power"
const ResetTypeHardware = "hw"
const ResetTypeManual = "man"

type ResetResponse struct {
	Reset Reset `json:"reset"`
}

type Reset struct {
	OperatingStatus string   `json:"operating_status"`
	ServerIP        string   `json:"server_ip"`
	ServerNumber    int      `json:"server_number"`
	Type            []string `json:"type"`
}

type ResetPostResponse struct {
	Reset ResetPost `json:"reset"`
}

type ResetPost struct {
	ServerIP string `json:"server_ip"`
	Type     string `json:"type"`
}

type ResetSetInput struct {
	Type string
}
