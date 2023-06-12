build:
	docker build -t debug-container .

buildx-create:
	docker buildx create --use --name buildx_instance

build-cross-platform:
	docker buildx build --platform linux/amd64,linux/arm64 -t kacy/debug-container . --push

run:
	docker run -p=8080:8080 --rm -it debug-container
