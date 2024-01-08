all: build

build:
	go build -o bd

clean:
	rm ./bd
