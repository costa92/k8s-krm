IMAGE_NAME = "costa92/krm-api:latest"
GIT_REVISION = $(shell git show -s --pretty=format:%h)

dev/dockerfile:
	docker build -t $(IMAGE_NAME) -f cmd/krm/Dockerfile .