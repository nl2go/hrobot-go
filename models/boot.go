package models

type RescueGetResponse struct {
	Rescue RescueOptions `json:"rescue"`
}

type RescueOptions struct {
	ServerIP      string        `json:"server_ip"`
	ServerNumber  int           `json:"server_number"`
	Os            []string      `json:"os"`
	Arch          []int         `json:"arch"`
	Active        bool          `json:"active"`
	Password      interface{}   `json:"password"`
	AuthorizedKey []interface{} `json:"authorized_key"`
	HostKey       []interface{} `json:"host_key"`
}

type RescueSetInput struct {
	OS            string
	Arch          int
	AuthorizedKey string
}

type RescueSetResponse struct {
	Rescue RescueValues `json:"rescue"`
}

type RescueValues struct {
	ServerIP      string        `json:"server_ip"`
	ServerNumber  int           `json:"server_number"`
	Os            string        `json:"os"`
	Arch          int           `json:"arch"`
	Active        bool          `json:"active"`
	Password      string        `json:"password"`
	AuthorizedKey []interface{} `json:"authorized_key"`
	HostKey       []interface{} `json:"host_key"`
}
