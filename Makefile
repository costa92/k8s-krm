IMAGE_NAME = "costa92/krm-api:latest"
GIT_REVISION = $(shell git show -s --pretty=format:%h)
GO_MOD_DOMAIN = "github.com/costa92/krm-api"

vet:
	go vet ./...

fmt:
	@gofumpt -version || go install mvdan.cc/gofumpt@latest
	gofumpt -extra -w -d .
	@gci -v || go install github.com/daixiang0/gci@latest
	#gci write -s standard -s default -s 'Prefix($(GO_MOD_DOMAIN))' -s 'Prefix($(GO_MOD_NAME))' --skip-generated .
	gci write -s standard -s default -s 'Prefix($(GO_MOD_DOMAIN))' --skip-generated .

dev/dockerfile:
	docker build -t $(IMAGE_NAME) -f cmd/krm/Dockerfile .