IMAGE_NAME = "costa92/krm:latest"
GIT_REVISION = $(shell git show -s --pretty=format:%h)
GO_MOD_DOMAIN = "github.com/costa92/krm"

vet:
	go vet ./...

fmt:
	@gofumpt -version || go install mvdan.cc/gofumpt@latest
	gofumpt -extra -w -d .
	@gci -v || go install github.com/daixiang0/gci@latest
	#gci write -s standard -s default -s 'Prefix($(GO_MOD_DOMAIN))' -s 'Prefix($(GO_MOD_NAME))' --skip-generated .
	gci write -s standard -s default -s 'Prefix($(GO_MOD_DOMAIN))' --skip-generated .



.PHONY: install
install:
	go version
	go install mvdan.cc/gofumpt@latest
	gofumpt --version
	go install github.com/daixiang0/gci@latest
	gci --version
	go install golang.org/x/tools/gopls@latest
	gopls version
	# https://golangci-lint.run/usage/install/#local-installation
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.54.2
	which golangci-lint
	golangci-lint --version

lint:
	golangci-lint version
	golangci-lint run -v --color always --out-format colored-line-number

dev/dockerfile:
	docker build -t $(IMAGE_NAME) -f cmd/krm/Dockerfile .