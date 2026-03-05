APP := oaicheck

.PHONY: fmt test build tidy clean

fmt:
	go fmt ./...

test:
	go test ./...

build:
	go build -o $(APP) .

tidy:
	go mod tidy

clean:
	rm -f $(APP)
