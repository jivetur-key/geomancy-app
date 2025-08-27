BINARY_NAME=geomancer

.PHONY: all build run clean

build:
	go build -o $(BINARY_NAME) ./main.go

run:
	go run ./main.go

clean:
	rm -f $(BINARY_NAME)
