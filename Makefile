all: build

.PHONY: install
install:
	npm install gitbook-cli@latest -g

.PHONY: serve
serve:
	gitbook serve

.PHONY: build
build:
	gitbook build

.PHONY: deploy
deploy:
	./deploy.sh

.PHONY: deploy/all
deploy/all: build pdf ebook mobi deploy

.PHONY: pdf
pdf:
	gitbook pdf ./ ethereum-development-with-go.pdf

.PHONY: ebook
ebook:
	gitbook epub ./ ethereum-development-with-go.epub

.PHONY: mobi
mobi:
	gitbook mobi ./ ethereum-development-with-go.mobi

.PHONY: plugins/install
plugins/install:
	gitbook install
