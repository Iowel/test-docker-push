
docker-build-and-push 
	docker buildx build --no-cache --platform linux/amd64 -t <REGISTRY>/test-server:v0.0.1 .
	docker login -u <USERNAME> -p <PASSWORD> <REGISTRY>
	docker push <REGISTRY>/test-server:v0.0.1