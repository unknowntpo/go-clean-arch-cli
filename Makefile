BIN=tm

build:
	go build -o $(BIN) app/*.go

clean:
	-rm -f $(BIN)
