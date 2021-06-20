run: setup-local build-local run-local
	
setup-local:
	docker network create -d bridge video-task-network
	docker pull elasticsearch:7.13.2
	docker run -d --name elasticsearch --network=video-task-network -p 9200:9200 -e "discovery.type=single-node" elasticsearch:7.13.2
	docker pull redis:6.2.4
	docker run -d --name redis --network=video-task-network -p 6379:6379 redis:6.2.4

build-local:
	docker build -t video-parser-image .
	docker tag video-parser-image video-parser-image:latest

run-local:
	docker stop video-parser
	docker run --rm --name video-parser --network=video-task-network -d video-parser-image:latest
