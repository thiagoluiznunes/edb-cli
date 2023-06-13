
install:
	go mod tidy

build: install
	go build -o bin/edbctl

release: build
	mv bin/edbctl /usr/local/bin
	chmod +x /usr/local/bin
