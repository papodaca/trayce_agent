package docker

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/docker/docker/client"
	"github.com/evanrolfe/dockerdog/api"
)

type Containers struct {
	containerIds []string
	dockerClient *client.Client
	filterCmd    string
}

func NewContainers(filterCmd string) *Containers {
	dockerC, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	return &Containers{
		containerIds: []string{},
		dockerClient: dockerC,
		filterCmd:    filterCmd,
	}
}

func (c *Containers) GetPidsToIntercept() []int {
	pids := []int{}

	for _, containerId := range c.containerIds {
		pids = append(pids, c.getPidsForContainer(containerId)...)
	}

	return pids
}

func (c *Containers) GetIdFromPid(pid int) string {
	return ""
}

func (c *Containers) SetSettings(settings *api.Settings) {
	c.containerIds = settings.ContainerIds
}

func (c *Containers) getPidsForContainer(containerId string) []int {
	result, err := c.dockerClient.ContainerTop(context.Background(), containerId, []string{})
	if err != nil {
		fmt.Println("[Containers] ContainerTop() error:", err)
	}

	// Extract the index of the PID in the results
	indexPid := -1
	for i, title := range result.Titles {
		if title == "PID" {
			indexPid = i
		}
	}
	if indexPid == -1 {
		panic("no index found for PID from docker.ContainerTop()")
	}

	// Extract the index of the Command in the results
	indexCmd := -1
	for i, title := range result.Titles {
		if title == "CMD" {
			indexCmd = i
		}
	}
	if indexCmd == -1 {
		panic("no index found for CMD from docker.ContainerTop()")
	}

	// Collect the PIDs
	pids := []int{}
	for i := 0; i < len(result.Processes); i++ {
		cmd := result.Processes[i][indexCmd]
		pidStr := result.Processes[i][indexPid]
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			panic(err)
		}

		if c.filterCmd == "" || strings.Contains(cmd, c.filterCmd) {
			pids = append(pids, pid)
		}
	}

	return pids
}
