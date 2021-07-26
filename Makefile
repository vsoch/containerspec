all:
	gofmt -s -w .
	go build -o containerspec
	
run:
	go run main.go
