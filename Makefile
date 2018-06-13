all:
	@echo "no default"

install:
	npm install gitbook-cli@latest -g

serve:
	gitbook serve

build:
	gitbook build

deploy:
	./deploy.sh

deploy/all: build pdf ebook mobi deploy

pdf:
	gitbook pdf ./ ethereum-development-with-go.pdf

ebook:
	gitbook epub ./ ethereum-development-with-go.epub

mobi:
	gitbook mobi ./ ethereum-development-with-go.mobi

plugins/install:
	gitbook install
