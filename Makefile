build:
	sudo docker build -t slack-repeat-bot .

test:
	go test -timeout 300ms -v -coverprofile coverage.out ./...
	go tool cover -html=coverage.out

start: build
	sudo docker run slack-repeat-bot ${ARGS}

lint:
	go vet
