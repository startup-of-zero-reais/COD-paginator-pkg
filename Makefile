GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

usage:
	@printf "\nUtilize um dos comandos:\n\nModo de uso:\n\tmake run\n\tmake tests\n"

run:
	echo "FAKE RUN"

.PHONY: gen_mocks
gen_mocks:
	mockery --all --keeptree

.PHONY: tests
tests:
	mkdir -p tests
	$(GOTEST) -coverprofile=./tests/coverage.out ./...
	$(GOCOVER) -func=./tests/coverage.out
	$(GOCOVER) -html=./tests/coverage.out