build-mongo:
	docker run --name mongo -p 27017:27017 -d mongo

run: main.go
	go mod tidy
	go build -o main main.go
	./main

.phony: build-mongo
