
TOOLS=go run ./tools
VETARGS?=-all
COVERPROFILE=coverage.out

PACKAGE_LIST := $$(go list ./... | grep -v '/vendor/')

.PHONY: all alpha build deps fmt fmtcheck generate release test coverage  version vet

all: build

alpha:
	@$(TOOLS) version --bump patch --prerelease alpha
	git add sl/version.go
	git commit -m "Bump version"
	git push

build: fmtcheck vet deps
	go build ./...

deps:
	go mod vendor

fmt:
	gofmt -w `find . -name '*.go' | grep -v vendor`

fmtcheck:
	@fmt_list=$$(gofmt -e -l `find . -name '*.go' | grep -v vendor`) && \
	[ -z $${fmt_list} ] || \
		(echo "gofmt needs to be run on the following files:" \
			&& echo "$${fmt_list}" && \
			echo "You can run 'make fmt' to format code" && false)

generate:
	go run ./tools generate

release: build
	@NEW_VERSION=$$($(TOOLS) version --bump patch) && \
	git add sl/version.go && \
	git commit -m "Cut release $${NEW_VERSION}" && \
	git tag $${NEW_VERSION} && \
	git push && \
	git push origin $${NEW_VERSION}

test: fmtcheck vet deps
	go test $(PACKAGE_LIST) -timeout=30s

coverage:
	@echo "Running unit tests. Cover profile saved to $(COVERPROFILE) ...\n"
	go test $(PACKAGE_LIST) -timeout=30s  -coverprofile=$(COVERPROFILE)
	@echo "\nBuilding function coverage report...\n"
	go tool cover -func=$(COVERPROFILE)

version:
	@$(TOOLS) version

# vet runs the Go source code static analysis tool `vet` to find
# any common errors.
vet:
	@go vet $(VETARGS) $$(go list ./... | grep -v datatypes) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi
