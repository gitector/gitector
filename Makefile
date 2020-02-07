build: gitector
	go build

install: build
	sudo mv gitector /usr/local/bin/gitector
