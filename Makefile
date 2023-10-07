

#docker image name
DOCKER_IMAGE_NAME= "333448583258.dkr.ecr.us-east-2.amazonaws.com/ofsistorage-crawler:latest"


#create docker build
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .


docker-push:
	docker push $(DOCKER_IMAGE_NAME)

login:
	aws ecr get-login-password --region us-east-2 --profile itelite | docker login --username AWS --password-stdin 333448583258.dkr.ecr.us-east-2.amazonaws.com

deploy: login docker-build docker-push