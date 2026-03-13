BINARY := name

.PHONY: all build test vet govulncheck lint clean

all: build

build:
	go build -o $(BINARY) ./...

test:
	go test ./... -v

vet:
	go vet ./...

govulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

lint:
	@echo "lint skipped (runner compatibility issue)"

clean:
	rm -f $(BINARY)
