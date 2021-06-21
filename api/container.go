package api

import (
	"bytes"
	"bzz-collector/vars"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainersInspec() []*vars.ContainerStatusInfo {

	var ContainersInfo []*vars.ContainerStatusInfo

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return []*vars.ContainerStatusInfo{}
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return []*vars.ContainerStatusInfo{}
	}

	queryPortKey := "1635/tcp"

	ip, err := CheckIP()
	if err != nil {
		ip = "127.0.0.1"
	}

	for _, container := range containers {

		containerInspecJson, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			continue
		}

		ContainerStatus := containerInspecJson.State.Status
		ContainerName := containerInspecJson.Name

		ContainerNetworkData := containerInspecJson.NetworkSettings
		containerPortsMappers := ContainerNetworkData.Ports

		var ContainerPort string
		var PeersCount int
		var ChequeCount int

		for key, ports := range containerPortsMappers {
			if string(key) == queryPortKey {
				ContainerPort = ports[0].HostPort
				PeersCount = GetPeersCount(ContainerPort)
				ChequeCount = GetChequeCount(ContainerPort)
			}
		}

		currentContainerInfo := &vars.ContainerStatusInfo{
			Name:        ContainerName,
			Port:        ContainerPort,
			Status:      ContainerStatus,
			PeersCount:  PeersCount,
			ChequeCount: ChequeCount,
			IpAddress:   ip,
		}

		ContainersInfo = append(ContainersInfo, currentContainerInfo)

	}
	return ContainersInfo
}

func GetChequeCount(port string) int {
	url := fmt.Sprintf("http://localhost:%s/chequebook/cheque", port)
	resp, err := http.Get(url)
	if err != nil {
		return -1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}

	var chequeData map[string][]interface{}
	json.Unmarshal(body, &chequeData)

	return len(chequeData["lastcheques"])
}

func GetPeersCount(port string) int {
	url := fmt.Sprintf("http://localhost:%s/peers", port)
	resp, err := http.Get(url)
	if err != nil {
		return -1
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}

	var peersData map[string][]interface{}
	json.Unmarshal(body, &peersData)

	return len(peersData["peers"])
}

func CheckIP() (string, error) {
	rsp, err := http.Get("http://checkip.amazonaws.com")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(buf)), nil
}
