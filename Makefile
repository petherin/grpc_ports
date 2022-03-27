runsvc: ##@application Run service as Go app
	cd portsvc && go run cmd/grpc/main.go

proto: ##@gRPC Generate code from proto file
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative portsvc/proto/ports.proto

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