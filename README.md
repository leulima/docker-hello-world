# docker-hello-world

## Build docker local image

	$ docker build . --tag hello-world

## Run
	$ docker run hello-world


## Build docker for GitHub Container Registry image

	# docker build . --tag ghcr.io/leulima/hello-world:latest

## Run
	$ docker run ghcr.io/leulima/hello-world:latest

## Push
	$ docker push ghcr.io/leulima/hello-world:latest

### Reference
https://blog.codecentric.de/en/2021/03/github-container-registry/