export tag=v1.0
build:
	echo "building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t homework/httpserver:${tag} .

push:release
	echo "pushing homework/httpserver"
	docker push homework/httpserver:${tag}
