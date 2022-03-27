IMAGE_NAME_CLI=petherin/portclient
IMAGE_NAME_SVC=petherin/portsvc
IMAGE_TAG=latest
SEMVER_IMAGE_TAG=$(tag)
CONTAINER_NAME=ports_client
DOCKER_COMPOSE=docker-compose

start: ##@application Start application in containers.
	$(DOCKER_COMPOSE) up -d

stop: ##@application Stops Go containers
	$(DOCKER_COMPOSE) down -v --remove-orphans
	$(DOCKER_COMPOSE) rm -v -f -s

logs: ##@application Outputs container logs
	$(DOCKER_COMPOSE) logs -f --tail 100 $(CONTAINER_NAME)

proto: ##@gRPC Generate code from proto file. Requires protoc installs. Would be better to run this in a container to have all the dependencies inside it.
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative portsvc/proto/ports.proto

build: ##@docker Builds all images from the Dockerfiles and applies two tags: 'latest' and provided tag e.g, tag=0.0.1 make build
	docker build --no-cache -f Dockerfile_client -t $(IMAGE_NAME_CLI):$(IMAGE_TAG) -t $(IMAGE_NAME_CLI):$(SEMVER_IMAGE_TAG) . && \
	docker build --no-cache -f Dockerfile_service -t $(IMAGE_NAME_SVC):$(IMAGE_TAG) -t $(IMAGE_NAME_SVC):$(SEMVER_IMAGE_TAG) .

push: ##@docker Pushes image to Docker with tags: 'latest' and provided tag e.g. tag=0.0.1 make push
	docker push $(IMAGE_NAME_CLI):$(SEMVER_IMAGE_TAG) && \
    docker push $(IMAGE_NAME_CLI):$(IMAGE_TAG) && \
    docker push $(IMAGE_NAME_SVC):$(SEMVER_IMAGE_TAG)  && \
  	docker push $(IMAGE_NAME_SVC):$(IMAGE_TAG)

login: ##@docker Login to Docker Hub (need to provide password)
	docker login --username petherin

# Color settings for the making the help information look pretty
GREEN  := $(shell tput -Txterm setaf 2)
WHITE  := $(shell tput -Txterm setaf 7)
YELLOW := $(shell tput -Txterm setaf 3)
RESET  := $(shell tput -Txterm sgr0)

# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
# A category can be added with @category
# link: https://gist.github.com/prwhite/8168133#gistcomment-1727513
HELP_FUN = \
	%help; \
	while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
	print "usage: make [target]\n\n"; \
	for (sort keys %help) { \
	print "${WHITE}$$_:${RESET}\n"; \
	for (@{$$help{$$_}}) { \
	$$sep = " " x (32 - length $$_->[0]); \
	print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
	}; \
	print "\n"; }

help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

# If you get 'debug/test/start is up to date', it's because Makefile assumes the commands refers to files you want to create.
# If it finds a file matching a command, it tells you it's already up to date.
# Add .PHONY so Makefile doesn't treat the command as a file. See https://stackoverflow.com/questions/3931741/why-does-make-think-the-target-is-up-to-date
.PHONY: help

# Running just the `make` command will now print out the help information
# instead of printing the first command in the file which is the "start" command
.DEFAULT_GOAL := help
