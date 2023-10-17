run:
		go run cmd/main.go


test:
		go test ./internal/delivery/archiveInfo

build:
		docker build -t app .

run-img:
		docker run --name=app -p 8081:8081 --rm -d app