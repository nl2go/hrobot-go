package client

import "github.com/nl2go/hrobot-go/models"

type RobotClient interface {
	SetBaseURL(baseURL string)
	SetUserAgent(userAgent string)
	GetVersion() string
	ServerGetList() ([]models.Server, error)
	ServerGet(ip string) (*models.Server, error)
	ServerSetName(ip string, input *models.ServerSetNameInput) (*models.Server, error)
	ServerReverse(ip string) (*models.Cancellation, error)
	KeyGetList() ([]models.Key, error)
	IPGetList() ([]models.IP, error)
	RDnsGetList() ([]models.Rdns, error)
	RDnsGet(ip string) (*models.Rdns, error)
	BootRescueGet(ip string) (*models.Rescue, error)
	BootRescueSet(ip string, input *models.RescueSetInput) (*models.Rescue, error)
	ResetGet(ip string) (*models.Reset, error)
	ResetSet(ip string, input *models.ResetSetInput) (*models.ResetPost, error)
	FailoverGetList() ([]models.Failover, error)
	FailoverGet(ip string) (*models.Failover, error)
}
