SERVICE =order-pack-calculator

clean:
	rm -rf ./bin

test:
	go test ./...

build: clean
	mkdir bin
	GOOS=linux GOARCH=amd64 go build -o bin/$(SERVICE) main.go

build-docker:
	docker build -t $(SERVICE) .

run-docker:
	docker run -p 8090:8090 -d $(SERVICE)