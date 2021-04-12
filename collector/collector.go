package collector

import (
	"github.com/mamau/starter/libs"
)

type Collector struct {
	entity Collectorable
}

func NewCollector(c Collectorable) *Collector {
	return &Collector{
		entity: c,
	}
}

func (c *Collector) DockerConfigCommand() []string {
	return libs.MergeSliceOfString([]string{
		c.entity.GetDockerConfig().GetUserId(),
		c.entity.GetDockerConfig().GetEnvironmentVariables(),
		c.entity.GetDockerConfig().GetHosts(),
		c.entity.GetDockerConfig().GetPorts(),
		c.entity.GetDockerConfig().GetDns(),
		c.entity.GetWorkDir(),
		c.entity.GetDockerConfig().GetVolumes(),
		c.entity.GetProjectVolume(),
		c.entity.GetImage(),
	})
}

func (c *Collector) ClientCommand() []string {
	execCommand := c.entity.GetImageCommand()

	preCommand := c.entity.GetDockerConfig().GetPreCommands()
	if len(preCommand) > 0 {
		preCommand += ";"
	}

	clientCommand := c.entity.GetClientCommand()
	postCommand := c.entity.GetDockerConfig().GetPostCommands()
	if len(postCommand) > 0 {
		clientCommand += ";"
	}

	fullCommand := []string{
		execCommand,
		preCommand,
		clientCommand,
		postCommand,
	}

	return libs.DeleteEmpty(libs.MergeSliceOfString(fullCommand))
}

func (c *Collector) CollectCommand() []string {
	bc := c.GetBeginCommand()
	bc = append(bc, c.DockerConfigCommand()...)
	bc = append(bc, c.entity.GetClientSignature(c.ClientCommand())...)
	return bc
}

func (c *Collector) GetBeginCommand() []string {
	var bc []string

	bc = append(bc, c.entity.GetDockerConfig().GetDockerCommand())
	bc = append(bc, c.entity.GetDockerConfig().GetFlags())
	bc = append(bc, c.entity.GetDockerConfig().GetDetached())
	bc = append(bc, c.entity.GetDockerConfig().GetCleanUp())

	return libs.DeleteEmpty(bc)
}
