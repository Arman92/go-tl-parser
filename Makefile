.PHONY: clean build 

all: release

clean:
	rm -f migrate

build:
	rm -rf ./build
	-mkdir ./build 
	go build -a -o build/tlparser

release:  build 
