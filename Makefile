NAME=goinsta

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


TARGET_MAX_CHAR_NUM=20


define colored
	@echo '${GREEN}$1${RESET}'
endef

## Show help
help:
	${call colored, help is running...}
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)


## Format code.
fmt:
	${call colored, fmt is running...}
	./scripts/fmt.sh
.PHONY: fmt

## vet project
vet:
	${call colored, vet is running...}
	./scripts/vet-lint.sh
.PHONY: vet

## Lint fmt
fmt-lint:
	${call colored, fmt-lint is running...}
	./scripts/fmt-lint.sh
.PHONY: fmt-lint

## Installs tools from vendor.
install-tools:
	./scripts/install-tools.sh
.PHONY: install-tools

## Sync vendor of root project and tools.
sync-vendor:
	./scripts/sync-vendor.sh
.PHONY: sync-vendor

## Test the project.
test:
	${call colored, test is running...}
	./scripts/run-tests.sh
.PHONY: test


## Test coverage report.
test-cover:
	${call colored, test-cover is running...}
	./scripts/test-coverage.sh
.PHONY: test-cover