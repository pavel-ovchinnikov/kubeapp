# Variables
APP_NAME := kubeapp1
SRC_DIR := src
DOCKER_IMAGE := $(APP_NAME):latest

# Docker build target
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Docker run target
.PHONY: docker-build docker-run
docker-run:
	@echo "Running Docker container..."
	docker run --rm -it $(DOCKER_IMAGE)
