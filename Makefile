build-mongo:
	docker run --name mongo -p 27017:27017 -d mongo:5.0

main: main.go configs/*.go controllers/*.go models/*.go routes/*.go clients/*.go
	go mod tidy
	go build -o main main.go

run: main
	./main

clean:
	rm -rf main

.phony: build-mongo run clean
