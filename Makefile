GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_DEPS) -u
GO_FMT=gofmt
GO_INSTALL=$(GO_CMD) install
GO_RUN=$(GO_CMD) run
GO_TEST=$(GO_CMD) test
TOOLS=$(GO_RUN) tools/*.go

PACKAGE_LIST := $$(go list ./... | grep -v '/vendor/')

.PHONY: all alpha build deps fmt fmtcheck generate install release test test_deps update_deps version

all: build

alpha:
	@$(TOOLS) version --bump patch --prerelease alpha
	git add sl/version.go
	git commit -m "Bump version"
	git push

build: fmtcheck deps
	$(GO_BUILD) ./...

deps:
	$(GO_DEPS) ./...

fmt:
	@$(GO_FMT) -w `find . -name '*.go' | grep -v vendor`

fmtcheck:
	@fmt_list=$$($(GO_FMT) -e -l `find . -name '*.go' | grep -v vendor`) && \
	[ -z $${fmt_list} ] || \
		(echo "gofmt needs to be run on the following files:" \
			&& echo "$${fmt_list}" && \
			echo "You can run 'make fmt' to format code" && false)

generate:
	@$(TOOLS) generate

install: fmtcheck deps
	@$(GO_INSTALL) ./...

release: build
	@NEW_VERSION=$$($(TOOLS) version --bump patch) && \
	git add sl/version.go && \
	git commit -m "Cut release $${NEW_VERSION}" && \
	git tag $${NEW_VERSION} && \
	git push && \
	git push origin $${NEW_VERSION}

test: fmtcheck test_deps
	@$(GO_TEST) $(PACKAGE_LIST) -timeout=30s -parallel=4

test_deps:
	$(GO_DEPS) -t ./...

update_deps:
	$(GO_DEPS_UPDATE) ./...

version:
	@$(TOOLS) version
