VERSION := $(shell go run ./tools/version 2>/dev/null || echo "dev")

export CGO_ENABLED := 0

.PHONY: all
all: build

.PHONY: build
build: build-$(shell go env GOOS)-$(shell go env GOARCH)

.PHONY: build-darwin-arm64
build-darwin-arm64: export GOARCH := arm64
build-darwin-arm64: export GOOS := darwin
build-darwin-arm64: bin/darwin-arm64 generate
	# go vet ./...
	go build -o bin/$(GOOS)-$(GOARCH)/ ./...

.PHONY: build-linux-amd64
build-linux-amd64: export GOARCH := amd64
build-linux-amd64: export GOOS := linux
build-linux-amd64: bin/linux-amd64 generate
	# go vet ./...
	go build -o bin/$(GOOS)-$(GOARCH)/ ./...

.PHONY: build-linux-arm64
build-linux-arm64: export GOARCH := arm64
build-linux-arm64: export GOOS := linux
build-linux-arm64: bin/linux-arm64 generate
	# go vet ./...
	go build -o bin/$(GOOS)-$(GOARCH)/ ./...

bin/linux-amd64 bin/darwin-arm64 bin/linux-arm64:
	@mkdir -p $@

.PHONY: dev
dev: build
	./scripts/run-local.sh 2>&1 | tee log/dev.log

.PHONY: deploy
deploy: build-linux-amd64
	fly deploy --local-only --strategy immediate

.PHONY: docker
docker: docker-local

.PHONY: docker-deploy
docker-deploy: docker-build-deploy
	./scripts/run-docker.sh amd64 2>&1 | tee log/docker-deploy.log

.PHONY: docker-build-deploy
docker-build-deploy: build-linux-amd64
	./scripts/docker-build.sh amd64 $(VERSION)

.PHONY: docker-local
docker-local: docker-build-local
	./scripts/run-docker.sh arm64 2>&1 | tee log/docker-local.log

.PHONY: docker-build-local
docker-build-local: build-linux-arm64
	./scripts/docker-build.sh arm64 $(VERSION)

.PHONY: generate
generate: internal/server/parser/parser.tab.go \
	internal/text/messages.go \
	internal/version/version.go

.DELETE_ON_ERROR: internal/server/parser/parser.tab.go
internal/server/parser/parser.tab.go: internal/server/parser/parser.tab.y
	@-go get -tool golang.org/x/tools/cmd/goyacc
	@go tool goyacc -o $@ -v $(<D)/y.output $<

.DELETE_ON_ERROR: internal/text/messages.go
internal/text/messages.go: data/messages.txt tools/gen-text.awk tools/gen-text.sh tools/gen-text.pl
	@tools/gen-text.sh $< >$@

.DELETE_ON_ERROR: internal/version/version.go
internal/version/version.go: tools/genversion/main.go
	@go run ./tools/genversion

.PHONY: test
test: bin/$(shell go env GOOS)-$(shell go env GOARCH) generate
	go test ./...

.PHONY: lint
lint: generate
	golangci-lint run

.PHONY: version
version:
	@go run ./tools/version

.PHONY: clean
clean: clean-generated
	$(RM) -r bin/

.PHONY: clean-generated
clean-generated:
	$(RM) internal/server/parser/parser.tab.go internal/server/parser/y.output
	$(RM) internal/text/messages.go
	$(RM) internal/version/version.go
	$(RM) internal/workbench/messages.go
