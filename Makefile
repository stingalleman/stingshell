.DEFAULT_GOAL := run

# https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04#step-4-%E2%80%94-building-executables-for-different-architectures

run:
	go run .

build-darwin:
	GOOS=darwin GOARCH=amd64 go build .

build-linux:
	GOOS=linux GOARCH=amd64 go build .

build-windows:
	GOOS=windows GOARCH=amd64 go build .