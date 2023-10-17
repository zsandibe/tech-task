build:
	docker build -t archive .
run-img:
	docker run --name=archive -p 8081:8081 --rm -d archive 
run:
	go run cmd/main.go
stop:
	docker stop archive