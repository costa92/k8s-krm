
.PHONY: _install.golangci-lint
_install.golangci-lint:
	@echo "===========> Install golangci-lint $(GOLANGCI_LINT_VERSION)"
	@chmod +x $(SCRIPTS_DIR)/add-completion.sh
	@echo "===> execute: golangci-lint completion bash > golangci-lint-completion.bash"
	@$(SCRIPTS_DIR)/add-completion.sh golangci-lint bash


.PHONY: tools.verify.%
tools.verify.%:
	@echo "===========> Verify $*"
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi  # if not found, install it