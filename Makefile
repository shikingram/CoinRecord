BINARY_NAME=CoinRecord

build:
	go build -o $(BINARY_NAME) -v

clean:
	go clean
