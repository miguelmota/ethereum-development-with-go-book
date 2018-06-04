all:
	@echo "no default"

install:
	npm install gitbook-cli@latest -g

serve:
	gitbook serve

build:
	gitbook build

deploy: build pdf ebook mobi
	mv ethereum-*.pdf _book/
	mv ethereum-*.epub _book/
	mv ethereum-*.mobi _book/
	./deploy.sh

pdf:
	gitbook pdf ./ ethereum-development-with-go.pdf

ebook:
	gitbook epub ./ ethereum-development-with-go.epub

mobi:
	gitbook mobi ./ ethereum-development-with-go.mobi

plugins/install:
	gitbook install
