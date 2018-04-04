.PHONY: clean build 

all: release

clean:
	rm -rf ./build

build:
	rm -rf ./build
	-mkdir ./build 
	go build -a -o build/tlparser

release:  build 
