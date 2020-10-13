# Go parameters
GO           = go
TIMEOUT_UNIT = 5m

.PHONY: all
all: test build-mac

.PHONY: build
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -o stress-test-windows -v ./cmd/...
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -o stress-test-linux -v ./cmd/...
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build  -o stress-test-mac -v ./cmd/... 

.PHONY: test
test:
	$(GO) test -timeout $(TIMEOUT_UNIT) -v ./test/...

.PHONY: clean
clean:
	$(GO) clean
	@rm -rf test/tests.* test/coverage.*