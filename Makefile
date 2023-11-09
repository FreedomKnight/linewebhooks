run: main
	./main

build-mongo:
	docker run --name mongo -p 27017:27017 -d mongo:5.0

main: main.go */*.go
	go mod tidy
	go build -o main main.go

clean:
	rm -rf main

.phony: build-mongo run clean
