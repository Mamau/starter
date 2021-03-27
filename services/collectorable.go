package services

import (
	"github.com/mamau/starter/config/docker"
	"github.com/mamau/starter/entity"
)

type Collectorable interface {
	GetDockerConfig() *docker.Docker
	GetCommandConfig() *entity.Command
	GetClientSignature(cmd []string) []string
	GetImage() string
}