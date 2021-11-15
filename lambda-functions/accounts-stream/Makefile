all: tidy build zip clean
tidy:
	GOOS=linux go mod tidy
build: tidy
	GOOS=linux go build main.go
zip:
	zip function.zip main
clean:
	rm -rf main
