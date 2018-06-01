all:
	@echo "no default"

install:
	npm install gitbook-cli@latest -g

serve:
	gitbook serve

build:
	gitbook build

deploy: build pdf
	mv ethereum-*.pdf _book/
	./deploy.sh

pdf:
	gitbook pdf ./ ethereum-development-with-go.pdf

ebook:
	gitbook epub ./ ethereum-development-with-go.epub

plugins/install:
	gitbook install
