
all: fmtcheck
	go build github.ibm.com/riethm/gopherlayer.git/...

fmt:
	gofmt -w `find . -name '*.go' | grep -v vendor`

fmtcheck:
	@[ -z $$(gofmt -e -l `find . -name '*.go' | grep -v vendor`) ] || \
		(echo "run 'make fmt'" && false)