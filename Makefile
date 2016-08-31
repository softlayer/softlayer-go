
all:
	go build github.ibm.com/riethm/gopherlayer.git/...

fmt:
	gofmt -w `find . -name '*.go' | grep -v vendor`
