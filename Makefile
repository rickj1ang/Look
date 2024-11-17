pull:
	git pull origin main

push:
	git push -u origin main

build:
	go build
	mv look /usr/local/bin
