export tag=v1.0
export dockerhub=sczhlpzz2008
export username=******* #处于安全考虑，隐去用户名和密码 希望理解
export password=******* #处于安全考虑，隐去用户名和密码 希望理解
build:
	echo "building httpserver binary"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .

release: build
	echo "building httpserver container"
	docker build -t ${dockerhub}/homework_httpserver:${tag} .

push:release
	echo "pushing homework_httpserver"
	docker login -u ${username} -p ${password}
	docker push ${dockerhub}/homework_httpserver:${tag}
