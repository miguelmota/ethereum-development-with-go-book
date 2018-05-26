all:
	@echo "no default"

serve:
	gitbook serve

build:
	gitbook build

deploy: build pdf
	mv ethereum-*.pdf _book/
	./deploy.sh

pdf:
	gitbook pdf ./ ethereum-development-with-go
