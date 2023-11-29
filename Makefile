run:
#Example: make run CEP=99999999
	@go run main.go -c $(CEP)

build:
	@go build -o ./bin/viacep .