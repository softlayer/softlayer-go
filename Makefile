GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_DEPS=$(GO_CMD) get -d -v
GO_FMT=gofmt
GO_INSTALL=$(GO_CMD) install
GO_TEST=$(GO_CMD) test

PACKAGE_LIST := $$(go list ./... | grep -v '/vendor/')

.PHONY: all alpha build deps fmt fmtcheck generate install install_tools release test

all: build

alpha:
	@$(TOOLS) version --bump patch --prerelease alpha && \
	git add version.go && \
	git commit -m "Bump version"
	@$(GO_INSTALL) . ./tools

build: fmtcheck
	$(GO_BUILD) ./...

deps:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Install dependencies for $$p ..."; \
		$(GO_DEPS) $$p || exit 1; \
	done

fmt:
	@$(GO_FMT) -w `find . -name '*.go' | grep -v vendor`

fmtcheck:
	@fmt_list=$$($(GO_FMT) -e -l `find . -name '*.go' | grep -v vendor`) && \
	[ -z $${fmt_list} ] || \
		(echo "gofmt needs to be run on the following files:" \
			&& echo "$${fmt_list}" && \
			echo "You can run 'make fmt' to format code" && false)

generate: install_tools
	@tools generate

install: fmtcheck
	@$(GO_INSTALL) ./...

install_tools: fmtcheck
	@$(GO_INSTALL) . ./tools

release: build install_tools
	@NEW_VERSION=$$(tools version --bump patch) && \
	git add version.go && \
	git commit -m "Cut release $${NEW_VERSION}" && \
	git tag $${NEW_VERSION}
	@$(GO_INSTALL) . ./tools

test: fmtcheck
	@$(GO_TEST) $(PACKAGE_LIST) -timeout=30s -parallel=4

version: install_tools
	@tools version