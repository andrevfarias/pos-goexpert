run:
	go run cmd/cli/main.go --url='https://httpstat.us/Random/200,201,404,500-504' --requests=50 --concurrency=30

docker-build:
	docker build -t goex-stress-tester .

docker-run:
	docker run --rm goex-stress-tester --url='https://www.google.com' --requests=50 --concurrency=3

docker-test: docker-build docker-run
