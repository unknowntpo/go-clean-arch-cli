BIN=tm

build:
	go build -o $(BIN)

clean:
	-rm -f $(BIN)
