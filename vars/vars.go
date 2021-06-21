package vars

import "fmt"

type ContainerStatusInfo struct {
	Name        string
	Port        string
	Status      string
	PeersCount  int
	ChequeCount int
	IpAddress   string
}

func (c *ContainerStatusInfo) String() string {
	msg := fmt.Sprintf("Name: %s Port: %s Status: %s PeersCount: %d ChequeCount: %d IpAddress: %s\n", c.Name, c.Port, c.Status, c.PeersCount, c.ChequeCount, c.IpAddress)
	return msg
}
