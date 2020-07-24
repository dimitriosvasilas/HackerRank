
export GO111MODULE=on

all : test

.PHONY: test
## test: Run tests
test: ; $(info $(M) running $(NAME:%=% )tests...)
	@go test -v

.PHONY: help
## help: Prints this message
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'