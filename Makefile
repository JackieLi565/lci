IMAGE = jackie565/lci:latest
CONTAINER = lci-app

build:
	docker build -t $(IMAGE) .

run:
	docker run --name $(CONTAINER) $(IMAGE)

clean:
	docker rm $(CONTAINER)
	docker rmi $(IMAGE)