VERSION := $(shell cat internal/version/VERSION)+$(shell git rev-parse --short=7 HEAD)

GOCACHE = $(shell go env GOCACHE)
GOPATH = $(shell go env GOPATH)
GOVERSION = $(patsubst go%,%,$(shell go env GOVERSION))

.PHONY: all
all: build website

.PHONY: build
build: generate
ifeq ($(shell uname),Linux)
	make internal/fed
	make internal/ibgames
	make fedtpd
	make httpd
	make login
	make modemd
	make perivale
	# make perivale-go
	make workbench
	# make workbench-go
else
	@docker buildx build \
		--build-arg "GOVERSION=$(GOVERSION)" \
		--file Dockerfile.build \
		--platform linux/amd64 \
		--tag federation-build:latest \
		.
	@docker run \
		--env "GOCACHE=$(GOCACHE)" \
		--env "GOPATH=$(GOPATH)" \
		--env "HOME=$(HOME)" \
		--mount "type=bind,source=$(GOCACHE),target=$(GOCACHE)" \
		--mount "type=bind,source=$(GOPATH),target=$(GOPATH)" \
		--mount "type=bind,source=$(HOME)/.config/git,target=$(HOME)/.config/git,ro" \
		--mount "type=bind,source=$(PWD),target=$(PWD)" \
		--mount type=tmpfs,dst=/tmp \
		--platform linux/amd64 \
		--read-only \
		--rm \
		--user "$(shell id -u):$(shell id -g)" \
		--workdir "$(PWD)" \
		federation-build:latest \
		make "$@"
endif

.PHONY: fedtpd httpd login perivale-go workbench-go
fedtpd httpd login perivale-go workbench-go: bin/$(shell go env GOOS)-$(shell go env GOARCH)
	go build -o bin/$(shell go env GOOS)-$(shell go env GOARCH)/$@ -ldflags=-w ./cmd/$@

.PHONY: modemd perivale workbench
modemd perivale workbench: bin/$(shell go env GOOS)-$(shell go env GOARCH)
	make -C cmd/$@

bin/linux-amd64:
	@mkdir -p $@

.PHONY: internal/fed
internal/fed: internal/ibgames
	$(MAKE) -C $@ all

.PHONY: internal/ibgames
internal/ibgames:
	$(MAKE) -C $@ all

.PHONY: dev
dev: build
	./scripts/run-local.sh

.PHONY: deploy
deploy: generate website build
	fly deploy --local-only --strategy immediate

.PHONY: docker
docker: generate website build
	./scripts/docker-build.sh "$(VERSION)"
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
	cd web && npm ci
	cd web && hugo build --minify

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
clean: clean-coverage clean-fed clean-generated clean-ibgames clean-modemd clean-perivale clean-website clean-workbench
	$(RM) -r bin/

.PHONY: clean-coverage
clean-coverage:
	$(RM) coverage.out

.PHONY: clean-fed
clean-fed:
	make -C internal/fed clean

.PHONY: clean-generated
clean-generated:
	$(RM) internal/server/parser/parser.tab.go internal/server/parser/y.output
	$(RM) internal/text/messages.go
	$(RM) internal/workbench/messages.go

.PHONY: clean-ibgames
clean-ibgames:
	make -C internal/ibgames clean

.PHONY: clean-modemd
clean-modemd:
	make -C cmd/modemd clean

.PHONY: clean-perivale
clean-perivale:
	make -C cmd/perivale clean

.PHONY: clean-website
clean-website:
	$(RM) -r web/node_modules/ web/public/

.PHONY: clean-workbench
clean-workbench:
	make -C cmd/workbench clean
