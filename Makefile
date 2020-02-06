run:
	go run main.go token.go lexicalAnalyzer.go
build:
	go build -o bin/main main.go token.go lexicalAnalyzer.go