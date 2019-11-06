package models

type RescueGetResponse struct {
	Rescue Rescue `json:"rescue"`
}

type AuthorizedKey struct {
	Key Key `json:"key"`
}

type Rescue struct {
	ServerIP      string          `json:"server_ip"`
	ServerNumber  int             `json:"server_number"`
	Os            interface{}     `json:"os"`   // unfortunately the API returns an array vs. a single value based on state (active/inactive)
	Arch          interface{}     `json:"arch"` // unfortunately the API returns an array vs. a single value based on state (active/inactive)
	Active        bool            `json:"active"`
	Password      string          `json:"password"`
	AuthorizedKey []AuthorizedKey `json:"authorized_key"`
	HostKey       []interface{}   `json:"host_key"`
}

type RescueSetInput struct {
	OS            string
	Arch          int
	AuthorizedKey string
}
