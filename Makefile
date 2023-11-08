build-mongo:
	docker run --name mongo -p 27017:27017 -d mongo

main: main.go
	go build -o main main.go

run: main
	./main

.phony: build-mongo
