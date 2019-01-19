default: install
build:
	@mkdir -p build/
	@go build -o build/creditcard main.go
install:
	@go get -u -v github.com/kavirajk/creditcard
clean:
	rm -rf build/
