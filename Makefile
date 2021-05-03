PHONY: .all

build: 
	go build cmd/main.go

list:
	./main list
