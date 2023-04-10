#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match --tags 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::') # grab everything after the space in "github.com/tendermint/tendermint v0.34.7"
HTTPS_GIT := https://github.com/furynet/furynetwork.git
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR)/proto:/workspace --workdir /workspace bufbuild/buf
BUILDDIR ?= $(CURDIR)/build

GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(FURY_BUILD_OPTIONS)))
  build_tags += gcc
endif

ifeq (secp,$(findstring secp,$(FURY_BUILD_OPTIONS)))
  build_tags += libsecp256k1_sdk
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=fury \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=furyd \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(FURY_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(FURY_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(FURY_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(FURY_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif

ifeq (,$(findstring nostrip,$(FURY_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(FURY_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

###############################################################################
###                                  Build                                  ###
###############################################################################

check_version:
ifneq ($(GO_MINOR_VERSION),18)
	@echo "ERROR: Go version 1.18 is required for this version of FURY. Go 1.19 has changes that are believed to break consensus."
	exit 1
endif

all: install lint test

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): check_version go.sum $(BUILDDIR)/
	GOWORK=off go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	echo "Ensure dependencies have not been modified ..." >&2
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/furyd -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf $(CURDIR)/artifacts/

distclean: clean
	rm -rf vendor/


###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.11.2
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-gen proto-docs

new-proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protoc-gen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@$(protoImage) find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

proto-update-deps:
	@echo "Updating Protobuf dependencies"
	$(DOCKER) run --rm -v $(CURDIR)/proto:/workspace --workdir /workspace $(protoImageName) buf mod update

containerProtoVer=v0.2
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=cosmos-sdk-proto-gen-$(containerProtoVer)

proto-gen:
	@echo "Generating Protobuf files"
	docker rm $(containerProtoGen) || true
	docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) sh ./scripts/protocgen.sh

proto-docs:
	@echo "Generating Protobuf docs"
	docker rm $(containerProtoGen) || true
	docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) sh ./scripts/protoc-docs-gen.sh


###############################################################################
###                          Tools & Dependencies                           ###
###############################################################################

TOOLS_DESTDIR  ?= $(GOPATH)/bin
STATIK         = $(TOOLS_DESTDIR)/statik
RUNSIM         = $(TOOLS_DESTDIR)/runsim

# Install the runsim binary with a temporary workaround of entering an outside
# directory as the "go get" command ignores the -mod option and will polute the
# go.{mod, sum} files.
#
# ref: https://github.com/golang/go/issues/30515
runsim: $(RUNSIM)
$(RUNSIM):
	@echo "Installing runsim..."
	@(cd /tmp && ${GO_MOD} go get github.com/cosmos/tools/cmd/runsim@master)

statik: $(STATIK)
$(STATIK):
	@echo "Installing statik..."
	@(cd /tmp && go get github.com/rakyll/statik@v0.1.6)

contract-tools:
ifeq (, $(shell which stringer))
	@echo "Installing stringer..."
	@go get golang.org/x/tools/cmd/stringer
else
	@echo "stringer already installed; skipping..."
endif

ifeq (, $(shell which go-bindata.exe))
	@echo "Installing go-bindata..."
	@go get github.com/kevinburke/go-bindata/go-bindata
else
	@echo "go-bindata already installed; skipping..."
endif

ifeq (, $(shell which gencodec))
	@echo "Installing gencodec..."
	@go get github.com/fjl/gencodec
else
	@echo "gencodec already installed; skipping..."
endif

ifeq (, $(shell which protoc-gen-go))
	@echo "Installing protoc-gen-go..."
	@go get github.com/fjl/gencodec github.com/golang/protobuf/protoc-gen-go
else
	@echo "protoc-gen-go already installed; skipping..."
endif

ifeq (, $(shell which protoc))
	@echo "Please istalling protobuf according to your OS"
	@echo "macOS: brew install protobuf"
	@echo "linux: apt-get install -f -y protobuf-compiler"
else
	@echo "protoc already installed; skipping..."
endif

ifeq (, $(shell which solcjs))
	@echo "Installing solcjs..."
	@npm install -g solc@0.5.11
else
	@echo "solcjs already installed; skipping..."
endif

docs-tools:
ifeq (, $(shell which yarn))
	@echo "Installing yarn..."
	@npm install -g yarn
else
	@echo "yarn already installed; skipping..."
endif

tools: tools-stamp
tools-stamp: contract-tools docs-tools proto-tools statik runsim
	# Create dummy file to satisfy dependency and avoid
	# rebuilding when this Makefile target is hit twice
	# in a row.
	touch $@

tools-clean:
	rm -f $(RUNSIM)
	rm -f tools-stamp

docs-tools-stamp: docs-tools
	# Create dummy file to satisfy dependency and avoid
	# rebuilding when this Makefile target is hit twice
	# in a row.
	touch $@

.PHONY: runsim statik tools contract-tools docs-tools proto-tools  tools-stamp tools-clean docs-tools-stamp

go.sum: go.mod
	echo "Ensure dependencies have not been modified ..." >&2
	go mod verify
	go mod tidy



update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
.PHONY: update-swagger-docs

###############################################################################
###                           Tests & Simulation                            ###
###############################################################################

PACKAGES_UNIT=$(shell go list ./... | grep -E -v 'tests/simulator')
TEST_PACKAGES=./...

test: test-unit
test-all: test-unit test-race test-cover

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock norace' $(PACKAGES_UNIT)

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' $(PACKAGES_UNIT)

test-cover:
	@VERSION=$(VERSION) mkdir -p .audit/gotest >> /dev/null && go test -timeout 30m -mod=readonly -coverprofile=".audit/gotest/coverage.out" -json -covermode=atomic -tags="norace" $(PACKAGES_UNIT)  > .audit/gotest/report.json

benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_UNIT)

################################################################################
###                                 Linting                                  ###
################################################################################

golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

lint:
	@echo "🤖 Running linter..."
	go run $(golangci_lint_cmd) run --timeout=10m
	@echo "✅ Completed linting!"


format:
	$(golangci_lint_cmd) run ./... --fix
	@go run mvdan.cc/gofumpt -l -w x/ app/
	$(DOCKER) run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md" --fix

mdlint:
	@echo "--> Running markdown linter"
	$(DOCKER) run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md"

markdown:
	$(DOCKER) run -v $(PWD):/workdir ghcr.io/igorshubovych/markdownlint-cli:latest "**/*.md" --fix

###############################################################################
###                              Documentation                              ###
###############################################################################

update-swagger-docs: statik
	$(BINDIR)/statik -src=client/docs/swagger-ui -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
.PHONY: update-swagger-docs

godocs:
	@echo "--> Wait a few seconds and visit http://localhost:6060/pkg/github.com/FURY-Network/FURY/types"
	godoc -http=:6060

# This builds a docs site for each branch/tag in `./docs/versions`
# and copies each site to a version prefixed path. The last entry inside
# the `versions` file will be the default root index.html.
build-docs:
	@cd docs && \
	while read -r branch path_prefix; do \
		(git checkout $${branch} && npm install && VUEPRESS_BASE="/$${path_prefix}/" npm run build) ; \
		mkdir -p ~/output/$${path_prefix} ; \
		cp -r .vuepress/dist/* ~/output/$${path_prefix}/ ; \
		cp ~/output/$${path_prefix}/index.html ~/output ; \
	done < versions ;
.PHONY: build-docs

sync-docs:
	cd ~/output && \
	echo "role_arn = ${DEPLOYMENT_ROLE_ARN}" >> /root/.aws/config ; \
	echo "CI job = ${CIRCLE_BUILD_URL}" >> version.html ; \
	aws s3 sync . s3://${WEBSITE_BUCKET} --profile terraform --delete ; \
	aws cloudfront create-invalidation --distribution-id ${CF_DISTRIBUTION_ID} --profile terraform --path "/*" ;
.PHONY: sync-docs

###############################################################################
###                              SONARQUBE                                  ###
###############################################################################

analyze: ## analyze the project for code quality
	@$(MAKE) .sonar-scanner

SONAR_TOKEN ?= "SonarScannerToken"
SONAR_HOST_URL ?=http://localhost:9000
SONAR_PROJECT_KEY ?= $(PROJECT_KEY)

SONAR_WDR ?= /usr/src
SONAR_PROJECT_KEY := $(subst $(subst ,, ),:,$(subst /,--,$(SONAR_PROJECT_KEY)))
.sonar-scanner:
	$(DOCKER) run \
		--rm \
		-e SONAR_HOST_URL="$(SONAR_HOST_URL)" \
		-e SONAR_LOGIN="$(SONAR_TOKEN)" \
		-e SONAR_PROJECTKEY="$(SONAR_PROJECT_KEY)" \
		-v "${CURDIR}:$(SONAR_WDR)" \
		-w $(SONAR_WDR) \
		sonarsource/sonar-scanner-cli \
		-D"sonar.projectKey=$(SONAR_PROJECT_KEY)" \
		-Dproject.settings=sonar-project.properties
