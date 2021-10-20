mkfile_path:=$(abspath $(lastword $(MAKEFILE_LIST)))
mkfile_dirpath:=$(shell cd $(shell dirname $(mkfile_path)); pwd)

cmdDir:=cmd
outputDir:=_output
binaryName:=ip-thing
hackDir:=hack

## version mgmt
VERSION := $(shell git rev-parse --short HEAD)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

help: ## show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

clean: ## cleans output
	@rm -r $(mkfile_dirpath)/$(outputDir) || exit 0

prep-outputdir: ## creates output directory
	@mkdir -p $(mkfile_dirpath)/$(outputDir)

#Depend on get-kubernetes?
build-binary: prep-outputdir ## builds binary and drop it  in output directory
	@echo "** building binary with version:$(VERSION)"
	@go build -o $(mkfile_dirpath)/$(outputDir)/$(binaryName) $(GOFLAGS) $(mkfile_dirpath)/$(cmdDir)
	@echo "** built binary is at:$(mkfile_dirpath)/$(outputDir)/$(binaryName) "

unit-tests:  ## runs unit tests
	@echo "** running unit test in @ $(mkfile_dirpath)/pkg/generators/types"
	@go test $(mkfile_dirpath)/pkg/generators/types  -count=1 $(ADD_TEST_ARGS)  || exit 1
	@echo "** running unit test in @ $(mkfile_dirpath)/pkg/generators/"
	@go test $(mkfile_dirpath)/pkg/generators/  -count=1 $(ADD_TEST_ARGS)  || exit 1
