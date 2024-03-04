gen:
	go generate ./...

docker-build:
	docker buildx build . --build-arg APP_NAME=gogolook-exercise -f docker/Dockerfile -t gogolook-exercise

docker-run:
	docker run --name gogolook-exercise -d -p 8080:8080 gogolook-exercise

tests:
	go test -v  ./usecase/.