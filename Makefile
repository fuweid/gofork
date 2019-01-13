NAME := Wei Fu
EMAIL := fuweid89@gmail.com

# command name
COMMAND=gofork

# install into dir
DESTDIR ?= /usr
INSTALLDIR=${DESTDIR}/bin

.PHONY: binary install help

binary: ## build the binary
	go build -o bin/gofork github.com/fuweid/gofork/cmd

install: ## install the binary
	@mkdir -p $(INSTALLDIR)
	@install bin/$(COMMAND) $(INSTALLDIR)

help: ## help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
