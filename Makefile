export tag=v1.0
export dockerhub=sczhlpzz2008
build:
	echo "building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t httpserver:${tag} .

push:release
	echo "pushing homework_httpserver"
	docker push ${dockerhub}/homework_httpserver:${tag}
