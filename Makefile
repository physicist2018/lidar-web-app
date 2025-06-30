build:
	go build -o bin/myapp cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

docker-build:
	docker build -t myapp .

docker-run:
	docker run -p 8080:8080 myapp