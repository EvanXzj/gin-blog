build: go-build docker-build

go-build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

docker-build:
	docker build -t xuzhijian/gin-blog .
