package mocks

// this generates a mock for the DockerClient interface
//go:generate mockgen -destination=docker/dockerClient.go -package=mocks go.arcalot.io/imagebuilder/internal/docker DockerClient

// generate a mock for the ContainerEngineService interface
//go:generate mockgen -destination=ce_service/ce_service.go -package=mocks go.arcalot.io/imagebuilder/internal/ce_service ContainerEngineService
