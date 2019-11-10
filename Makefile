build:
	go build

test:
	go test -timeout 300ms -v -coverprofile coverage.out ./...
	go tool cover -html=coverage.out

start: build
	./slack-repeat-bot ${ARGS}

lint:
	go vet
