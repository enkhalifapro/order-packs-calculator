SERVICE =order-pack-calculator

clean:
	rm -rf ./bin

test:
	go test ./...

build: clean
	mkdir bin

	GOOS=linux GOARCH=amd64 go build -v -a -tags scheduler -o bin/$(SERVICE) main.go

build-docker:
	docker build -t $(SERVICE) .