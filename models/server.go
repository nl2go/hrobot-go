package models

type ServerResponse struct {
	Server Server `json:"server"`
}

type Subnet struct {
	IP   string `json:"ip"`
	Mask string `json:"mask"`
}

type Server struct {
	ServerIP     string   `json:"server_ip"`
	ServerNumber int      `json:"server_number"`
	ServerName   string   `json:"server_name"`
	Product      string   `json:"product"`
	Dc           string   `json:"dc"`
	Traffic      string   `json:"traffic"`
	Flatrate     bool     `json:"flatrate"`
	Status       string   `json:"status"`
	Throttled    bool     `json:"throttled"`
	Cancelled    bool     `json:"cancelled"`
	PaidUntil    string   `json:"paid_until"`
	IP           []string `json:"ip"`
	Subnet       []Subnet `json:"subnet"`
	Reset        bool     `json:"reset"`
	Rescue       bool     `json:"rescue"`
	Vnc          bool     `json:"vnc"`
	Windows      bool     `json:"windows"`
	Plesk        bool     `json:"plesk"`
	Cpanel       bool     `json:"cpanel"`
	Wol          bool     `json:"wol"`
	HotSwap      bool     `json:"hot_swap"`
}

type ServerSetNameInput struct {
	Name string
}
