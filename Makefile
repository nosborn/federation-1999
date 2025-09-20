VERSION := $(shell cat internal/version/VERSION)+$(shell git rev-parse --short=7 HEAD)

GOCACHE = $(shell go env GOCACHE)
GOPATH = $(shell go env GOPATH)
GOVERSION = $(patsubst go%,%,$(shell go env GOVERSION))

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
ifeq ($(shell uname),Darwin)
	@docker run \
		--env GOCACHE=$(GOCACHE) \
		--env GOPATH=$(GOPATH) \
		--env HOME=$(HOME) \
		--mount type=bind,source=$(GOCACHE),target=$(GOCACHE) \
		--mount type=bind,source=$(GOPATH),target=$(GOPATH) \
		--mount type=bind,source=$(HOME)/.config/git,target=$(HOME)/.config/git,ro \
		--mount type=bind,source=$(PWD),target=$(PWD) \
		--mount type=tmpfs,dst=/tmp \
		--platform $(GOOS)/$(GOARCH) \
		--read-only \
		--rm \
		--user $(shell id -u):$(shell id -g) \
		--workdir $(PWD) \
		golang:$(GOVERSION) \
		make $@
else
	# go vet ./...
	go build -o bin/$(GOOS)-$(GOARCH)/ ./...
endif

bin/darwin-arm64 bin/linux-amd64:
	@mkdir -p $@

.PHONY: dev
dev: build
	./scripts/run-local.sh

.PHONY: deploy
deploy: build-linux-amd64 website
	fly deploy --local-only --strategy immediate

.PHONY: docker
docker: build-linux-amd64 website
	./scripts/docker-build.sh $(VERSION)
	./scripts/run-docker.sh 2>&1 | tee log/docker-local.log

.PHONY: generate
generate: internal/server/parser/parser.tab.go \
	internal/text/messages.go

.DELETE_ON_ERROR: internal/server/parser/parser.tab.go
internal/server/parser/parser.tab.go: internal/server/parser/parser.tab.y
	@go tool goyacc -l -o $@ -v $(<D)/y.output $<

.DELETE_ON_ERROR: internal/text/messages.go
internal/text/messages.go: data/messages.txt tools/gen-text.awk tools/gen-text.sh tools/gen-text.pl
	@tools/gen-text.sh $< >$@

.PHONY: website
website:
	cd web && hugo build
	find web/public -name '*.html' -exec tidy -indent -wrap 0 -upper -quiet -modify {} 2>/dev/null \;
	find web/public -name '*.html' -exec sed -i '' '/META name="generator"/d' {} \;

.PHONY: test
test: bin/$(shell go env GOOS)-$(shell go env GOARCH) generate
ifeq ($(shell go env GOOS),darwin)
	go test -coverprofile=coverage.out -coverpkg=./cmd/...,./internal/... ./...
	@echo "Coverage Summary:"
	@go tool cover -func=coverage.out | grep "^total:" | awk '{print "Total coverage: " $3}'
else
	go test ./...
endif

.PHONY: lint
lint: generate
	golangci-lint run

.PHONY: clean
clean: clean-coverage clean-generated clean-website
	$(RM) -r bin/

.PHONY: clean-coverage
clean-coverage:
	$(RM) coverage.out

.PHONY: clean-generated
clean-generated:
	$(RM) internal/server/parser/parser.tab.go internal/server/parser/y.output
	$(RM) internal/text/messages.go
	$(RM) internal/workbench/messages.go

.PHONY: clean-website
clean-website:
	$(RM) -r web/public/
