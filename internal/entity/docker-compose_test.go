package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDCToCommand(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.ToCommand([]string{}))

	dc.Path = "/some/path"
	dc.Verbose = true

	result := strings.Join(dc.ToCommand([]string{"up"}), " ")
	e := "--file /some/path --verbose up"
	assert.Equal(t, result, e)

	dc.RemoveOrphans = true
	result = strings.Join(dc.ToCommand([]string{"up"}), " ")
	e = "--file /some/path --remove-orphans --verbose up"
	assert.Equal(t, result, e)

	dc.EnvFile = "/some/path/to/.env"
	result = strings.Join(dc.ToCommand([]string{"up"}), " ")
	e = "--file /some/path --env-file=/some/path/to/.env --remove-orphans --verbose up"
	assert.Equal(t, result, e)

	dc.BuildArgs = []string{"VAR1=123", "VAR2=45"}
	result = strings.Join(dc.ToCommand([]string{"up"}), " ")
	e = "--file /some/path --env-file=/some/path/to/.env --remove-orphans --build-arg VAR1=123 --build-arg VAR2=45 --verbose up"
	assert.Equal(t, result, e)

	dc.User = "1000"
	result = strings.Join(dc.ToCommand([]string{"up"}), " ")
	e = "--file /some/path -u 1000 --env-file=/some/path/to/.env --remove-orphans --build-arg VAR1=123 --build-arg VAR2=45 --verbose up"
	assert.Equal(t, result, e)

	dc.Detach = true
	result = strings.Join(dc.ToCommand([]string{"up"}), " ")
	e = "--file /some/path -u 1000 -d --env-file=/some/path/to/.env --remove-orphans --build-arg VAR1=123 --build-arg VAR2=45 --verbose up"
	assert.Equal(t, result, e)
}

func TestGetProjectDirectory(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetProjectDirectory())

	dc.ProjectDirectory = "/some/dir"
	assert.Equal(t, dc.GetProjectDirectory(), "--project-directory /some/dir")
}

func TestGetPath(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetPath())

	dc.Path = "/some/path"
	assert.Equal(t, dc.GetPath(), "--file /some/path")
}

func TestGetProjectName(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetProjectName())

	dc.ProjectName = "some-p-name"
	assert.Equal(t, dc.GetProjectName(), "--project-name some-p-name")
}

func TestGetLogLevel(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetLogLevel())

	dc.LogLevel = "DEBUG"
	assert.Equal(t, dc.GetLogLevel(), "--log-level DEBUG")
}

func TestGetVerbose(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetVerbose())

	dc.Verbose = true
	assert.Equal(t, dc.GetVerbose(), "--verbose")
}

func TestDCGetName(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetName())

	dc.Name = "some-name"
	assert.Equal(t, dc.GetName(), "some-name")
}

func TestDCGetDescription(t *testing.T) {
	dc := DockerCompose{}
	assert.Empty(t, dc.GetDescription())

	dc.Description = "some description"
	assert.Equal(t, dc.GetDescription(), "some description")
}

func TestDCGetExecCommand(t *testing.T) {
	dc := DockerCompose{}
	assert.Equal(t, dc.GetExecCommand(), string(DOCKER_COMPOSE))
}
