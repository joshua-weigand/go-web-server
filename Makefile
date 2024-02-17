build:
	go build -C src/ -o ../bin/application

run: build
	./bin/application

test:
	go test -v ./...