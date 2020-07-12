PROJECT="QQGroupBot/qqbot"
GOPATH ?= $(shell go env GOPATH)
P=8

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

GO              := GO111MODULE=on go
GOBUILD         := $(GO) build $(BUILD_FLAG) -tags codes
GOBUILDCOVERAGE := GOPATH=$(GOPATH) cd tidb-server; $(GO) test -coverpkg="../..." -c .
GOTEST          := $(GO) test -p $(P)
OVERALLS        := GO111MODULE=on overalls
STATICCHECK     := GO111MODULE=on staticcheck

ARCH      := "`uname -s`"
LINUX     := "Linux"
MAC       := "Darwin"
PACKAGE_LIST  := go list ./...| grep -vE "cmd"
PACKAGES  := $$($(PACKAGE_LIST))
PACKAGE_DIRECTORIES := $(PACKAGE_LIST) | sed 's|github.com/mapleFU/$(PROJECT)/||'
FILES     := $$(find $$($(PACKAGE_DIRECTORIES)) -name "*.go")
# go source files, ignore vendor directory
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

default:
	$(GOBUILD)

format: fmt

fmt:
	@gofmt -s -l -w $(SRC)

check:
	@test -z $(shell gofmt -l main.go | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet ${SRC}

# lint:tools/bin/revive
# 	@echo "linting"
# 	@tools/bin/revive -formatter friendly -config tools/check/revive.toml $(FILES)

# tools/bin/revive: tools/check/go.mod
# 	cd tools/check; \
# 	$(GO) build -o ../bin/revive github.com/mgechev/revive

# tools/bin/goword: tools/check/go.mod
# 	cd tools/check; \
# 	$(GO) build -o ../bin/goword github.com/chzchzchz/goword