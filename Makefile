run:
	clear
	go run main.go token.go tokenTypes.go lexicalAnalyzer.go
build:
	clear
	go build -o bin/main main.go token.go tokenTypes.go lexicalAnalyzer.go