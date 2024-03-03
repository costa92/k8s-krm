
## Tool Binaries
GOLANGCI_LINT := golangci-lint
KUBE_LINT := kube-linter

.PHONY: lint.run
lint.run: lint.ci lint.kubefiles lint.dockerfiles lint.charts ## Run all available linters.


.PHONY: lint.golangci-lint
lint.golangci-lint: tools.verify.golangci-lint ## Run golangci to lint source codes.
	@echo "===========> Run golangci to lint source codes"
	@$(GOLANGCI_LINT) run -c $(KRM_ROOT)/.golangci.yaml $(KRM_ROOT)/...
