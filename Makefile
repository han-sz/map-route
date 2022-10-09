
.PHONY: build run

build:
	go build -o build/map-route ./map-route
run:
	go run ./map-route